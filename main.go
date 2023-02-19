package main

/*
stopAllCh := make(chan struct{}) // If close this chanel all goroutine will cancel
If close this chanel all goroutine will cancel
If close this chanel all goroutine will cancel
If close this chanel all goroutine will cancel
*/

import (
	"fmt"
	"makelintdemo/sub"
	"time"
)

// this is main func.
func main() {
	stopAllCh := make(chan struct{}) // If close this chanel all goroutine will cancel.
	sub.Sub(stopAllCh)
	fmt.Println("sub() 回来了") // do it
	time.Sleep(time.Millisecond * 200000)
}
