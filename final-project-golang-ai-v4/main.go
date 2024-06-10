package main

import (
	"net/http"
)

type AIModelConnector struct {
	Client *http.Client
}

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type Response struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

func CsvToSlice(data string) (map[string][]string, error) {
	return nil, nil // TODO: replace this
}

func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	return Response{}, nil // TODO: replace this
}

func main() {
	// TODO: answer here
}
