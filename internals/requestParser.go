package internals

import (
	"bytes"
	"errors"
)
type requestLine struct {
	method string
	url string
	httpVersion string
}


var CRLF = []byte("\r\n")

// func httpParser(data []byte) {
//
// }

func parseRequestLine(data []byte, done bool) (*requestLine, bool, error) {

	// Because this line is fixed... we can proceed with this logic i guess.... 
	// if we don't get this line in the first try... we are going to return error that it's not a http request...

	if done {
		return nil, true, nil
	}

	if i := bytes.Index(data, CRLF); i != -1 {
		line := data[:i]	
		 
		parts := bytes.Split(line, []byte(" ")) 

		if len(parts) != 3 {
			return nil, false, errors.New("INVALID HTTP REQUEST.. CLIENT SEND THE CORRECT HTTP PACKET AFTER VARIFYING IT WITH YOUR DADDY") 
		}


		return &requestLine{
			method: string(parts[0]), 
			url: string(parts[1]), 
			httpVersion: string(parts[2]), 
		}, true, nil
	}


	return nil, false, errors.New("THIS IS NOT A VALID HTTP REQUEST, SINCE \r\n IS NOT PRESENT IN THE FIRST 1024 Bytes FOR THE REQUESTLINE") 
}

func parseHeaderFields(data []byte) {
}

func parseBody(data []byte) {
}
