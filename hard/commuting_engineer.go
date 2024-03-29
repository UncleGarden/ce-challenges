package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type place struct {
	id int
	x  float64
	y  float64
}

func dist(a, b *place) float64 {
	x, y := a.x-b.x, a.y-b.y
	return math.Sqrt(x*x + y*y)
}

func tour(nodes []place, di []int) (ret float64) {
	for i := 1; i < len(di); i++ {
		ret += dist(&nodes[di[i-1]], &nodes[di[i]])
	}
	return ret
}

func find_path(o, p []int, d, mx float64, n int, nodes []place) ([]int, float64) {
	var o2, p2 []int
	if len(p) == 0 {
		p2 = []int{o[0]}
	} else {
		p2 = make([]int, len(p))
		copy(p2, p)
	}
	o2 = make([]int, len(o))
	copy(o2, o)
	route := []int{}
	for i := 0; i < len(o2); i++ {
		found := false
		for j := 0; j < len(p2); j++ {
			if p2[j] == o2[i] {
				found = true
				break
			}
		}
		if !found {
			route = append(route, o2[i])
		}
	}
	if len(route) == 1 {
		return append(p2, route[0]), d + dist(&nodes[p2[len(p2)-1]], &nodes[route[0]])
	}
	m2, mp := mx, o
	for i := 0; i < len(route); i++ {
		var mindi, maxdi float64
		for j := 0; j < len(route); j++ {
			var mindis1 float64
			if j == i {
				mindis1 = math.MaxFloat64
			} else {
				mindis1 = dist(&nodes[route[i]], &nodes[route[j]])
			}
			mindis2 := math.MaxFloat64
			for k := 0; k < len(route); k++ {
				if j == k {
					continue
				}
				dis := dist(&nodes[route[j]], &nodes[route[k]])
				if dis < mindis2 {
					if dis < mindis1 {
						mindis1, mindis2 = dis, mindis1
					} else {
						mindis2 = dis
					}
				}
			}
			if j == i {
				mindi += mindis1 / 2
			} else {
				mindi += (mindis1 + mindis2) / 2
				if maxdi < mindis2 {
					maxdi = mindis2
				}
			}
		}
		mindi -= maxdi / 2

		d2 := d + dist(&nodes[p2[len(p2)-1]], &nodes[route[i]])
		if d2+mindi < m2 {
			di, du := find_path(o2, append(p2, route[i]), d2, m2, n, nodes)
			if du < m2 {
				m2 = du
				mp = di
			}
		}
	}
	return mp, m2
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	places, path := []place{}, []int{}
	reader := bufio.NewReader(data)
	for i := 0; ; i++ {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		t := strings.Fields(s)
		var (
			id   int
			x, y float64
		)
		fmt.Sscan(t[0], &id)
		fmt.Sscanf(t[len(t)-2], "(%f,", &x)
		fmt.Sscanf(t[len(t)-1], "%f)", &y)
		places = append(places, place{id, x, y})
		path = append(path, i)
	}
	shortest_path, _ := find_path(path, []int{}, 0, tour(places, path), len(path), places)
	for _, i := range shortest_path {
		fmt.Println(places[i].id)
	}
}
