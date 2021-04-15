package twitter

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)


func UserAuth() {

	config := oauth1.NewConfig(creds.GetString("twitterapikey"), creds.GetString("twittersecret"))
	token := oauth1.NewToken(creds.GetString("twitteraccesstoken"), creds.GetString("twitteraccesssecret"))
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's ACCOUNT:\n%+v\n", user)

	//Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	fmt.Printf("User's HOME TIMELINE:\n%+v\n", user)

	//Mention Timeline
	mentionTimelineParams := &twitter.MentionTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ = client.Timelines.MentionTimeline(mentionTimelineParams)
	fmt.Printf("User's MENTION TIMELINE:\n%+v\n", tweets)

	// Retweets of Me Timeline
	retweetTimelineParams := &twitter.RetweetsOfMeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ = client.Timelines.RetweetsOfMeTimeline(retweetTimelineParams)
	fmt.Printf("User's 'RETWEETS OF ME' TIMELINE:\n%+v\n", tweets)

	/*requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		log.Fatal(err)
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, req, authorizationURL.String(), http.StatusFound)

	requestToken, verifier, err := oauth1.ParseAuthorizationCallback(req)
	if err != nil {
		log.Fatal(err)
	}

	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, verifier)
	if err != nil {
		log.Fatal(err)
	}*/

}
