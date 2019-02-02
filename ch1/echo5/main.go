/*
Echo 5 shows how much time each version of our echo method
will take
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
