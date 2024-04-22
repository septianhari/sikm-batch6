package main

import (
	"errors"
	"fmt"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	// Check for empty domain
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
		return
	}

	// Check for invalid domain
	if !website.Valid {
		chErr <- errors.New("domain not valid")
		return
	}

	// Check for invalid RefIPs
	if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
		return
	}

	// Fill TLD and IDN_TLD fields
	TLD, IDN_TLD := GetTLD(website.Domain)
	website.TLD = TLD
	website.IDN_TLD = IDN_TLD

	// Send the updated website data to the channel
	ch <- website
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)

	// Start goroutines to process each website data
	for _, website := range data {
		go ProcessGetTLD(website, ch, errCh)
	}

	// Collect processed website data
	var processedData []RowData
	for i := 0; i < len(data); i++ {
		select {
		case website := <-ch:
			// Check if the website matches the provided TLD
			if website.TLD == TLD {
				processedData = append(processedData, website)
			}
		case err := <-errCh:
			// Return the first error encountered
			return nil, err
		}
	}

	return processedData, nil
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
