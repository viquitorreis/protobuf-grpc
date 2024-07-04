package basic

import (
	"log"
	"proto-course/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 6969,
		Slug:          "/v1-test",
		Title:         "Titulo teste para a v1",
		HtmlContent:   "<p>Dummy parágrafo para o UserContent</p>",
		AuthorId:      99,
	}

	WriteProtoToFile(&uc, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Ler v1")
	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v1.bin", &uc)

	log.Println(&uc)

	ucJson, err := protojson.Marshal(&uc)
	if err != nil {
		log.Fatal("erro ao fazer marshal do user content:", err)
	}

	log.Println(string(ucJson))
}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 6970,
		Slug:          "/v2-test",
		Title:         "Titulo teste para a v2",
		HtmlContent:   "<p>Dummy parágrafo para o UserContent VERSAO 2</p>",
		AuthorId:      100,
		// Category: "categoria teste",
	}

	WriteProtoToFile(&uc, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Ler v2")
	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v2.bin", &uc)

	log.Println(&uc)

	ucJson, err := protojson.Marshal(&uc)
	if err != nil {
		log.Fatal("erro ao fazer marshal do user content:", err)
	}

	log.Println(string(ucJson))
}

func BasicWriteUserContentV3() {
	uc := basic.UserContent{
		UserContentId: 7000,
		Slug:          "/v3-test",
		Title:         "Titulo teste para a v2",
		HtmlContent:   "<p>VERSAO 3 Dummy parágrafo para o UserContent VERSAO 3</p>",
		AuthorId:      100,
		// Category: "VERSAO 3 categoria teste VERSAO 3",
	}

	WriteProtoToFile(&uc, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Ler v3")
	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v3.bin", &uc)

	log.Println(&uc)

	ucJson, err := protojson.Marshal(&uc)
	if err != nil {
		log.Fatal("erro ao fazer marshal do user content:", err)
	}

	log.Println(string(ucJson))
}
