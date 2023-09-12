package raindrops

import "fmt"

func Convert(number int) string {
	response := ""
	if number % 3 == 0 {
		response = fmt.Sprintf("%sPling", response)
	}
	if number % 5 == 0 {
		response = fmt.Sprintf("%sPlang", response)
	}
	if number % 7 == 0 {
		response = fmt.Sprintf("%sPlong", response)
	}
	if response == "" {
		response = fmt.Sprintf("%v", number)
	}
	return response
}
