// Fetch Excercise 1.7: Uses Copy rather than ReadAll
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// This is the main function
// It sings and dances.
// Sometimes does tricks on the side.
func main() {
	for _, url := range os.Args[1:] {
		// Exercise 1.8: Add scheme if not already there.
		theURL := "http://"
		if !strings.HasPrefix(url, "http://") {
			theURL += url
		}
		resp, err := http.Get(theURL)
		statusCode := resp.StatusCode
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("The Status Code: %d\n", statusCode)
		fmt.Printf("The Number of Bytes Written: %d\n", b)
	}
}
