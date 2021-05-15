package main

import(
	"fmt"
	"net"
	"time"
	"bufio"
	"strings"
) 

func handle(conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		data = strings.TrimSpace(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data,"  written at ",time.Now().Format("2006-01-02 3:4:5"))
	}
	conn.Close()
}

func main() {
	fmt.Println("[*]Server started at ",time.Now().Format("2006-01-02 3:4:5"))
	//Listen
	ln, err := net.Listen("tcp",":9000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	//Accepting new cons and handling them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}