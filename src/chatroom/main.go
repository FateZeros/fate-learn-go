package main

import (
	"fmt"
	"chatroom/user"
)

func main() {
	s := user.Hello()
	fmt.Printf("s: %v\n", s)
}