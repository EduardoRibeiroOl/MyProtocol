package main

import (
	"fmt"
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
	
	if err != nil{
		fmt.Printf("erro ao carregar arquivo")
	}
	
	String_Size := len(fileInfo.Name())

	//Retorna Tamanho do nome, nome, tamanho e tipo do arquivo
	return String_Size, fileInfo.Name(), fileInfo.Size(), archivetype   
}

func Type(name_file string, size_file int64) (string, int64){
		
	StartEvent_Type := []byte{0x01} //Tipo de evento Start
	Chunlk_Type := []byte{0x02} //Tipo de evento Chunk
	End_Type := []byte{0x03} //Tipo de evento End
	ACK_Type := []byte{0x04} //Tipo de evento ACK
	Erro_Type := []byte{0x05} //Tipo de evento Erro
	
		
	fmt.Println("Start Event Type:", StartEvent_Type, "Chunk Type:", Chunlk_Type, "End Type:", End_Type, "ACK Type:", ACK_Type, "Erro Type:", Erro_Type)
	return name_file, size_file
	//tem que retornar o tipo de arquivo que vai ser dependendo da função chamada 
	
}

/*
func Start_header() (int16){
	
}

func Chunk_SEQ() (){ 

}

func Payload() (int64){
	
}

func CRC32()(int32){

}*/


func main(){

	_, Archive_Name, file_size, _ := MetaData() /*se não quiser retorno coloque = "_" */
	
	fmt.Println(Type(Archive_Name, file_size))
}