package main

import (
	"context"
	"log"
	"os"

	"unmatched_picker/internal/bot"
)

func main() {
	tgBot, err := bot.NewBot(os.Getenv("TG_BOT_API_TOKEN"))
	if err != nil {
		log.Fatalf("failed to init tg bot: %v", err)
	}

	// TODO: graceful shutdown
	tgBot.Listen(context.Background())

	log.Println("stopped")

	//manager := pickManager.NewPickManager(3)
	//err = manager.RandDistribute()
	//if err != nil {
	//	log.Fatalf("failed to distribute: %v", err)
	//}
	//
	//fmt.Println(manager)
}
