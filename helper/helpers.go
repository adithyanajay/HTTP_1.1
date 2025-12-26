package helper

import (
	"fmt"
	"os"
	"net"
	"http.adithyaajay.duck/globals"
	"strings"
	"strconv"
)

func GetAddr(listener net.Listener) (*globals.TcpServerAddr) {
	
	protocol := listener.Addr().Network()
	addr := listener.Addr().String()
	data := strings.Split(addr, ":")

	if len(data) != 2 {
		fmt.Printf("Error: server IP and PORT not detecting... %v", data)
		os.Exit(1)
	}

	port, err := strconv.Atoi(data[1]) 

	if err != nil {
		fmt.Printf("Error: server IP and PORT not detecting... %v", data)
		os.Exit(1)
	}
		
	return &globals.TcpServerAddr {
		Protocol: protocol, 
		IpAddr: data[0],
		PORT : port,
	}
}



