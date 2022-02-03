package main

import (
	"fmt"
	"net"
	"time"
)

/****************************************************************

Code for sending and recieving messages using UDP

 - Reciever
	Starts listening to input adress and prints out the recieved messages

 - main / UDP-server
	Starts a server which sends a message to a given adress every second 


*****************************************************************/

func main() {
	go reciever("0.0.0.0:30000")
	go reciever("0.0.0.0:20012")

	addr, err := net.ResolveUDPAddr("udp4", "10.100.23.240:20012")
	conn, err := net.DialUDP("udp4", nil, addr)
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	for {
		conn.Write([]byte("hello from #12"))
		time.Sleep(1000 * time.Millisecond)
	}

}

func reciever(strAdress string) {
	addr, err := net.ResolveUDPAddr("udp4", strAdress)
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("A dial error!")
	}
	defer conn.Close()
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("A read error!")
		}
		fmt.Println(n, "bytes recieved")
		fmt.Println(string(buffer[:n]))
		time.Sleep(1000 * time.Millisecond)
	}
}