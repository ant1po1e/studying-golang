package study_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Antipole"
		fmt.Println("Sended data to Channel")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second	)
}