package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	f, err := os.Open("testfiles/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	root, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	//	ctx := context.Background()
	var timings []string
	ctx := context.WithValue(context.Background(), timings, &timings)

	var wg sync.WaitGroup
	processNode(ctx, &wg, root)
	wg.Wait()
	v, _ := ctx.Value("timings").([]string)
	fmt.Printf("%v\n", v)
}

func processNode(ctx context.Context, wg *sync.WaitGroup, node *html.Node) html.Node {
	if node.Type == html.ElementNode && strings.EqualFold(node.Data, "article") {
		wg.Add(1)
		go processArticle(ctx, wg, node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		processNode(ctx, wg, c)
	}
	return *node
}

func processArticle(ctx context.Context, wg *sync.WaitGroup, node *html.Node) {
	defer wg.Done()

	id := node.Attr[0].Val
	v, _ := ctx.Value("timings").([]string)
	v = append(v, fmt.Sprintf("%s, ", id))

	fmt.Printf("node.Data: %s DONE!\n", id)
}
