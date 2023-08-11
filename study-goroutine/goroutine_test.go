package study_goroutine

import (
	"fmt"
	"testing"
	"time"
)



func HelloWorld() {
	fmt.Println("Hello World")
}



func TestGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}



func DisplayNumber(number int) {
	fmt.Println(number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

