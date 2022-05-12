package service

type VideoServiceImpl struct {
	UserService //依赖注入===》依赖抽象不依赖具体，依赖接口不依赖实现
}

func (Video *VideoServiceImpl) Feed(time int) bool {
	//逻辑代码
	Video.UserService.IsAdmin(time)
	//Video.IsAdmin(time)
	return true
}
