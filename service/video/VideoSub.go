package video

import "awesomeProject/service/user"

type VideoSub struct {
}

func (video *VideoSub) IsAdmin(id int) user.Answer {
	a := user.Answer{
		Key:  0,
		Kshf: 0,
	}
	return a
}
