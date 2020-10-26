package ip2regionserver

import (
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

var region *ip2region.Ip2Region

func init() {
	var err error
	region, err = ip2region.New("../conf/ip2region.db")
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
