package twitter

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var creds *viper.Viper

func GetTweets(handle string) {
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
		ScreenName: handle,
	})
	if err != nil {
		log.Fatal(err)
	}

	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{UserID: user.ID})
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tweets {
		//parse tweet data
		if t.RetweetedStatus == nil {
			fmt.Println(t.CreatedAt)
			fmt.Println(t.Text)
			//print links
			for _, m := range t.Entities.Media {
				fmt.Println(m.MediaURL)
			}
			fmt.Printf("\n\n")
		}
	}

	trendsList, _, err := client.Trends.Place(23424977, &twitter.TrendsPlaceParams{WOEID: 23424977})
	if err != nil {
		log.Fatal(err)
	}
	//markov chain mapping each word of a tweet into the next word in the senctence
	fmt.Printf("TRENDING TOPICS USA \n\n")
	for _, x := range trendsList {
		for _, y := range x.Trends {

			fmt.Printf(`"%s" `, y.Name)
			fmt.Println(y.TweetVolume, "NEW TWEETS")
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
