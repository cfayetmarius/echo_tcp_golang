package main

import(
	"fmt"
	"net"
	"time"
	"os"
	"log"
	"bufio"
) 


func parse() string{
	return(os.Args[1])
}

func connect(addr string) net.Conn {
	conn, err := net.Dial("tcp",addr)
	if err != nil {
		log.Fatal(err)
	}
	return(conn)
}

func senddata(data string, conn net.Conn) {
	_, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("[*]Client started at ",time.Now().Format("2006-01-02 3:4:5"))
	//getting args as address
	serv := parse()
	//connecting
	conn := connect(serv)
	for {
		//getting what client says via input
		fmt.Printf(">")
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')
		//sending data to server
		senddata(msg,conn)
	}
	defer conn.Close()
}
