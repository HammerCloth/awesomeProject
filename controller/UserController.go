package controller

import (
	"awesomeProject/service/user"
	"fmt"
)

func gigigi() {
	user := user.UserServiceImpl{}
	answer := user.IsAdmin(2)
	fmt.Print(answer.Kshf)
}
