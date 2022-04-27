package main

import (
	"fmt"
	"time"
	"twitter-lofi/usecase"
)

const TIME_BETWEEN_POSTS = 5

func postLofi(countPost int) {
	search := "but it's lofi remix"

	lofiUsecase := usecase.LofiUseCase{}
	lofiPosted := lofiUsecase.Execute(search)

	fmt.Println("O lofi postado foi o ", lofiPosted.Title)
	fmt.Println("Irei postar mais um lofi em 10 minutos")

	time.Sleep(TIME_BETWEEN_POSTS * time.Minute)

	if countPost < 5 {
		postLofi(countPost)
	}
}

func main() {
	initCount := 0

	postLofi(initCount)
}
