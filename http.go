package ip2regionserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AppConf struct {
	Listen string `json:listen`
	Mode   string `json:mode`
}

func HttpListen() {
	defer region.Close()

	s, err := ioutil.ReadFile("../conf/app.json")
	CheckError(err)

	cfg := AppConf{}
	err = json.Unmarshal(s, &cfg)
	CheckError(err)

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

	http.ListenAndServe(cfg.Listen, nil)
}

func GinListen() {
	defer region.Close()

	s, err := ioutil.ReadFile("../conf/app.json")
	CheckError(err)

	cfg := AppConf{}
	err = json.Unmarshal(s, &cfg)
	CheckError(err)

	r := gin.Default()
	if cfg.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	r.GET("/query", func(c *gin.Context) {
		ip := c.DefaultQuery("ip", "0.0.0.0")
		ipInfo, err := region.MemorySearch(ip)
		if err != nil {
			c.JSON(http.StatusOK, []string{})
		} else {
			c.JSON(http.StatusOK, ipInfo)
		}
	})
	r.Run(cfg.Listen)
}
