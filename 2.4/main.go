package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		При записи в неинициализированную map будет паника.
		Чтобы изюежать панику, можно использовать sync.Map
	*/

	var m sync.Map
	i := 0
	for _, word := range []string{"hello", "world", "from", "the", "best", "language", "in", "the", "world"} {
		i++
		m.Store(word, i)
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
