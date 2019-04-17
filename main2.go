package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {
	fmt.Println("http server")
	region, err := ip2region.New("conf/ip2region.db")
	if err != nil {
		panic("ip2region.db err")
	}
	defer region.Close()

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		if ip == "" {
			ip = "0.0.0.0"
		}
		ipInfo, err := region.MemorySearch(ip)
		if err != nil {
			fmt.Fprintf(w, "[]")
		} else {
			b, err := json.Marshal(ipInfo)
			if err != nil {
				b = []byte("[]")
			}
			fmt.Fprintf(w, string(b))
		}

	})
	http.ListenAndServe(":8082", nil)
}
