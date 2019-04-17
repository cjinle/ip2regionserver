package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

type AppConf struct {
	Listen string `json:listen`
	Mode   string `json:mode`
}

func main() {
	region, err := ip2region.New("conf/ip2region.db")
	if err != nil {
		panic("ip2region.db err")
	}
	defer region.Close()

	s, err := ioutil.ReadFile("conf/app.json")
	if err != nil {
		panic("app.json err")
	}
	cfg := AppConf{}
	err = json.Unmarshal(s, &cfg)
	if err != nil {
		panic("Unmarshal app.json err")
	}

	r := gin.Default()
	if cfg.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	r.GET("/query", func(c *gin.Context) {
		ip := c.DefaultQuery("ip", "0.0.0.0")
		ipInfo, err := region.MemorySearch(ip)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, []string{})
		} else {
			c.JSON(http.StatusOK, ipInfo)
		}
	})
	r.Run(cfg.Listen)
}
