package main

func main() {
	ch := make(chan bool, 1) // создание канала с bool значением
	ch <- true               // пишем в канал true
	go func() {              // горутина читает
		<-ch
	}()
	ch <- true // пишем в каналл
	close(ch)
}

// применила буферизвцию канала и закрыла канал после всех операций
