// https://atcoder.jp/contests/abc051/tasks/abc051_b

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//INF : 十分大きな値
const INF = 1000000009

func Max(nums ...int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var K, S int
	fmt.Fscan(r, &K, &S)
	ans := 0
	for i := 0; i <= K; i++ {
		for j := 0; j <= K; j++ {
			if S-i-j>=0 && S-i-j<=K{
				ans++
			}
		}
	}
	fmt.Fprintln(w, ans)
}
