package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gogorepos/goskeleton/snmp"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ip := request.Form.Get("ip")
		if ip == "" {
			ip = "169.254.0.5"
		}
		_, _ = writer.Write(snmp.Run("169.254.0.5"))
	})
	http.HandleFunc("/speed", func(writer http.ResponseWriter, request *http.Request) {
		ip := request.Form.Get("ip")
		if ip == "" {
			ip = "169.254.0.5"
		}
		_, _ = writer.Write(snmp.RunWithSpeed(""))
	})
	log.Println("Listen in 0.0.0.0:9899")
	log.Fatalln(http.ListenAndServe(":9899", nil))
}

var lastSpeed []int
var speed = make([]int, 32)

func getSpeed() {
	s, _ := snmp.NewSNMP("169.254.0.5")
	r, _ := s.Walk(snmp.IfInOctetOid)
	for _, result := range r {
		lastSpeed = append(lastSpeed, result.Int())
	}
	t := time.NewTicker(time.Second)
	for {
		<-t.C
		r, _ = s.Walk(snmp.IfInOctetOid)
		for i, result := range r {
			speed[i] = result.Int() - lastSpeed[i]
		}
		log.Println(speed)
	}
}
