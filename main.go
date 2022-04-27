package main

import (
	"fmt"
	"twitter-lofi/usecase"
)

const TIME_BETWEEN_POSTS = 5

func postLofi(countPost int) {
	search := "but it's lofi remix"

	lofiUsecase := usecase.LofiUseCase{}
	lofiPosted := lofiUsecase.Execute(search)

	fmt.Println("O lofi postado foi o ", lofiPosted.Title)
}

func main() {
	initCount := 0

	postLofi(initCount)
}
