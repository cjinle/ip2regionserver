package main

import (
	"github.com/cjinle/test/ip2regionserver"
)

func main() {
	// ip2regionserver.GinListen()
	// ip2regionserver.HttpListen()
	ip2regionserver.TcpListen()
}
