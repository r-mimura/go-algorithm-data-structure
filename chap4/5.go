// https://atcoder.jp/contests/abc114/tasks/abc114_c

package main

import (
	"bufio"
	"fmt"
	"os"
)

//INF : 十分大きな値
const INF = 1000000009

var ans, n int

func dfs(x int, three, five, seven bool) {
	if x > n {
		return
	}
	if three && five && seven {
		ans++
	}
	dfs(10*x+3, true, five, seven)
	dfs(10*x+5, three, true, seven)
	dfs(10*x+7, three, five, true)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n)
	dfs(0, false, false, false)
	fmt.Fprintln(w, ans)

}
