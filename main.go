package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"io/ioutil"
	"fmt"
	"net"
	"bytes"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "PONG"+"\n")
	})
	r.GET("/info", func(c *gin.Context) {
		serviceName := os.Getenv("SERVICE_NAME")
		hostname, _ := os.Hostname()
		mac := getMacAddr()

		data, err := ioutil.ReadFile("/tmp/conf")
		if err != nil {
			fmt.Println(err.Error())
		}
		resStr := serviceName + " : " + hostname + " : " + mac + "\n" + string(data)
		c.String(200, strings.ToUpper(resStr))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	nic := os.Getenv("NIC")
	if nic == "" {
		nic = "0.0.0.0"
	}

	fmt.Println("Goona use follinwg port: " + port)
	fmt.Println("Goona use follinwg nic: " + nic)
	r.Run(nic + ":" + port)
}

func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}
