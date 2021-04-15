package main

import (
	"bufio"
	"log"
	"os"

	"github.com/cjaston3/DashBored/pkg/feed/twitter"
)

func main() {
<<<<<<< HEAD
	twitter.UserAuth()
	fmt.Println("Enter a Twitter handle: ")
=======
>>>>>>> parent of 4592dbb (finalized twitter API calls)
	reader := bufio.NewReader(os.Stdin)
	instr, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	twitter.GetTweets(string(instr))
}
