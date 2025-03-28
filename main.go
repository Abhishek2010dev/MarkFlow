package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Hello struct{}

type HelloMessage interface {
	Hello()
}
