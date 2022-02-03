package main

import (
	"fmt"
	"net"
	"time"
)

/****************************************************************

Code for sending and recieving messages using TCP

 - main 
	Establishes a tcp connection to a school server using DialTCP for sending messages to server

	also makes own tcp server and lets school server connect to it and allow it to send messages to us

 - Read
	function for reading and printing messages recieved from a tcp connection

 - Write
	Function for sending a given str message to a tcp connection

*****************************************************************/

func main() {
	//go reciever("0.0.0.0:30000")
	addy := "10.100.23.240:33546"

	addr, err := net.ResolveTCPAddr("tcp", addy)
	conn, err := net.DialTCP("tcp", nil, addr)
	fmt.Println(conn.LocalAddr().String())
	errorHandling(err)
	go read(conn)
	go write(conn, "We coneted to server")

	conn.Write([]byte("Connect to: 10.100.23.218:20012\x00"))

	laddy := "10.100.23.218:20012"
	laddr, err := net.ResolveTCPAddr("tcp", laddy)
	listen, err := net.ListenTCP("tcp", laddr)
	for {
		inConn, err := listen.AcceptTCP()
		errorHandling(err)
		go read(inConn)
		go write(inConn, "Server hit us up")
	}

}

func read(conn *net.TCPConn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		errorHandling(err)
		fmt.Println(n, "bytes recieved. Local:", conn.LocalAddr().String(), " Remote:", conn.RemoteAddr().String())
		fmt.Println("   ", string(buffer[:n]))
		//time.Sleep(1000 * time.Millisecond)
	}
}

func write(conn *net.TCPConn, msg string) {
	for {
		time.Sleep(10 * time.Millisecond)
		conn.Write([]byte(msg + "\x00"))
		//errorHandling(err)
		time.Sleep(1000 * time.Millisecond)
	}
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}