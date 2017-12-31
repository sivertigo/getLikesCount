package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	RenderedBody String
	Id           int `json:"id"`
	LikesCount   int `json:"likes_count"`
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
