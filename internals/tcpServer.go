package internals 

import (
	"fmt"
	"net"
	"io"
	"os"
	// "http.adithyaajay.duck/globals"
	"http.adithyaajay.duck/helper"
)

// Opening a TCP connection 
func OpenServerConn() {
	// var tcpAddr globals.tcpServerAddr{}

	fmt.Println("Server launching ")

	listner, err := net.Listen("tcp", "localhost:")

	serverDetails := helper.GetAddr(listner)	

	if err != nil {
		fmt.Errorf("Error: Listening the tcp server %w", err)	
		os.Exit(1)
	}

	defer listner.Close()

	fmt.Printf("%s Server %s started in PORT:%d\n",serverDetails.Protocol, serverDetails.IpAddr, serverDetails.PORT)
	acceptClientsRequest(listner)
	
}

// Accepting connection and reading bytes from clients 
func acceptClientsRequest(listener net.Listener) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Errorf("Error: Server not accepting Clients... %w", err)
			os.Exit(1)
		}

		go func() {
			defer conn.Close()
			buf :=  make([]byte,1024)
			// data := make([]byte, 0, 4096) //Max 4kb
			dataSize := 0
			done := false
			var requestLine *requestLine	
			for {
				if dataSize > 4096 {
					// TODO make a response with packet exceeds limit
					fmt.Println("Packet exceeded the limit", dataSize)
					return
				}

				n, err := conn.Read(buf)

				if err != nil {
					if err == io.EOF {
						fmt.Println("The client disconnected: ", conn.RemoteAddr())
						return
					}

					fmt.Errorf("Error: Malfunction reading the packet... %w", err)
					os.Exit(1)
				}

				dataSize += n

				// fmt.Printf("%q \n", buf[:n])

				// parserData, err := httpParser(buf)
				requestLine, done, err = parseRequestLine(buf, done)

				if err != nil {
					//TODO Send a response to the client that the http request is not valid
					return
				}

				fmt.Println()
				fmt.Printf("METHOD: %s\n", requestLine.method)
				fmt.Printf("URL: %s\n", requestLine.url)
				fmt.Printf("VERSION: %s\n", requestLine.httpVersion)


				// Send back the response to the Client 
				// generateResponse(parserData)

			}
		}()


	}

}



