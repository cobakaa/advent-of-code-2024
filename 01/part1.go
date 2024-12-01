package main

import (
	"fmt"
	"io"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func scan() ([]int, []int) {
	a := make([]int, 0)
	b := make([]int, 0)

	for {
		var t1, t2 int
		_, err := fmt.Scanf("%d %d", &t1, &t2)
		a = append(a, t1)
		b = append(b, t2)

		if err == io.EOF {
			break
		}
	}

	return a, b
}

func part1(a, b []int) {

	sort.Ints(a)
	sort.Ints(b)

	sum := 0
	for i := 0; i < len(a); i++ {
		sum += abs(a[i] - b[i])
	}

	fmt.Println(sum)
}

func part2(a, b []int) {
	sum := 0
	
	m := make(map[int]int)
	for i := 0; i < len(b); i++ {
		m[b[i]]++
	}

	for i := 0; i < len(a); i++ {
		sum += a[i] * m[a[i]]
	}

	fmt.Println(sum)
}

func main() {
	a, b := scan()
	part1(a, b)
	part2(a, b)
}