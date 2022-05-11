package main

import (
	"awesomeProject/service"
	"fmt"
)

func main() {
	userService := service.UserServiceImpl{}
	video := service.VideoServiceImpl{
		UserService: &userService,
	}
	key := video.Feed(1)
	fmt.Print(key)
}
