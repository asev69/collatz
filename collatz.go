package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n uint
	fmt.Scan(&n)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, n)
	for n != 1 {
		if n % 2 == 0 {
			n /= 2
		} else {
			n = n * 3 + 1
		}
		fmt.Fprint(w, " ", n)
	}
	fmt.Fprintln(w)
	w.Flush()
}
