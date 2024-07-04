package basic

import (
	"log"
	"proto-course/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicReadUserPayment() {
	log.Println("BasicReadUserPayment")

	up := basic.UserPayment{}

	ReadProtoFromFile("user_content_v1.bin", &up)

	log.Println(&up)

	upJson, err := protojson.Marshal(&up)
	if err != nil {
		log.Fatal("erro ao fazer marshal do user payment:", err)
	}

	log.Println(string(upJson))
}
