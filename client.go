package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080") // IP do outro PC
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Digite: ")
        text, _ := reader.ReadString('\n')

        conn.Write([]byte(text))

        reply, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println("Servidor respondeu:", reply)
    }
}
