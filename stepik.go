package main

import (
	"fmt"
	"sort"
)

func compression(rectan []int) {
	cnt := 0
	cnt_1 := 0
	arr_x := make([]int, len(rectan)/2)
	arr_y := make([]int, len(rectan)/2)
	for i := 0; i < len(rectan); i++ {
		if i%2 == 0 {
			arr_x[cnt] = rectan[i]
			cnt++
		} else {
			arr_y[cnt_1] = rectan[i]
			cnt_1++
		}
	}
	sort.Ints(arr_x)
	sort.Ints(arr_y)
	fmt.Println(arr_x, arr_y)
	fmt.Print(rectan)
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	compression(list)
	return list
}

func main() {
	var n int
	fmt.Scan(&n)
	rectan := make([]int, n*4)
	for i := 0; i < n*4; i++ {
		fmt.Scan(&rectan[i])
	}
	removeDuplicateInt(rectan)
}
