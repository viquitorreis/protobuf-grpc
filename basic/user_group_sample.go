package basic

import (
	"log"
	"proto-course/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	group := basic.UserGroup{
		GroupId:   1,
		GroupName: "Admin",
		Roles:     []string{"admin", "root"},
		Users: []*basic.User{
			{
				Id:       1,
				Username: "victorreis",
				IsActive: true,
				Password: []byte("helloworld"),
				Emails:   []string{"victor@reis.com"},
			},
			{
				Id:       2,
				Username: "randomwomen",
				IsActive: true,
				Password: []byte("pass123"),
				Emails:   []string{"random@email.com"},
			},
		},
	}

	log.Println("grupo criado:", &group)

	groupCopy := &group
	jsonBytes, err := protojson.Marshal(groupCopy)
	if err != nil {
		log.Fatal("erro no marshal base UserGroup:", err)
	}
	log.Println(string(jsonBytes))
}
