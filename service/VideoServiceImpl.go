package service

type VideoServiceImpl struct {
	UserService
}

func (Video *VideoServiceImpl) Feed(time int) bool {
	Video.IsAdmin(time)
	return true
}
