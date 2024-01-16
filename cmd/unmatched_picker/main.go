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
	err = tgBot.Listen(context.Background())
	if err != nil {
		log.Fatalf("error while listening tgbot: %v", err)
	}

	log.Println("stopped")

	//manager := pickManager.NewPickManager(3)
	//err = manager.RandDistribute()
	//if err != nil {
	//	log.Fatalf("failed to distribute: %v", err)
	//}
	//
	//fmt.Println(manager)
}
