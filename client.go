package main

import (
	"bufio"
	"fmt"
	"github.com/yamakenji24/tcp-playground/handler"
	"io"
	"log"
	"net"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func read(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", ":8080")
	handler.ErrHandler(err)

	go write(conn)
	read(os.Stdout, conn)
}

func nextText() string {
	fmt.Println("入力待ち~: ")
	sc.Scan()
	return sc.Text()
}

func write(conn net.Conn) {
	for {
		text := nextText()
		b := []byte(text)
		_, err := io.WriteString(conn, string(b))
		if err != nil {
			log.Println("Write ", err)
			return
		}
	}
}