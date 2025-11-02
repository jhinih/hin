package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8999")
	i := 0
	for {
		i++
		conn.Write([]byte("你好" + strconv.Itoa(i)))

		time.Sleep(3 * time.Second)

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
		}
		fmt.Println(string(buf[:n]))
	}
}
