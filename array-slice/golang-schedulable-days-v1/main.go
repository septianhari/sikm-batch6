package main

func SchedulableDays(date1 []int, date2 []int) []int {
	var duaduanyabisa []int
	duaduanyabisa = make([]int, 0)

	// for i := 0; i <= 31; i++ {
	// 	isImamAvailable := false
	// 	isPermanaAvailable := false

	// 	for _, date := range date1 {
	// 		if date == i {
	// 			isImamAvailable = true
	// 		}
	// 	}

	// 	for _, date := range date2 {
	// 		if date == i {
	// 			isPermanaAvailable = true
	// 		}
	// 	}

	// 	if isImamAvailable && isPermanaAvailable {
	// 		result = append(result, i)
	// 	}
	// }

	for _, dateImam := range date1 {
		for _, datePermana := range date2 {

			if dateImam == datePermana {
				duaduanyabisa = append(duaduanyabisa, dateImam)
				break
			}
		}
	}

	return duaduanyabisa
}
