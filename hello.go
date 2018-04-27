package main

import (
	"fmt"
	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

func main() {
	priorities := []float32{77.7, 22.2, 44.4, 55.5, 11.1, 88.8, 33.3, 99.9, 0.0, 66.6}
	data := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	pq := prque.New()
	for i := 0; i < len(data); i++ {
		pq.Push(data[i], priorities[i])
	}
	for !pq.Empty() {
		val, prio := pq.Pop()
		fmt.Printf("%.1f:%s\n", prio, val)
	}
}
