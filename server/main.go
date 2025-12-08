package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

// Considerando: Nome_Func(quantidade_de Bytes)
// Start(2) => 16bits
// Type(1) => 8bits
// SEQ(4) => 32bits
// Length(4) => 32bits
// Payload(variable) => N*8bits | filename_length(1) | filename(N) | file_size(8 bytes) |
// CRC32(4) => 32bits

func MetaData() (int, string, int64, string) { //

	namefile := "oitudobem.txt"
	archivetype := filepath.Ext(namefile)
	fileInfo, err := os.Stat(namefile)

	if err != nil {
		fmt.Printf("erro ao carregar arquivo")
	}

	String_Size := len(fileInfo.Name())

	//Retorna Tamanho do nome, nome, tamanho e tipo do arquivo
	return String_Size, fileInfo.Name(), fileInfo.Size(), archivetype
}

func Type(Identify_byte byte, name_file string, size_file int64) (string, string, int64) { //primeira string é teste

	//identificador de inicio de comunicação
	StartEvent_Type := []byte{0x01} //Tipo de evento Start
	Chunk_Type := []byte{0x02}      //Tipo de evento Chunk
	End_Type := []byte{0x03}        //Tipo de evento End
	ACK_Type := []byte{0x04}        //Tipo de evento ACK
	Erro_Type := []byte{0x05}       //Tipo de evento Erro

	if Identify_byte == StartEvent_Type[0] {
		return "start", name_file, size_file
	}
	if Identify_byte == Chunk_Type[0] {
		return "chunk", name_file, size_file
	}
	if Identify_byte == End_Type[0] {
		return "end", name_file, size_file
	}
	if Identify_byte == ACK_Type[0] {
		return "ack", name_file, size_file
	}
	if Identify_byte == Erro_Type[0] {
		return "erro", name_file, size_file
	}

	fmt.Println("Start Event Type:", StartEvent_Type, "Chunk Type:", Chunk_Type, "End Type:", End_Type, "ACK Type:", ACK_Type, "Erro Type:", Erro_Type)
	return "", name_file, size_file

}

func Start_header() {
	_, Archive_Name, file_size, _ := MetaData() //se não quiser retorno coloque = "_"
	fmt.Println(Type(0x02, Archive_Name, file_size))
}

// divisa de trabalho :

func Read_file() ([]byte, error) {

	file, err := os.OpenFile("../oitudobem.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return nil, err
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return nil, err
	}

	fmt.Println("Conteúdo do arquivo:", string(data))
	return data, nil
}

func handle(conn net.Conn) {
	defer conn.Close()

	//reader := bufio.NewReader(conn)

	//buffer := make([]byte, 1024) buffer em martiz de 1kB
	buffer, err := Read_file()
	if err != nil {
		return
	}

	pos := 0
	tamanho := len(buffer)

	for pos < tamanho{
	

			end := pos + 1024
			if end > tamanho {
				end = tamanho
			}

			parte := buffer[pos:end]
			
			byte_sender, err := conn.Write(parte) 
			if err != nil {
				log.Println("Client saiu", err)
				return
			}
			
			pos += len(parte)
			fmt.Println("_", byte_sender)
	}
}

func main() {

	//Abre servidor TCP
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("esperando conexão")

	for { //loop infinito para aceitar conexões
		conn, err := ln.Accept() //aceita conexão
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}

}
