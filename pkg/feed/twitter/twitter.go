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
<<<<<<< HEAD
	//markov chain mapping each word of a tweet into the next word in the senctence
	fmt.Printf("TRENDING TOPICS USA \n\n")
	for _, x := range trendsList {
		for _, y := range x.Trends {

			fmt.Printf(`"%s" `, y.Name)
<<<<<<< HEAD
<<<<<<< HEAD
			fmt.Println(y.TweetVolume, "NEW TWEETS")
=======
	for _, t := range tweets { // TODO: parse tweet data
		for _, m := range t.Entities.Media {
			fmt.Println(m.MediaURL)
>>>>>>> parent of a04e09d (twitter trends + exported GetTweets)
=======
			fmt.Println(y.TweetVolume, "NEW TWEETS \n")
>>>>>>> parent of 4592dbb (finalized twitter API calls)
=======
			fmt.Println(y.TweetVolume, "NEW TWEETS \n")
>>>>>>> parent of 4592dbb (finalized twitter API calls)
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
