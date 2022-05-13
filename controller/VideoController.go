package controller

import (
	"awesomeProject/service/video"
)

func Feed(time int) bool {
	videoService := video.VideoServiceImpl{
		UserService: &video.VideoSub{},
	}
	return videoService.Feed(time)
}
