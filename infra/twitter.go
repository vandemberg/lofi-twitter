package infra

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	config "twitter-lofi/config"

	"github.com/kurrik/oauth1a"
)

func Twitter(twitterText string) {
	consumerKey := config.GoDotEnvVariable("TWITTER_CLIENT_KEY")
	consumerSecret := config.GoDotEnvVariable("TWITTER_CLIENT_SECRET")
	accessToken := config.GoDotEnvVariable("TWITTER_ACCESS_TOKEN")
	accessSecret := config.GoDotEnvVariable("TWITTER_TOKEN_SECRET")

	payload := strings.NewReader(`{"text":"` + twitterText + `"}`)

	service := &oauth1a.Service{
		RequestURL:   "https://api.twitter.com/oauth/request_token",
		AuthorizeURL: "https://api.twitter.com/oauth/request_token",
		AccessURL:    "https://api.twitter.com/oauth/request_token",
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			CallbackURL:    "https://twittter.com/vandemberglima",
		},
		Signer: new(oauth1a.HmacSha1Signer),
	}

	userConfig := oauth1a.NewAuthorizedConfig(accessToken, accessSecret)

	httpRequest, _ := http.NewRequest("POST", "https://api.twitter.com/2/tweets", payload)
	httpRequest.Header.Add("Content-Type", "application/json")
	service.Sign(httpRequest, userConfig)

	httpClient := new(http.Client)
	httpResponse, err := httpClient.Do(httpRequest)

	if err != nil {
		fmt.Println(err)
		return
	}

	buff := bytes.NewBuffer(nil)
	httpResponse.Write(buff)
}
