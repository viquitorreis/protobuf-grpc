package basic

import (
	"log"
	"proto-course/protogen/user"

	"google.golang.org/protobuf/encoding/protojson"
)

func BaseUser() {
	newUser := user.User{
		Id:       1,
		Username: "victorreis",
		IsActive: true,
		Password: []byte("helloworld"),
		Emails:   []string{"teste@teste.com", "hello@world.com"},
		Gender:   user.Gender_GENDER_MALE,
	}

	jsonBytes, err := protojson.Marshal(&newUser)
	if err != nil {
		log.Fatal("erro no marshal base User:", err)
	}

	log.Println(string(jsonBytes))
}

func ProtoJsonUser() {
	p := user.User{
		Id:       2,
		Username: "randomwomen",
		IsActive: true,
		Password: []byte("pass123"),
		Emails:   []string{"marie@teste.com", "test@teste.com"},
		Gender:   user.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonProtoUser() {
	json := `{
				"id":2,
				"username":"randomwomen",
				"is_active":true,
				"password":"cGFzczEyMw==",
				"emails":["marie@teste.com","test@teste.com"],
				"gender":"GENDER_FEMALE"}
			`

	p := user.User{}
	err := protojson.Unmarshal([]byte(json), &p)
	if err != nil {
		log.Fatal("erro em json to proto user:", err)
	}

	log.Println(&p)
}
