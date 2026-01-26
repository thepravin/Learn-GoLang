package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)

	//---------------------------------------------------

	currentTime2 := time.Now()
	// Use the "3:04 PM" format for a 12-hour clock with AM/PM
	formattedTime2 := currentTime2.Format("2006-01-02 3:04 PM")
	fmt.Println("Formatted Time:", formattedTime2)

	//---------------------------------------------------

	dateStr := "2023-11-25"
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err == nil {
		fmt.Println("Parsed Time:", parsedTime)
	} else {
		fmt.Println("Error parsing time:", err)
	}

	//------------------ arithmatic operations --------------------

	currentTime3 := time.Now()

	// Add 1 day
	newTime := currentTime3.Add(24 * time.Hour)
	fmt.Println("Current Time:", currentTime3)
	fmt.Println("New Time (after adding 1 day):", newTime)

	// Add 1 day formating
	newTime2 := currentTime3.Format("02.01.2006 Monday")
	fmt.Println("New Time:", newTime2)

}
