package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

type Stat struct {
	Succ int
	Fail int
}

func main() {
	fmt.Println("http server")
	region, err := ip2region.New("conf/ip2region.db")
	if err != nil {
		panic("ip2region.db err")
	}
	defer region.Close()
	stat := Stat{0, 0}

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		if ip == "" {
			ip = "0.0.0.0"
		}
		ipInfo, err := region.MemorySearch(ip)
		if err != nil {
			fmt.Fprintf(w, "[]")
			stat.Fail = stat.Fail + 1
		} else {
			b, err := json.Marshal(ipInfo)
			if err != nil {
				b = []byte("[]")
			}
			stat.Succ = stat.Succ + 1
			fmt.Fprintf(w, string(b))
		}

	})
	http.HandleFunc("/stat", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(stat)
		fmt.Fprintf(w, string(b))
	})
	http.ListenAndServe(":8080", nil)
}
