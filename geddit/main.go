package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	textapi "github.com/AYLIEN/aylien_textapi_go"
	"github.com/ChimeraCoder/anaconda"
	"log"
	// "net/url" //for URL
	"os"

	"github.com/nicolas-martin/goredditgo"
)

const row = "========"

func main() {
	//subReddits := []string{"startups", "futurism", "technology"}
	if len(os.Args) == 0 {
		log.Fatalln("Must provide the subreddit")
	}

	subReddits := os.Args[1:]

	keys, err := readKeys("/Users/nma/go/src/github.com/nicolas-martin/reddit/keys")
	if err != nil {
		log.Fatal(err)
	}

	api, err := getAPI(keys[0], keys[1], keys[2], keys[3])
	if err != nil {
		log.Fatal(err)
	}

	for _, subReddit := range subReddits {

		selectedItem, err := reddit.FindBestThread(subReddit)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n%s\n%s\n", subReddit, row)

		encoded, _ := json.MarshalIndent(selectedItem, "", "    ")
		fmt.Println(string(encoded))

		hashtags := getHashTags(selectedItem.URL, keys[4], keys[5])
		fmt.Printf("\n%s\n%v\n", hashtags.Hashtags, row)

		var tweetText string

		if len(hashtags.Hashtags) > 1 {
			tweetText = fmt.Sprintf("%s %s %s %s", selectedItem.Title, selectedItem.URL, hashtags.Hashtags[0], hashtags.Hashtags[1])
		} else if len(hashtags.Hashtags) == 0 {
			tweetText = fmt.Sprintf("%s %s", selectedItem.Title, selectedItem.URL)
		} else {
			tweetText = fmt.Sprintf("%s %s %s", selectedItem.Title, hashtags.Hashtags[0], selectedItem.URL)
		}

		fmt.Println(api.GetBlocksIds(nil))
		//_, err = api.PostTweet(tweetText, nil)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Output: \n%s\n", tweetText)

	}
}

func readKeys(fileToRead string) ([]string, error) {
	file, err := os.Open(fileToRead)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getAPI(consumerKey string, consumerSecret, accessToken string, accessTokenScret string) (*anaconda.TwitterApi, error) {

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenScret)

	return api, nil
}

func getHashTags(url string, ApplicationID string, ApplicationKey string) *textapi.HashtagsResponse {
	auth := textapi.Auth{ApplicationID: ApplicationID, ApplicationKey: ApplicationKey}
	client, err := textapi.NewClient(auth, true)
	if err != nil {
		panic(err)
	}
	params := &textapi.HashtagsParams{URL: url}
	hashtags, err := client.Hashtags(params)
	if err != nil {
		panic(err)
	}

	return hashtags
}
