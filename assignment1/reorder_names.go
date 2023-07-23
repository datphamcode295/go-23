package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("missing arguments")
		return
	}

	countryCode := args[len(args)-1]
	var isValid = validateCountryCode(countryCode)
	if isValid == false {
		fmt.Println("invalid country code")
		return
	}

	firstName := args[0]
	lastName := args[1]
	midName := strings.Join(args[2:len(args)-1], " ")

	orderedNames := orderedNames(firstName, lastName, midName, countryCode)

	fmt.Println(orderedNames)
}

func validateCountryCode(countryCode string) bool {
	var countryCodes = []string{"US", "VN"}
	countryCode = strings.ToUpper(countryCode)

	for _, code := range countryCodes {
		if code == countryCode {
			return true
		}
	}

	return false
}

func orderedNames(firstName string, lastName string, midName string, countryCode string) string {
	if midName != "" {
		midName = midName + " "
	}

	if countryCode == "US" {
		return firstName + " " + midName + lastName
	}

	return lastName + " " + midName + firstName
}
