package basic

import (
	"proto-course/protogen/user"
)

func BaseUser() *user.User {
	newUser := user.User{
		Id:       1,
		Username: "victorreis",
		IsActive: true,
		Password: []byte("helloworld"),
		Emails:   []string{"teste@teste.com", "hello@world.com"},
		Gender:   user.Gender_GENDER_MALE,
	}

	// como a struct tem um Mutex, precisamos retornar um ponteiro para a struct, e não usar ela como ponteiro por sí própria, pois é mais seguro
	return &newUser
}
