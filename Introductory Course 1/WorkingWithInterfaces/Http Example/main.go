package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.twitter.com")
	if err != nil {
		fmt.Println("Error is : ", err)
		os.Exit(1)
	}

	// fmt.Println("Data present in the the byte slice from the response", resp, err)
	// getting data from io.Read
	// dataTobeRead := make([]byte, 999999)
	// resp.Body.Read(dataTobeRead)

	// fmt.Println(string(dataTobeRead))

	// Code to simplfy the process of reading data from response

	lw := logWriter{}
	fmt.Println("Data to be evaluated")
	// fmt.Println(resp.Body)
	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Data to be proccessed")
	fmt.Println("Data received ", string(bs))
	fmt.Println("Data written is ", len(bs))
	return len(bs), nil
}
