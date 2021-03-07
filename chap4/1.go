package main

import (
	"bufio"
	"fmt"
	"os"
)

func tri(n int) int {
	if n == 0 || n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	return tri(n-1) + tri(n-2) + tri(n-3)

}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	fmt.Fprintln(w, tri(n))

}
