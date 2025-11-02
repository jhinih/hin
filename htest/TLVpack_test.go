package htest

import (
	"fmt"
	"github.com/jhinih/hin/hpack"
	"io"
	"net"
	"strconv"
	"testing"
)

func TestTLVPack(t *testing.T) {
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Listener error:", err)
	}
	defer listenner.Close()

	//server
	go func() {
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("Accept error:", err)
			}
			go func(conn net.Conn) {
				p := hpack.NewTLVPack()
				for {
					head := make([]byte, p.GetHeadLen())
					_, err := io.ReadFull(conn, head)
					if err != nil {
						fmt.Println("Read head error:", err)
					}

					msgHandler, err := p.UnPack(head)
					if err != nil {
						fmt.Println("Unpack error:", err)
					}

					if msgHandler.GetDataLen() > 0 {
						msg := msgHandler.(*hpack.Message)
						msg.Data = make([]byte, msg.GetDataLen())

						_, err = io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("Read data error:", err)
						}
						fmt.Println("——————>Recv data:", strconv.Itoa(int(msg.ID)), strconv.Itoa(int(msg.DataLen)), string(msg.Data))
						fmt.Println("Unpack data:", msg.GetData())
					}

				}

			}(conn)
		}
	}()

	//client
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Dial error:", err)
	}
	p := hpack.NewTLVPack()

	msg1 := &hpack.Message{
		ID:      1,
		DataLen: 3,
		Data:    []byte{'h', 'i', 'n'},
	}
	sendMsg1, err := p.Pack(msg1)
	if err != nil {
		fmt.Println("Pack error:", err)
	}
	msg2 := &hpack.Message{
		ID:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendMsg2, err := p.Pack(msg2)
	if err != nil {
		fmt.Println("Pack error:", err)
	}

	sendMsg := append(sendMsg1, sendMsg2...)

	conn.Write(sendMsg)

	select {}
}
