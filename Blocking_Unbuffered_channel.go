package main

import (
	"fmt"
)

func f(ball chan string, done chan bool) {
	fmt.Println(<- ball)
	done <- true
}

func main() {
	ball := make(chan string)
	done := make(chan bool)
	
  go f(ball, done)
  
	ball <- "Shubham"
	<-done
}
