package globals 

//Global Port variable for the TCP Server
type TcpServerAddr struct {
	Protocol string
	IpAddr string
	PORT int
}


// Struct to store the client response
type clientResponse struct {
	tcpStream []byte 
}


