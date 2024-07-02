package basic

import (
	"log"
	"os"
	"proto-course/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtobufAsJSON(msg proto.Message, fname string) {
	// convertendo msg em json
	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatal("erro em marshal user para json:", err)
	}

	// escrevendo como json
	err = os.WriteFile(fname, jsonBytes, 0644)
	if err != nil {
		log.Fatal("erro em escrever arquivo bytes:", err)
	}
}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	// lendo arquivo
	b, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal("erro ao ler arquivo:", err)
	}

	if err := protojson.Unmarshal(b, dest); err != nil {
		log.Fatal("erro ao fazer unmarshal:", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()
	WriteProtobufAsJSON(u, "protobuf_json.json")
}

func ReadFromJSONSample() {
	dest := basic.User{}

	ReadProtoFromJSON("protobuf_json.json", &dest)

	log.Println(&dest)
}
