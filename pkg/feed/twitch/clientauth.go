package twitch

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

var (
	creds    *viper.Viper
	clientID = creds.GetString("twitchclientid")
	// Consider storing the secret in an environment variable or a dedicated storage system.
	clientSecret = creds.GetString("twitchclientsecret")
	oauth2Config *clientcredentials.Config
)

func TwitchClientAuth() {
	oauth2Config = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Access token: %s\n", token.AccessToken)
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
