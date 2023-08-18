package main

import (
	"fmt"
	"time"
)

func sum(nums []int, worker chan int) {
	fmt.Print("Running sum!!")
	time.Sleep(time.Millisecond * 1000)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	worker <- sum
}

func createArray(init int, end int) []int {
	array := make([]int, end-init+1)
	for num := init; num <= end; num++ {
		array = append(array, num)
	}
	return array
}

func main() {
	worker := make(chan int)

	go sum(createArray(1, 100), worker)
	go sum(createArray(1, 10000), worker)

	sum1 := <-worker
	sum2 := <-worker

	fmt.Println(sum1, sum2)
}
