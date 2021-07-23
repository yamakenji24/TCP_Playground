package main

import (
	"context"
	"fmt"
	"github.com/yamakenji24/tcp-playground/handler"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tcpHandler(ctx)
}

func tcpHandler(ctx context.Context) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	handler.ErrHandler(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	handler.ErrHandler(err)

	defer listener.Close()

	log.Println("Starting TCP Server ")
	for {
		select {
		case <- ctx.Done():
			log.Println("Stopping TCP Server ")
			return
		default:
			listener.SetDeadline(time.Now().Add(time.Second * 30))
			conn, err := listener.AcceptTCP()
			if err != nil {
				switch err := err.(type) {
				case net.Error:
					if err.Timeout() {
						log.Println("Connection close")
						return
					}
				default:
					log.Println("Error on something else")
					return
				}
			}
			go echoHandler(conn)
		}
	}
}

func echoHandler(conn *net.TCPConn) {
	defer conn.Close()

	buf := make([]byte, 4*1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			if ne, ok := err.(net.Error); ok {
				switch {
				case ne.Temporary():
					continue
				}
			}
			return
		}
		text := fmt.Sprintf("response from server: %s \n", string(buf))
		log.Println("From client: ", string(buf))
		_, err = io.WriteString(conn, text)
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}