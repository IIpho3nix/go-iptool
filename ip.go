package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"io/ioutil"
    "net/http"
)

func main() {
	local := flag.Bool("local", false, "Get local IP address")
	public := flag.Bool("public", false, "Get public IP address")
	flag.BoolVar(local, "l", false, "Get local IP address")
	flag.BoolVar(public, "p", false, "Get public IP address")

	flag.Parse()

	if !*local && !*public {
		*public = true
	}

	if *local && *public {
		flag.Usage()
		os.Exit(1)
	}

	if *local {
		localIP := getLocalIP()
		fmt.Println(localIP)
	} else {
		publicIP := getPublicIP()
		fmt.Println(publicIP)
	}
}

func getPublicIP() string {
    resp, err := http.Get("https://api.ipify.org")
    if err != nil {
        return ""
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return ""
    }
    return string(body)
}

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}