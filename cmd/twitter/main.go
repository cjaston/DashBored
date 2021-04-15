package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cjaston3/DashBored/pkg/feed/twitter"
)

func main() {
	twitter.UserAuth()
	fmt.Println("Enter a Twitter handle: ")
	reader := bufio.NewReader(os.Stdin)
	instr, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	twitter.GetTweets(string(instr))
}
