package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	m := [][]int{}
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		t := strings.Fields(strings.TrimSpace(s))
		r := []int{}
		for _, i := range t {
			var x int
			fmt.Sscan(i, &x)
			r = append(r, x)
		}
		m = append(m, r)
	}
	n := len(m)
	ps := [][]int{}
	for i := 0; i < n; i++ {
		r := make([]int, n)
		ps = append(ps, r)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ps[j][i] = m[j][i]
			if j > 0 {
				ps[j][i] += ps[j-1][i]
			}
		}
	}
	var maxmat int
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			var minmat, submat int
			for k := 0; k < n; k++ {
				submat += ps[j][k]
				if i > 0 {
					submat -= ps[i-1][k]
				}
				if submat < minmat {
					minmat = submat
				}
				if submat - minmat > maxmat {
					maxmat = submat - minmat
				}
			}
		}
	}
	fmt.Println(maxmat)
}
