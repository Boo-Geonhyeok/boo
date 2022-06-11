package handler

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rest/models"
	"strings"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var SearchedSpotifyStruct models.SpotifyStruct
var SearchedYoutubeStruct models.YoutubeStruct

func searchSpotify(songTitle string) {
	accessToken := authSpotify()

	songTitle = hex.EncodeToString([]byte(songTitle))
	songTitle = insertInto(songTitle, 2, '%')

	url := "https://api.spotify.com/v1/search?q=" + songTitle + "&type=track&market=KR&limit=10&offset=0"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &SearchedSpotifyStruct)

	if err != nil {
		panic(err)
	}
}

func searchYoutube(query string) {
	query = strings.ReplaceAll(query, " ", "%20")
	url := "https://www.googleapis.com/youtube/v3/search?part=snippet&key=AIzaSyAP0JhjXnvfl9-jz2iHr8c0n7y3vxMvXzs&q=" + query + "&type=video&topicId=music&videoEmbeddable=true"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &SearchedYoutubeStruct)

	if err != nil {
		panic(err)
	}
}

func authSpotify() string {
	authConfig := &clientcredentials.Config{
		ClientID:     "a19a65fbd87b475c8274134323040d1d",
		ClientSecret: "b85db7814f8746fcbe11fcb3ed39f8d6",
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieve access token: %v", err)
	}

	return accessToken.AccessToken
}

func insertInto(s string, interval int, sep rune) string {
	var buffer bytes.Buffer
	before := interval - 1
	last := len(s) - 1
	for i, char := range s {
		buffer.WriteRune(char)
		if i%interval == before && i != last {
			buffer.WriteRune(sep)
		}
	}
	data := buffer.String()
	upperData := "%" + strings.ToUpper(data)
	return upperData
}
