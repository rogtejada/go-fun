//fetchs a given url and write output to stdout

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const (
		prefix = "http://"
	)

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprint(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprint(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

	}
}
