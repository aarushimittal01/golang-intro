package channels

import (
	"fmt"
	"os"
	"time"
)

//print hello and world alternatively with 1 sec delay after hello and 2 sec delay after world

func print(message string, delay time.Duration, input chan string, output chan string) {
	for {
		<-input
		time.Sleep(delay * time.Second)
		fmt.Println(message)
		output <- "end"
	}
}

func Channel_hello_world() {
	ch1, ch2 := make(chan string), make(chan string)
	go print("hello", 1, ch1, ch2)
	go print("world", 2, ch2, ch1)
	ch1 <- "start"
	fmt.Println("Channel code executed")
	var exitVar string
	fmt.Println("Enter a key to exit")
	fmt.Scanf("%v", &exitVar) //if this line is omitted, the code main thread exits before the sleep
	os.Exit(0)
}
