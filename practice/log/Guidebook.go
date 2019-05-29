package main

import (
	"fmt"
	"sort"
)

type list struct {
	c string
	s int
	o int
}
type Lists []list

func (l Lists) Len() int {
	return len(l)
}

func (l Lists) Less(i, j int) bool {
	if l[i].c != l[j].c {
		return l[i].c < l[j].c
	} else {
		return l[i].s > l[j].s
	}
}

func (l Lists) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	var (
		N int
	)
	fmt.Scan(&N)
	lists := make([]list, N)
	for i := 0; i < N; i++ {
		var (
			city string
			s    int
		)
		fmt.Scan(&city, &s)
		lists[i] = list{
			c: city,
			s: s,
			o: i + 1,
		}
	}

	sort.Sort(Lists(lists))
	for _, list := range lists {
		fmt.Println(list.o)
	}
}
