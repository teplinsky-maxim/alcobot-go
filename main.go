package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	err = InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		results := make(tb.Results, 1)

		article := &tb.ArticleResult{
			Title: "How much should you drink today?",
			Text:  GenerateAnswer(uint(q.From.ID), q.From.Username),
		}

		results[0] = article
		results[0].SetResultID(strconv.Itoa(1))

		err := b.Answer(q, &tb.QueryResponse{
			Results:    results,
			IsPersonal: true,
			CacheTime:  60,
		})

		if err != nil {
			log.Println(err)
		}
	})

	b.Start()
}
