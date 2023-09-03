package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

var (
	err error
	bot *tele.Bot
)

func init() {
	if bot, err = tele.NewBot(tele.Settings{
		Token:       os.Getenv("TOKEN"),
		Synchronous: true,
	}); err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Здаров братан, больше не пиши мне!")
	})
	bot.Handle("/help", func(c tele.Context) error {
		return c.Send(fmt.Sprintf("Братан, %s, я тебе ничем не помогу", c.Sender().FirstName))
	})
	bot.Handle(tele.OnQuery, func(c tele.Context) error {
		var (
			username    = c.Sender().Username
			description = "Узнай размер своего кутака"
		)

		if u := extractUsername(c.Query().Text); u != nil {
			username = *u
			description = "Узнай размер кутака " + username
		}

		var results tele.Results = []tele.Result{
			&tele.ArticleResult{
				ResultBase: tele.ResultBase{
					ID:        strconv.FormatInt(rand.Int63(), 10),
					Type:      "article",
					ParseMode: "Markdown",
				},
				Title:       "будь татарином, ярату кутак!",
				Text:        computeCock(username),
				Description: description,
				ThumbURL:    "https://attuale.ru/wp-content/uploads/2018/09/s-02b2e68f4f98047797058bcdc7740d566f223f4b.jpg",
			},
		}

		return c.Answer(&tele.QueryResponse{
			Results:   results,
			CacheTime: 60, // a minute
		})
	})
}

// Handler is the entry point for the AWS Lambda function or Yandex Cloud Functions
//
//goland:noinspection GoUnusedExportedFunction
func Handler(_ http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var u tele.Update
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Fatal(err)
	}
	bot.ProcessUpdate(u)
}
