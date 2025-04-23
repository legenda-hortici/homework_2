package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3) // чтобы избежать дедлока, можно использовать буфер
	wg := &sync.WaitGroup{}

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}

	/*
		но если не известно сколько горутин запущено, то можно использовать
		горутину для закрытия канала и дождаться завершения всех горутин
	*/
	go func() {
		wg.Wait()
		close(ch)
	}()

	var sum int
	for v := range ch {
		sum += v
	}

	fmt.Printf("result: %d\n", sum)
}
