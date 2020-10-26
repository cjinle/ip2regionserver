package ip2regionserver

import (
	"encoding/json"
	"fmt"
	"net"
	"io/ioutil"
)

var ConnMap map[string]net.Conn

func TcpStart() {
	defer region.Close()

	s, err := ioutil.ReadFile("../conf/app.json")
	CheckError(err)

	cfg := AppConf{}
	err = json.Unmarshal(s, &cfg)
	CheckError(err)

	ConnMap = make(map[string]net.Conn)
	netListen, err := net.Listen("tcp", cfg.Listen)
	CheckError(err)

	defer netListen.Close()

	fmt.Println("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		ConnMap[conn.RemoteAddr().String()] = conn
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		ipInfo, err := region.MemorySearch(string(buffer[:n]))
		if err != nil {
			conn.Write([]byte("[]"))
		} else {
			b, err := json.Marshal(ipInfo)
			if err != nil {
				b = []byte("[]")
			}
			fmt.Printf("request ip: %s\nresponse: %s\n", buffer[:n], b)
			conn.Write(b)
		}
	}
}

