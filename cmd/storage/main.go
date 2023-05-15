package main

import (
	"fmt"
	"log"

	// "github.com/google/uuid"
	"github.com/xor111xor/go-project-struct/v2/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("long string"))

	if err != nil {
		log.Fatal(err)
	}

	restoreFile, err := st.GetByID(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It is restored", restoreFile)

}
