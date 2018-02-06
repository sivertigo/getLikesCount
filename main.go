package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Result struct {
	CreatedAt  time.Time `json:"created_at"`
	LikesCount int       `json:"likes_count"`
	Title      string    `json:"title"`
	UpdatedAt  time.Time `json:"updated_at"`
	URL        string    `json:"url"`
}

func main() {
	//	lambda.Start(getLikesCount)
	getLikesCount()
}

func getLike(userId string) int {
	myurl := "http://qiita.com/api/v2/users/" + userId + "/items"
	client := &http.Client{}
	req, err := http.NewRequest("GET", myurl, nil)
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

	var results []Result

	err = json.Unmarshal([]byte(body), &results)
	if err != nil {
		log.Fatal(err)
	}
	var totalLikes int
	fmt.Print("ユーザー名" + userId + " ")
	for i := 0; i < len(results); i++ {
		if results[i].CreatedAt.Sub(time.Date(2018, time.January, 0, 0, 0, 0, 0, time.UTC)) > 0 {
			totalLikes += results[i].LikesCount
		}
	}
	fmt.Print(totalLikes)
	return totalLikes
}

func postLike(userId string, slackUrl string, total int) error {
	myUrl := "https://qiita.com/" + userId
	jsonStr := `{"text":"` + userId + `さんの2018年に書いた記事のいいね数は合計 ` + strconv.Itoa(total) + `です. URL: ` + myUrl + `"}`
	req, err := http.NewRequest(
		"POST",
		slackUrl,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func getLikesCount() {
	userId := []string{"sivertigo", "sadayuki-matsuno"}
	slackUrl := "https://hooks.slack.com/services/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"
	for i := 0; i < len(userId); i++ {
		postLike(userId[i], slackUrl, getLike(userId[i]))
	}
}
