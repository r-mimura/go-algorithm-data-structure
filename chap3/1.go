package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, v int
	fmt.Fscan(r, &N, &v)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &a[i])
	}
	var foundID = -1
	for i := 0; i < N; i++ {
		if a[i] == v {
			foundID = i
		}
	}
	fmt.Fprintln(w, foundID)

}
