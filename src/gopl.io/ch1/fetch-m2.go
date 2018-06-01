package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(strings.ToLower(url), "http://") {
			fmt.Printf("url: %s\n\n", url)
			url = "http://" + url
		}
		fmt.Printf("url: %s\n\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		n, err := io.Copy(os.Stdout, resp.Body)
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("\n\n%d bytes copied.\n", n)
	}
}