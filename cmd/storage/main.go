package main

import "fmt"
import "ozon/internal/storage"

func main() {
	st := storage.NewStorage()
	fmt.Println("hell", st)
}
