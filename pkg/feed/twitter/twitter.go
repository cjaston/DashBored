package twitter

// OAuth2

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var creds *viper.Viper

func shit() {
	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     creds.GetString("twitterapikey"),
		ClientSecret: creds.GetString("twittersecret"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	user, _, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: "chrizdashiz",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.ID) //https://pkg.go.dev/github.com/dghubble/go-twitter/twitter#MediaEntity
	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{UserID: user.ID})
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tweets { // TODO: parse tweet data
		for _, m := range t.Entities.Media {
			fmt.Println(m.MediaURL)
		}
	}
}

func init() {
	creds = viper.New()
	creds.SetConfigName("creds")
	creds.AddConfigPath(".")
	creds.AddConfigPath("../../..")
	if err := creds.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
