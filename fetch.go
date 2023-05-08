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
		if !strings.HasPrefix(url, "https") && !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}
		fmt.Println(url)

		resp, err := http.Get(url)
		fmt.Println(resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//b, err := io.ReadAll(resp.Body)
		file, err := os.Create("output.html")
		if err != nil {
			panic(err)
		}

		b, err := io.Copy(file, resp.Body)
		resp.Body.Close()
		file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

	}
}
