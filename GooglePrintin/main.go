package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

type logWriter struct {

}

func main() {
	fmt.Print("Working")

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(resp)
	bs := make([]byte, 9999)
	resp.Body.Read(bs)

	lw := logWriter{}

	fmt.Println(string(bs))

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte)(int, error) {
	return len(bs), nil
}
