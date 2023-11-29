package main

import (
	"fmt"
	"time"
)

// T1
func main() {
	ch := make(chan int)

	//T2
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("Enviando no canal", i)
		}
	}()
	//T3
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("Enviando no canal", i)
		}
	}()

	//T4
	go worker(ch, 1)
	go worker(ch, 2)
	worker(ch, 3)

}

func worker(ch chan int, workerID int) {
	for x := range ch {
		fmt.Println("Recebendo no canal", x, " worker ", workerID)
		time.Sleep(time.Second)
	}
}
