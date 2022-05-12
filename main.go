package main

import (
	"awesomeProject/controller"
	"fmt"
)

func main() {
	key := controller.Feed(1)
	fmt.Print(key)
}
