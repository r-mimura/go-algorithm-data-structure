// https://atcoder.jp/contests/abc081/tasks/abc081_b

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func min(nums ...int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

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
	ans := INF
	for i := 0; i < N; i++ {
		cnt := 0
		for a[i]%2 == 0 {
			a[i] /= 2
			cnt++
		}
		ans = min(ans, cnt)
	}
	fmt.Fprintln(w, ans)

}
