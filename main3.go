package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {
	region, err := ip2region.New("conf/ip2region.db")
	if err != nil {
		panic("ip2region.db err")
	}
	defer region.Close()

	netListen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn, region)
	}
}

func handleConnection(conn net.Conn, region *ip2region.Ip2Region) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data string:", string(buffer[:n]))
		ipInfo, _ := region.MemorySearch(string(buffer[:n]))
		b, err := json.Marshal(ipInfo)
		Log(string(b), err)
		n, err = conn.Write(b)
		Log(n, err)
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}
