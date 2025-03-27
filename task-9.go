package main

import (
	"fmt"
	"sync"
)

func merge(channels ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	mergeChannel := func(ch chan int) { // читаем из каждого канала
		defer wg.Done() // уменьшение счетчика горутин
		for val := range ch {
			out <- val //  пишем в канал
		}
	}

	for _, ch := range channels { // цикл по горутинам
		wg.Add(1) // счетчик горутин
		go mergeChannel(ch)
	}

	// Запускаем горутину для закрытия выходного канала после завершения всех
	go func() { // горутины
		wg.Wait()  // ожидание горутин
		close(out) // закрываем выходной канал
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	merged := merge(ch1, ch2)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1) // горутина, данные отправили, канал закрытли
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2) // горутина, данные еще отправили, канал закрыли
	}()

	for val := range merged {
		fmt.Println(val) // печать данных из единого канала
	}
}
