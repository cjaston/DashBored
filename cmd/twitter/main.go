package main

import (
	"bufio"
	"log"
	"os"

	"github.com/cjaston3/DashBored/pkg/feed/twitter"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	instr, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	twitter.GetTweets(string(instr))
}
