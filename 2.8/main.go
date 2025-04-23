package main

func main() {
	ch := make(chan bool, 1) // чтобы решить проблему, можно использовать буфер
	ch <- true
	go func() {
		<-ch
	}()
	ch <- true
}
