package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Response structure
type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

// Item post value
type Item struct {
	Title    string
	URL      string
	Comments int  `json:"num_comments"`
	IsSelf   bool `json:"is_self"`
	Score    int
	Domain   string
}

func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
		// nothing
	case 1:
		com = "1 comment"
	default:
		com = fmt.Sprintf("%d comments", i.Comments)

	}
	return fmt.Sprintf("%s(%s | %d score)\n%s\n", i.Title, com, i.Score, i.URL)
}

// Get a subreddits items as json
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}

	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data

	}

	return items, nil

}

// FindBestThread computers a score based on the
// nb of comments and the score of each threads.
func FindBestThread(subReddit string) (Item, error) {
	items, err := Get(subReddit)

	CommentsAvg := 0
	ScoreAvg := 0

	for _, item := range items {
		CommentsAvg += item.Comments
		ScoreAvg += item.Score
	}
	CommentsAvg = CommentsAvg / len(items)
	ScoreAvg = ScoreAvg / len(items)

	highestTotal := 0
	selectedItem := Item{}

	for _, item := range items {
		if len(item.Title) > 90 || item.IsSelf {
			continue
		}

		total := item.Comments + item.Score
		if total > highestTotal {
			highestTotal = total
			selectedItem = item
		}
	}
	return selectedItem, err
}

// GetAPI returns a new instance of a TwitterApi
