package main

import (
	"fmt"
	"time"
)

func goWorker(workerId int, inputChannel chan string, outputChannel chan string) {
	for inp := range inputChannel {
		fmt.Printf("Worker Id: %d is processing input: %s\n", workerId, inp)
		time.Sleep(1 * time.Second)
		outputChannel <- fmt.Sprintf("Pushed by worker Id: %d, Processed input %s\n", workerId, inp)
	}
}

func main() {
	inputChannel, outputChannel := make(chan string, 3), make(chan string, 3)

	for i := 0; i < 3; i++ {
		go goWorker(i+1, inputChannel, outputChannel)
	}

	for _, val := range []string{"SB", "BS", "MacBook", "Company", "Apple"} {
		inputChannel <- val
	}

	for i := 0; i < 5; i++ {
		select {
		case result := <-outputChannel:
			fmt.Println("RESULT: ", result)
		case <-time.After(1 * time.Second):
			fmt.Println("Timed out !!!")
		}
	}
}
