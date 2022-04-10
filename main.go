package main

import "time"

func main() {
	now := time.Now()
	for i := 0; i < 100; i++ {
		println(i)
	}
	defer func() {
		println(time.Since(now))
	}()
}
