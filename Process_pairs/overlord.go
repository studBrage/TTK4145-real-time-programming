package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"time"
)

/**************************************************************

Overlord is an extension of the suicide_bot

The program essentially creates a copy of itself and passes a counter to the copy
After a set time the program terminates itself and the copy takes over the counting
The copy makes its own copy after noticing the first sender has died, and then becomes the sender itself

***************************************************************/

func main() {
	important_baby := 0

	{
		laddy, _ := net.ResolveUDPAddr("udp4", "0.0.0.0:20012")
		conn, _ := net.ListenUDP("udp4", laddy)

		important_baby = reciever(conn)
		conn.Close()
	}

	laddy, _ := net.ResolveUDPAddr("udp4", "0.0.0.0:20012")
	conn, _ := net.DialUDP("udp4", nil, laddy)

	fmt.Println("Spawning backup")
	cmd := exec.Command("gnome-terminal", "--", "go", "run", "overlord.go")
	cmd.Run()

	time.Sleep(500 * time.Millisecond)

	for msgs := 0; msgs <= 5; msgs++ {
		fmt.Println(important_baby)
		//str := "hello this is GOD nr." + fmt.Sprint(msgs)
		conn.Write([]byte(fmt.Sprint(important_baby)))
		important_baby++
		time.Sleep(500 * time.Millisecond)
	}

	conn.Close()
}

func reciever(inconn *net.UDPConn) int {
	temp_baby := 0
	for {
		buffer := make([]byte, 100)
		inconn.SetReadDeadline(time.Now().Add(1 * time.Second))
		n, err := inconn.Read(buffer)
		if err != nil {
			//fmt.Println("Timeout goodbye")
			time.Sleep(500 * time.Millisecond)
			inconn.Close()
			return temp_baby
		}
		//fmt.Println(n, "bytes recieved")
		//fmt.Println(strconv.Atoi(string(buffer[:n])))
		temp_baby, _ = strconv.Atoi(string(buffer[:n]))
		time.Sleep(100 * time.Millisecond)
	}
}