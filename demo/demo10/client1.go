package main

import (
	"fmt"
	"github.com/jhinih/hin/hpack"
	"io"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8999")
	i := 0
	for {
		i++

		p := hpack.NewTLVPack()
		msg := hpack.NewMessage(1, []byte("hello"+strconv.Itoa(i)))

		binaryMsg, _ := p.Pack(msg)
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println(err)
		}
		time.Sleep(3 * time.Second)

		//回显
		binaryHead := make([]byte, p.GetHeadLen())
		if _, err := conn.Read(binaryHead); err != nil {
			fmt.Println(err)
		}
		Imessage, err := p.UnPack(binaryHead)
		if err != nil {
			fmt.Println(err)
		}
		if Imessage.GetDataLen() > 0 {
			msg := Imessage.(*hpack.Message)
			msg.Data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Recv————>" + strconv.Itoa(int(msg.ID)) + string(msg.Data))
		}
	}
}
