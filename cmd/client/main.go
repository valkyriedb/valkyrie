package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":42069")
	if err != nil {
		panic(err)
	}

	var i int
	buf := make([]byte, 1024)
	for {
		_, err := conn.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			panic(err)
		}
		_, err = conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf))
		i++
		time.Sleep(200 * time.Millisecond)
	}
}
