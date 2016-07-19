package main

import (
	"fmt"
	"time"
)

func main() {
	// create new channel of type int
	chStart := make(chan int, 5)
	for i := 0 ; i < 5 ; i++{
		chStart <-i
	}

	time.Sleep(2 * time.Second)

	for i := 0 ; i < 20 ; i++ {
		go printSmth(i, chStart)
	}

	for ;;{
		switch  {
		}
	}
}

func printSmth(tNum int, chStart chan int){
	<-chStart
	time.Sleep(5*time.Second)
	fmt.Println("ok from ", tNum)
	chStart<-1
}