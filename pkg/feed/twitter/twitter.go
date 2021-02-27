package twitter

// OAuth2

import (
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
}
func init() {
    creds = viper.New()
    creds.SetConfigName("creds")
    creds.AddConfigPath(".")
    if err := creds.ReadInConfig(); err != nil {
        log.Fatal(err)
    }
}
