package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listener, error := net.Listen("tcp", ":3383")

	if nil != error {
		log.Println(error)
	}
	defer listener.Close()

	for {
		conn, error := listener.Accept()
		if nil != error {
			log.Println(error)
			continue
		}
		defer conn.Close()
		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]

			start := time.Now()
			start_time := string(data)[:17]
			service_id := string(data)[17:28]
			request := string(data)[28:]

			result := ""
			if service_id == "service0001" {
				result = Service0001(request)
			}
			if service_id == "service0002" {
				result = Service0002(request)
			}

			elapsedTime := time.Since(start)
			// end := time.Since(start)
			second := elapsedTime.Milliseconds()

			_, err = conn.Write([]byte(fmt.Sprintf("%s%s%09d%s", start_time, service_id, second, result)))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
