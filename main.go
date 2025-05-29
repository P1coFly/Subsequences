package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ALPHA = 26
	NONE  = "NONE"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minAlphabetWindowLength(seq []int) int {
	count := make([]int, ALPHA+1)
	have := 0
	left := 0
	best := len(seq) + 1

	for right, v := range seq {
		if count[v] == 0 {
			have++
		}
		count[v]++

		for have == ALPHA {
			windowLen := right - left + 1
			best = min(best, windowLen)

			w := seq[left]
			count[w]--
			if count[w] == 0 {
				have--
			}
			left++
		}
	}

	if best == len(seq)+1 {
		return 0
	}
	return best
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	seq := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &seq[i])
	}

	res := minAlphabetWindowLength(seq)
	if res == 0 {
		fmt.Fprint(out, NONE)
	} else {
		fmt.Fprint(out, res)
	}
}
