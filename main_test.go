package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	html := `<DOCTYPE html><html></html>`
	target, _ := os.Open(os.DevNull)
	defer target.Close()

	for i := 0; i < b.N; i++ {
		content := ioutil.NopCloser(bytes.NewReader([]byte(html)))
		processAndWriteToConsole(content, target)
	}
}
