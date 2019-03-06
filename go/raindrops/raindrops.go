package raindrops

import "strconv"

//Convert takes in a number and returns the converted string
func Convert(number int) string {
	returnString := ""
	if number%3 == 0 {
		returnString += "Pling"
	}
	if number%5 == 0 {
		returnString += "Plang"
	}
	if number%7 == 0 {
		returnString += "Plong"
	}
	if returnString == "" {
		returnString += strconv.Itoa(number)
	}
	
	return returnString
}
