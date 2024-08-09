package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// Here we are making a get request to google.com
	resp, err := http.Get("https://google.com")

	//Here we are checking if there is an error
	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}
	// //Here we are creating a byte slice
	// bs := make([]byte, 99999)
	// //Here we are reading the body of the response
	// resp.Body.Read(bs)

	// //Here we are printing the body of the response and converting byte slice to string
	// fmt.Println(string(bs))

	//HEre we are using the io.Copy function to copy the body of the response to the standard output initialy doing the same as above

	//Basically first argument has to be implement writer interface, second argument should be reader
	//If looked into documentation tree of stdout, we would see that it writes at some point, so it can be used as writer
	//As for reader, resp.Body directly implements reader interface
	
	//This can be used as writer now, so basically it can break code if used incorrecly if return just be random
	// lw := logWriter{} 
	
	// io.Copy(lw, resp.Body)
		io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
