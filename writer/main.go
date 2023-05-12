package main

import (
	"log"
	"os"

	userpb "github.com/stickza-file-management-API/gen/go/user/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := userpb.User{
		Uuid:      "1-2-3-4",
		FullName:  "Shrio",
		BirthYear: 20,
	}

	b, err := proto.Marshal(&u)
	if err != nil {
		log.Fatalln("Marshalling error", err)

	}

	if err := os.WriteFile("user.bin", b, 0664); err != nil {
		log.Fatalln("Writing error", err)
	}
}
