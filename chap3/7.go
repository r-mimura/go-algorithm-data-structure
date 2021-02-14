// https://atcoder.jp/contests/abc045/tasks/arc061_a

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

	var S string
	fmt.Fscan(r, &S)
	ans := 0
	for bit := 0; bit < 1<<(len(S)-1); bit++ {
		cnt := 0
		for i := 0; i < len(S); i++ {
			cnt = cnt*10 + int(S[i]) - '0'
			if (bit>>i)&1 == 1 {
				ans += cnt
				cnt = 0
			}
		}
		ans += cnt
	}
	fmt.Fprintln(w, ans)

}
