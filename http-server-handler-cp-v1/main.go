package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()

		day := now.Day()
		month := now.Month()
		year := now.Year()

		fmt.Println(time.Monday)
		result := fmt.Sprintf("%s, %d %s %d", now.Weekday(), day, month, year)
		writer.Write([]byte(result))
	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
