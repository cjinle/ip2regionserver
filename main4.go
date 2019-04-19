package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {
	fmt.Println("start ... ")
	region, err := ip2region.New("conf/ip2region.db")
	if err != nil {
		panic("ip2region.db err")
	}
	defer region.Close()

	file, err := os.Open("/data/wwwroot/test/ip3.log")
	if err != nil {
		panic("ip data open err!")
	}
	defer file.Close()
	t1 := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := scanner.Text()
		ipInfo, _ := region.MemorySearch(ip)
		b, _ := json.Marshal(ipInfo)
		fmt.Println(string(b))
	}
	if scanner.Err() != nil {
		panic(err)
	}

	t2 := time.Now()
	fmt.Println(t1, t2)
}
