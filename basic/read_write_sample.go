package basic

import (
	"log"
	"os"
	"proto-course/protogen/basic"

	"google.golang.org/protobuf/proto"
)

func dummyUser() *basic.User {
	addr := basic.Address{
		Street:  "random street",
		City:    "random city",
		Country: "Brazil",
	}

	userSkillRating := map[string]uint32{
		"randomSkill":  5,
		"randomSkill2": 3,
		"randomSkill3": 1,
	}

	instantMsg := basic.InstantMessaging{
		InstantMessagingProduct:  "hello, how are you?",
		InstantMessagingUsername: "randUsername",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMsg,
	}

	newUser := basic.User{
		Id:                   1,
		Username:             "victorreis",
		IsActive:             true,
		Password:             []byte("helloworld"),
		Emails:               []string{"teste@teste.com", "hello@world.com"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		EletronicCommChannel: &ecomm,
		SkillRating:          userSkillRating,
	}

	return &newUser
}

func WriteProtoToFile(msg proto.Message, fname string) {
	// convertendo msg em bytes
	userBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("erro em marshal user para bytes:", err)
	}

	// escrevendo como bytes
	err = os.WriteFile(fname, []byte(userBytes), 0644)
	if err != nil {
		log.Fatal("erro em escrever arquivo bytes:", err)
	}
}

func WriteToFileHelper() {
	WriteProtoToFile(dummyUser(), "user.bin")
}

func ReadProtoFromFile(fname string, dest proto.Message) {
	// lendo arquivo
	b, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal("erro ao ler arquivo:", err)
	}

	if err := proto.Unmarshal(b, dest); err != nil {
		log.Fatal("erro ao fazer unmarshal:", err)
	}
}

func BasicReadingUserBytesFile(fname string) {
	// lendo arquivo bytes
	userBytes, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal("erro em ler arquivo bytes:", err)
	}

	// traduzindo bytes em mensagens proto
	u := basic.User{}
	err = proto.Unmarshal(userBytes, &u)
	if err != nil {
		log.Fatal("erro em unmarshal user bytes:", err)
	}

	log.Println(u.String())
}
