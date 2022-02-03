package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

/**************************************************************

Script for a program designed to terminate itself after a given time

Extended to add functionality for reacieving messages from a master over an UDP connection, 
and termiate after a set timeout from the master

***************************************************************/

func main() {

	laddy, _ := net.ResolveUDPAddr("udp4", "0.0.0.0:20012")
	conn, _ := net.ListenUDP("udp4", laddy)

	//conn.SetReadDeadline()
	//time.Sleep(3000 * time.Millisecond)
	reciever(conn)

	fmt.Println("Hi")
	for i := 5; i > 0; i-- {
		fmt.Println(i, " seconds until suicide")
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Println("Goodbye")
	time.Sleep(500 * time.Millisecond)
	conn.Close()
	os.Exit(3)
}

func reciever(inconn *net.UDPConn) {

	for {
		buffer := make([]byte, 100)
		inconn.SetReadDeadline(time.Now().Add(2 * time.Second))
		n, err := inconn.Read(buffer)
		if err != nil {
			fmt.Println("Timeout goodbye")
			time.Sleep(500 * time.Millisecond)
			inconn.Close()
			os.Exit(3)
		}
		//fmt.Println(n, "bytes recieved")
		fmt.Println(string(buffer[:n]))
		time.Sleep(1000 * time.Millisecond)
	}
}