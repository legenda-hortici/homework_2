package main

import (
	"context"
	"fmt"
	"sync"
)

func merge(ch ...chan int) <-chan int {
	// создаем результирующий канал
	result := make(chan int)

	var wg sync.WaitGroup
	// используем контекст для отмены операций, если один из каналов будет закрыт
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем горутину для каждого канала
	output := func(ch <-chan int) {
		defer wg.Done() // уменьшаем счетчик горутин

		for {
			select {
			case <-ctx.Done(): // если был сигнал отмены, то прерываем контекст
				return
			case v, ok := <-ch: // если канал закрыт, то закрываем каналы
				if !ok {
					cancel()
					return
				}
				select {
				case result <- v: // записываем значение в результирующий канал
				case <-ctx.Done(): // если был сигнал отмены, то завершаем контекст
					return
				}
			}
		}
	}

	// запускаем горутины для каждого канала
	for _, c := range ch {
		wg.Add(1)
		go output(c)
	}

	// ожидаем завершения всех горутин и закрываем результирующий канал
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	// создаем 2 канала
	ch1, ch2 := make(chan int), make(chan int)

	// заполняем первый канал числами
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// заполняем второй канал числами
	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	// объединяем каналы
	mergedChan := merge(ch1, ch2)

	// читаем объединенный канал
	for v := range mergedChan {
		fmt.Println(v)
	}
}
