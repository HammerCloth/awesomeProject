package user

type Answer struct {
	Key  int
	Kshf int
}

type UserService interface {
	IsAdmin(snd int) Answer
}
