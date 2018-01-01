package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Result []struct {
	RenderedBody   string      `json:"rendered_body"`
	Body           string      `json:"body"`
	Coediting      bool        `json:"coediting"`
	CommentsCount  int         `json:"comments_count"`
	CreatedAt      time.Time   `json:"created_at"`
	Group          interface{} `json:"group"`
	ID             string      `json:"id"`
	LikesCount     int         `json:"likes_count"`
	Private        bool        `json:"private"`
	ReactionsCount int         `json:"reactions_count"`
	Tags           []struct {
		Name     string        `json:"name"`
		Versions []interface{} `json:"versions"`
	} `json:"tags"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	User      struct {
		Description       string      `json:"description"`
		FacebookID        string      `json:"facebook_id"`
		FolloweesCount    int         `json:"followees_count"`
		FollowersCount    int         `json:"followers_count"`
		GithubLoginName   string      `json:"github_login_name"`
		ID                string      `json:"id"`
		ItemsCount        int         `json:"items_count"`
		LinkedinID        string      `json:"linkedin_id"`
		Location          string      `json:"location"`
		Name              string      `json:"name"`
		Organization      string      `json:"organization"`
		PermanentID       int         `json:"permanent_id"`
		ProfileImageURL   string      `json:"profile_image_url"`
		TwitterScreenName interface{} `json:"twitter_screen_name"`
		WebsiteURL        string      `json:"website_url"`
	} `json:"user"`
}

func main() {
	url := "http://qiita.com/api/v2/users/sivertigo/items"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("x-auth-token", "token1")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	var results []Result
	//	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
	//		panic(err)
	//	}
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
		fmt.Println(i)

	}
}
