package main

import "fmt"

func main() {
	c := make(map[string]int)
	c["a"] = 123
	for i := 0; i < 1000000000; i++ {
		go func() {
			for j := 0; j < 100000000; j++ {
				//c[fmt.Sprintf("%d", j)] = j
				fmt.Println(c["a"])
			}
		}()
	}
}
