package service

type answer struct {
	key  int
	kshf int
}

type UserService interface {
	IsAdmin(id int) answer
}
