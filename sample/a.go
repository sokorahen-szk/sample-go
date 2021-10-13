//パスカルの三角形

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)

	c := 0
	ac := make([]int, 0)
	an := make([]int, 0)
	ac = append(ac, 1)
	for {
		for i := 0; i <= c; i++ {
			if c == 0 {
				break
			}
			if i == 0 {
				ac = append(ac, an[i])
				continue
			} else if i == c {
				ac = append(ac, an[i-1])
				continue
			}
			ac = append(ac, an[i-1]+an[i])
		}
		c++

		fmt.Println(ac)

		an = ac
		ac = make([]int, 0)

		if n == c {
			break
		}
	}

}
