package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	data := os.Args
	fmt.Print(data)
	f, err := os.OpenFile(data[1], os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Error while reading the input file ", data[1])
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

}
