package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	ticketPrice := map[string]int{
		"VIP":     30,
		"Regular": 20,
		"Student": 10,
	}

	total := VIP*ticketPrice["VIP"] + regular*ticketPrice["Regular"] + student*ticketPrice["Student"]
	var discount float32

	if total >= 100 {
		if day%2 == 0 { // even day
			if VIP+regular+student >= 5 {
				discount = 0.2
			} else {
				discount = 0.1
			}
		} else { // odd day
			if VIP+regular+student >= 5 {
				discount = 0.25
			} else {
				discount = 0.15
			}
		}
	}

	return float32(total) - float32(total)*discount
}

func main() {
	// Test cases
	fmt.Println(GetTicketPrice(1, 1, 1, 20)) // Output: 60
	fmt.Println(GetTicketPrice(3, 3, 3, 20)) // Output: 144
	fmt.Println(GetTicketPrice(4, 0, 0, 21)) // Output: 102
}
