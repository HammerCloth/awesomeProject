package controller

import (
	"awesomeProject/service"
)

func Feed(time int) bool {
	videoService := service.VideoServiceImpl{
		UserService: &service.VideoSub{},
	}
	return videoService.Feed(time)
}
