package main

import (
	"fmt"
	"strings"

	"github.com/icelain/jokeapi"
)

type Joke struct {
	Joke      string	`json:"joke"`
	Category  string	`json:"category"`
}

func main() {
	api := jokeapi.New()
	api.Set(jokeapi.Params{Blacklist: []string{"nsfw"}, JokeType: "single", Categories: []string{"Programming"}, Lang: "en"})
  res, err := api.Fetch()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%s\n\n", strings.Join(res.Joke, ""))
}