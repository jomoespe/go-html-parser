/*
 * https://godoc.org/golang.org/x/net/html#Parse
 */
package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func walk(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "div" {
		for _, a := range node.Attr {
			if a.Key == "data-loc" {
				println("encontrado uno con data-loc", a.Val)
				break
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		walk(child)
	}
}

func processAndWriteToConsole(content io.ReadCloser, target io.Writer) {
	root, err := html.Parse(content) // parse and get a tree of nodes
	if err != nil {
		panic(err)
	}
	walk(root)
	err = html.Render(target, root) // render the node tree to STDOUT
}

func main() {
	content, err := fromFile("file.html")
	//content, err := fromUrl("http://www.google.es")
	if err != nil {
		panic(err)
	}
	defer content.Close()

	target, _ := os.Open(os.DevNull)
	defer content.Close()

	processAndWriteToConsole(content, target)
}

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
