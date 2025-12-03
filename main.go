package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func header(){
	fmt.Println("Hello")
}


func payload(){
	fmt.Println("Hello")
}

func Separator(){ /*Caso para recebimento*/

}

func MetaData() (string, int64, string) { /*saída de TAMANHO é um int de 64 bits*/
	
	namefile := "oitudobem.txt"
	archivetype := filepath.Ext(namefile)
	fileInfo, err := os.Stat(namefile)
	
	if err != nil{
		fmt.Printf("erro ao carregar arquivo")
	}
	
	
	return fileInfo.Name(), fileInfo.Size(), archivetype   
}

func main(){

	Name, Size, Type := MetaData() /*se não quiser retorno coloque = "_" */
	fmt.Println("", Name, Size, Type)
}