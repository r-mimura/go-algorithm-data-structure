package main

import (
	"bufio"
	"fmt"
	"os"
)

//INF : 十分大きな値
const INF = 1000000009

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	fmt.Fscan(r, &N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &a[i])
	}

	minVal := INF
	maxVal := -INF
	for i := 0; i < N; i++ {
		if a[i] < minVal {
			minVal = a[i]
		}
		if a[i] > maxVal {
			maxVal = a[i]
		}
	}
	fmt.Fprintln(w, maxVal-minVal)

}