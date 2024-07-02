package basic

import (
	"log"
	"math/rand"
	"proto-course/protogen/basic"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func BaseUser() {
	addr := basic.Address{
		Street:  "random street",
		City:    "random city",
		Country: "Brazil",
	}

	newUser := basic.User{
		Id:       1,
		Username: "victorreis",
		IsActive: true,
		Password: []byte("helloworld"),
		Emails:   []string{"teste@teste.com", "hello@world.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
	}

	jsonBytes, err := protojson.Marshal(&newUser)
	if err != nil {
		log.Fatal("erro no marshal base User:", err)
	}

	log.Println(string(jsonBytes))
}

func ProtoJsonUser() {
	p := basic.User{
		Id:       2,
		Username: "randomwomen",
		IsActive: true,
		Password: []byte("pass123"),
		Emails:   []string{"marie@teste.com", "test@teste.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
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

	p := basic.User{}
	err := protojson.Unmarshal([]byte(json), &p)
	if err != nil {
		log.Fatal("erro em json to proto user:", err)
	}

	log.Println(&p)
}

func UserWithCoordinates() {
	loc := basic.Address_Coordinate{
		Latitude:  123.123,
		Longitude: 321.321,
	}

	addr := basic.Address{
		Street:     "random street",
		City:       "random city",
		Country:    "Brazil",
		Coordinate: &loc,
	}

	newUser := basic.User{
		Id:                   1,
		Username:             "victorreis",
		IsActive:             true,
		Password:             []byte("helloworld"),
		Emails:               []string{"victor@reis.com", "victor@ftw.com"},
		Address:              &addr,
		CommunicationChannel: randomCommunicationChannel(),
	}

	log.Println("user created:", &newUser)
}

func randomCommunicationChannel() *anypb.Any {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	paperMail := basic.PaperMail{
		PaperMailAddress: "random street, random city, Brazil",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "randSocialMedia",
		SocialMediaUsername: "randUsername",
	}

	instantMsg := basic.InstantMessaging{
		InstantMessagingProduct:  "randProduct",
		InstantMessagingUsername: "randUsername",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	case 2:
		anypb.MarshalFrom(&a, &instantMsg, proto.MarshalOptions{})
	}

	return &a
}

func BasicUnmarshalAnyKnow() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "randSocialMedia",
		SocialMediaUsername: "randUsername",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// como sabemos o tipo de mensagem que está dentro do any, podemos fazer o unmarshal também
	// sem a necessidade de fazer um switch case
	sm := &basic.SocialMedia{}
	if err := a.UnmarshalTo(sm); err != nil {
		log.Fatal("erro em unmarshal any know:", err)
	}

	jsonBytes, err := protojson.Marshal(sm)
	if err != nil {
		log.Fatal("erro em marshal any know:", err)
	}

	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyUnknow() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage
	unmarshaled, err := a.UnmarshalNew()
	if err != nil {
		log.Fatal("erro em unmarshal any unknow:", err)
	}

	log.Println("Unmarshal feito como:", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, err := protojson.Marshal(unmarshaled)
	if err != nil {
		log.Fatal("erro em marshal any unknow:", err)
	}
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := &basic.PaperMail{}

	if a.MessageIs(pm) {
		log.Println("é um paper mail")
		if err := a.UnmarshalTo(pm); err != nil {
			log.Fatal("erro em unmarshal any is:", err)
		}

		jsonBytes, err := protojson.Marshal(pm)
		if err != nil {
			log.Fatal("erro em marshal any is:", err)
		}
		log.Println(string(jsonBytes))
	} else {
		log.Println("não é um paper mail e sim um:", a.TypeUrl)
	}
}

func BasicUserOneOf() {
	// socialMedia := basic.SocialMedia{
	// 	SocialMediaPlatform: "randSocialMedia",
	// 	SocialMediaUsername: "randUsername",
	// }

	// ecomm := basic.User_SocialMedia{
	// 	SocialMedia: &socialMedia,
	// }

	instantMsg := basic.InstantMessaging{
		InstantMessagingProduct:  "hello, how are you?",
		InstantMessagingUsername: "randUsername",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMsg,
	}

	u := basic.User{
		Id:                   96,
		IsActive:             true,
		Password:             []byte("helloworld"),
		Emails:               []string{"victor@victor.victor"},
		EletronicCommChannel: &ecomm,
	}

	jsonBytes, err := protojson.Marshal(&u)
	if err != nil {
		log.Fatal("erro em marshal user one of:", err)
	}

	log.Println(string(jsonBytes))
}

func BasicUserWithSkillRating() {
	userSkillRating := map[string]uint32{
		"randomSkill":  5,
		"randomSkill2": 3,
		"randomSkill3": 1,
	}

	u := basic.User{
		Id:          96,
		IsActive:    true,
		Password:    []byte("helloworld"),
		Emails:      []string{"victor@victor.victor"},
		SkillRating: userSkillRating,
	}

	jsonBytes, err := protojson.Marshal(&u)
	if err != nil {
		log.Fatal("erro em marshal user one of:", err)
	}

	log.Println(string(jsonBytes))
}
