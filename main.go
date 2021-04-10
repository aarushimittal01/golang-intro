package main

import (
	"fmt"
	"golang-intro/counter"
)

func main() {
	var msg string
	msg = "Hello World!"
	fmt.Println(msg)

	//Counters
	count := counter.Get_counter()
	up, down := counter.Get_value()

	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

	fmt.Println("Calling UP function")
	fmt.Println(up())
	fmt.Println(up())
	fmt.Println(up())
	fmt.Println(up())

	fmt.Println("Calling DOWN function")
	fmt.Println(down())
	fmt.Println(down())
	fmt.Println(down())
	fmt.Println(down())

	//Get Series
	even_nos := counter.Get_series(0, 2)
	fmt.Println(even_nos())
	fmt.Println(even_nos())
	fmt.Println(even_nos())
	fmt.Println(even_nos())

	odd_nos := counter.Get_series(-1, 2)
	fmt.Println(odd_nos())
	fmt.Println(odd_nos())
	fmt.Println(odd_nos())
	fmt.Println(odd_nos())

	nos_divisible_by_3 := counter.Get_series(0, 3)
	fmt.Println(nos_divisible_by_3())
	fmt.Println(nos_divisible_by_3())
	fmt.Println(nos_divisible_by_3())

	//condition for numbers
	get_even_nos := counter.Get_numbers(func(n int) bool {
		if n%2 == 0 {
			return true
		} else {
			return false
		}
	})

	get_odd_nos := counter.Get_numbers(func(n int) bool {
		if n%2 != 0 {
			return true
		} else {
			return false
		}
	})

	fmt.Println(get_even_nos())
	fmt.Println(get_even_nos())
	fmt.Println(get_even_nos())

	fmt.Println(get_odd_nos())
	fmt.Println(get_odd_nos())
	fmt.Println(get_odd_nos())
}
