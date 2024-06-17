package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

// CsvToSlice reads CSV data from a string and converts it to a map
func CsvToSlice(data string) (map[string][]string, error) {
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV data must contain at least a header row and one data row")
	}

	table := make(map[string][]string)
	headers := records[0]

	for _, header := range headers {
		table[header] = []string{}
	}

	for _, row := range records[1:] {
		for i, value := range row {
			table[headers[i]] = append(table[headers[i]], value)
		}
	}

	return table, nil
}

// ConnectAIModel sends a request to the AI model with the provided payload and token
func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	url := "https://huggingface.co/google/tapas-base-finetuned-wtq" // Placeholder URL

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		return Response{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func main() {
	// Open the CSV file
	file, err := os.Open("data-series.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read CSV data from the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert CSV data to a map
	table, err := CsvToSlice(string(data))
	if err != nil {
		fmt.Println("Error converting CSV to map:", err)
		return
	}

	// Define the query
	query := "Find the sum of column A"

	// Prepare the payload
	payload := Inputs{
		Table: table,
		Query: query,
	}

	// Create a new AIModelConnector
	connector := AIModelConnector{
		Client: &http.Client{},
	}

	// Define the token
	token := "your_api_token_here"

	// Send the request to the AI model
	response, err := connector.ConnectAIModel(payload, token)
	if err != nil {
		fmt.Println("Error connecting to AI model:", err)
		return
	}

	// Print the response
	fmt.Println("Answer:", response.Answer)
	fmt.Println("Coordinates:", response.Coordinates)
	fmt.Println("Cells:", response.Cells)
	fmt.Println("Aggregator:", response.Aggregator)
}
