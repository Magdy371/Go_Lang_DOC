package main

import (
	"fmt"
	"time"
)

func main() {
	var timeGoLainch time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", timeGoLainch)

	var timeNow time.Time = time.Now()
	fmt.Printf("Now created at %s\n", timeNow)

	// to add new day like tomowrrow
	tomorrow := timeNow.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow will be %s\n", tomorrow)
	//Nice format for date Time Object
	fmt.Printf("Tomorrow will be %s\n", tomorrow.Format(time.ANSIC))
}
