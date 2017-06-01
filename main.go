/*
 * https://godoc.org/golang.org/x/net/html#Parse
 */
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func fromFile(filename string) (content io.ReadCloser, err error) {
	content, err = os.Open(filename)
	return 
}

func fromUrl(url string) (content io.ReadCloser, err error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 
	}
	
	req.Header.Add("X-MyHeader", "My Value")

	resp, err := client.Do(req)
	if err != nil {
		return 
	}
	
	content = resp.Body 
	return 
}

func processAndWriteToConsole(content io.ReadCloser) {
	root, err := html.Parse(content) // parde and get a tree of nodes
	if err != nil {
		panic(err)
	}

	buff := bytes.NewBufferString("")
	err = html.Render(buff, root) // render the node tree to a buffer
	if err != nil {
		panic(err)
	}

	fmt.Println(buff.String()) // here we only print the "renderized" tree
}

func main() {
	content, err := fromFile("file.html")
	//content, err := fromUrl("http://www.google.es")
	if err != nil {
		panic(err)
	}
	defer content.Close()
	processAndWriteToConsole(content)
}
