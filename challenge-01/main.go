// challenge 01
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x, k int
	fmt.Scanf("%d %d %d", &n, &x, &k)

	dls  := make([]int, n, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&dls[i])
	}

	keys := make([]int, k)
	for j:=0; j<k; j++ {
		keys[j] = 9223372036854775807
	}

	msg := make(map[int]struct{})

	for i:=0; i<n; i++ {
		for j:=0; j<k; j++ {
			v := dls[i] + j*x
			if _, ok := msg[v]; !ok {
				msg[v] = struct{}{}
				if keys[k-1] > v {
					keys[k-1] = v
					sort.Ints(keys)
				}	else {
					break
				}
			}
		}
	}

	fmt.Printf("%d", keys[k-1])
}