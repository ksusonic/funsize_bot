package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func computeCock(username string) string {
	size := rand.Intn(51)
	var emoji string
	if size >= 15 {
		emoji = randomChoice([]string{"😏", "😱", "😂", "😁"})
	} else {
		emoji = randomChoice([]string{"😒", "☹️", "😣", "🥺"})
	}

	return fmt.Sprintf("Кутак @%s *%dсм* %s", username, size, emoji)
}

func randomChoice(choices []string) string {
	index := rand.Intn(len(choices))
	return choices[index]
}

func extractUsername(s string) *string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "@") && len(s) > 3 {
		s = strings.TrimPrefix(s, "@")
		return &s
	}

	return nil
}
