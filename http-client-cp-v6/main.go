package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := &http.Client{}
	animechan := []Animechan{}

	req, err := http.NewRequest("GET", "https://animechan.vercel.app/api/quotes/anime?title=naruto", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBody, &animechan); err != nil {
		return nil, err
	}

	fmt.Println(len(animechan))
	// Hit API https://animechan.vercel.app/api/quotes/anime?title=naruto with method GET:
	return animechan, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	// http.Post implementation
	resp, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	if err != nil {
		return Postman{}, err
	}
	defer resp.Body.Close()

	// read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Postman{}, err
	}

	postMan := Postman{}
	if err := json.Unmarshal(responseBody, &postMan); err != nil {
		return Postman{}, err
	}

	// // Hit API https://postman-echo.com/post with method POST
	return postMan, nil
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
