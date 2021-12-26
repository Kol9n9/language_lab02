package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"bytes"
)

func DownloadInfo(EOF bool, buffer *bytes.Buffer){
	for !EOF{
		fmt.Println(buffer.Len() / 1024, "Kb")
        time.Sleep(time.Second)
	}
}
func main() {
	fmt.Print("Enter URL:")

	var URL string
	if _, err := fmt.Scanln(&URL); err != nil {
		panic(err)
	}

	resp, err := http.Get(URL)
	if  err != nil {
		println("Error:", err.Error())
		os.Exit(-1)
	} else if resp.StatusCode != http.StatusOK{
		println("Error:", resp.Status)
		os.Exit(-2)
	}
	defer resp.Body.Close()

	fileName := URL[strings.LastIndex(URL, "/")+1 : len(URL)]
	file, err := os.Create(fileName)
	if err != nil {
		println("Error:", err.Error())
		os.Exit(-3)
	}
	defer file.Close()

	var buffer bytes.Buffer
	EOF := false
	tee := io.TeeReader(resp.Body, &buffer)
	fmt.Println("Start downloading. Already received:")
	go DownloadInfo(EOF, &buffer)
	io.Copy(file, tee)
	EOF = true
	fmt.Println("Download")
}
