package basic

import (
	"log"
	"proto-course/protogen/basic"
)

func BasicHello() {
	hello := &basic.Hello{
		Name: "Víctor",
	}

	log.Println(hello)
}
