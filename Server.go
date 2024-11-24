package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening on PORT 6379")
	listner, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := listner.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		resp := NewResp(conn)

		value, err := resp.Read()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		conn.Write([]byte("+OK\r\n"))

	}

}
