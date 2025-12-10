package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func MetaData() (int, string, int64, string) { //

	namefile := "oitudobem.txt"
	archivetype := filepath.Ext(namefile)
	fileInfo, err := os.Stat(namefile)

	if err != nil {
		fmt.Printf("erro ao carregar arquivo")
	}

	String_Size := len(fileInfo.Name())

	return String_Size, fileInfo.Name(), fileInfo.Size(), archivetype
}

func Type(Identify_byte byte) byte { //primeira string é teste

	//identificador de inicio de comunicação
	StartEvent_Type := []byte{0x01} //Tipo de evento Start
	Chunk_Type := []byte{0x02}      //Tipo de evento Chunk
	End_Type := []byte{0x03}        //Tipo de evento End
	ACK_Type := []byte{0x04}        //Tipo de evento ACK
	Erro_Type := []byte{0x05}       //Tipo de evento Erro

	if Identify_byte == StartEvent_Type[0] {
		return StartEvent_Type[0]
	}
	if Identify_byte == Chunk_Type[0] {
		return Chunk_Type[0]
	}
	if Identify_byte == End_Type[0] {
		return End_Type[0]
	}
	if Identify_byte == ACK_Type[0] {
		return ACK_Type[0]
	}
	if Identify_byte == Erro_Type[0] {
		return Erro_Type[0]
	}

	return StartEvent_Type[0]

}

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

func handler(conn net.Conn) {
	defer conn.Close()

	buffer, err := Read_file()
	if err != nil {
		return
	}

	header := [4]byte{}
	header_value := [1]byte{}
	file_size := [8]byte{}

	_, filenametmp, filesizetmp, _ := MetaData()
	file_name := []byte(filenametmp)

	file_name_size := make([]byte, 4)
	binary.BigEndian.PutUint32(file_name_size, uint32(len(filenametmp)))

	fmt.Println("Nome do arquivo:", string(file_name))
	pos := 0
	tamanho := len(buffer)

	for pos < tamanho {

		header[0] = Type(header_value[0]) //Start Event
		binary.BigEndian.PutUint64(file_size[:], uint64(filesizetmp))

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

		go handler(conn)
	}

}
