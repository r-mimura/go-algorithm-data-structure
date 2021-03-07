package main

import (
	"bufio"
	"fmt"
	"os"
)

var memo [50]int

func tri(n int) int {
	if n == 0 || n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	if memo[n] != -1 {
		return memo[n]
	}
	memo[n] = tri(n-1) + tri(n-2) + tri(n-3)
	return tri(n-1) + tri(n-2) + tri(n-3)

}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < 50; i++ {
		memo[i] = -1
	}
	var n int
	fmt.Fscan(r, &n)

	fmt.Fprintln(w, tri(n))

}
