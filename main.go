package main

import (
	"fmt"
	"time"
)


func main() {
	var cache1 = newCache()
	var mode, value, TTL int
	var key string
	go checkExpiry(&cache1)
	for {
		fmt.Println(time.Now().Unix())
		fmt.Println("Enter mode:\n 1. Set Value with default TTL\n 2. Set Value with TTL \n 3. Get Value \n 4. Delete Value \n 5. Print Map \n 6. Quit")
		fmt.Scanln(&mode)
		if mode == 1 {
			fmt.Print("Enter key, value: ")
			fmt.Scan(&key, &value)
			cache1.set(Parameters{data: int64(value), key: key})
		} else if mode == 2 {
			fmt.Print("Enter key, value, TTL: ")
			fmt.Scan(&key, &value, &TTL)
			cache1.set(Parameters{data: int64(value), key: key, TTL: int64(TTL)})
		} else if mode == 3 {
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			fmt.Println(cache1.get(key))
		} else if mode == 4 {
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			cache1.delete(key)
		} else if mode == 5 {
			fmt.Println(cache1.Map)
		} else if mode == 6 {
			break
		}
	}
}
