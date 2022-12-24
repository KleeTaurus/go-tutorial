package main

import (
	"fmt"
	"time"
)

func main() {
	// preDefinedFormat()
	// customizeLayout()
	// timeParsing()
	// parseInLocation()
	parseDuration()
}

func preDefinedFormat() {
	now := time.Now()
	fmt.Println(now.Format(time.ANSIC))
	fmt.Println(now.Format(time.UnixDate))
	fmt.Println(now.Format(time.RubyDate))
	fmt.Println(now.Format(time.RFC822))
	fmt.Println(now.Format(time.RFC822Z))
	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123))
	fmt.Println(now.Format(time.RFC1123Z))
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Format(time.RFC3339Nano))
}

func customizeLayout() {
	displayDate := func(format string) {
		fmt.Println(time.Now().Format(format))
	}

	// The output time will vary for you
	// but you can check the time format in which we get the output
	displayDate("20060102150405")                    // YYYYMMDDHHMMSS Format
	displayDate("2006/01/02 15/04/05")               // YYYY/MM/DD HH/MM/SS Format
	displayDate("02/01/2006 15/04/05")               // DD/MM/YYYY HH/MM/SS Format
	displayDate("02/01/06 15/04/05")                 // DD/MM/YY HH/MM/SS Format
	displayDate("2006-01-02 15:04:05")               // YYYY-MM-DD HH:MM:SS Format
	displayDate("15:04:05, 2006-Jan-02 Mon")         // HH:MM:SS, YYYY-Mon-DD DayShort Format
	displayDate("15:04:05, 2006-Jan-02 Monday")      // HH:MM:SS, YYYY-Mon-DD DayLong Format
	displayDate("15:04:05, 2006-January-02 MST Mon") // HH:MM:SS, YYYY-Mon-DD TZ DayShort Format
	displayDate("3:4:05, 2006-1-02 MST Mon")         // H:M:SS, YYYY-M-DD TZ DayShort Format
	displayDate("3:4:05 pm, 2006-1-02 MST Mon")      // H:M:SS Part_of_day(small), YYYY-M-DD TZ DayShort Format
	displayDate("3:4:05 PM, 2006-1-02 MST Mon")      // H:M:SS Part_of_day(small), YYYY-M-DD TZ DayShort Format
	displayDate("03:04:05.000")                      // HH:MM:SS.MilliSeconds
	displayDate("03:04:05.000000")                   // HH:MM:SS.MicroSeconds
	displayDate("03:04:05.000000000")                // HH:MM:SS.NanoSeconds
}

func timeParsing() {
	fmt.Println("Time parsing")

	dateString := "2014-11-09T12:40:45.15Z"
	dateString2 := "2022-10-12T10:15:16.371+07:00"

	time1, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		fmt.Println("Error while parsing date: ", err)
	}
	fmt.Println(time1)

	time2, err := time.Parse(time.RFC3339, dateString2)
	if err != nil {
		fmt.Println("Error while parsing date: ", err)
	}
	fmt.Println(time2)

	dateString3 := "2007-10-09 22:50:02.023"
	time3, err := time.Parse("2006-01-02 15:04:05.000", dateString3)
	if err != nil {
		fmt.Println("Error while parsing date: ", err)
	}
	fmt.Println(time3)
}

func parseInLocation() {
	PrintTime := func(label string, t *time.Time) {
		fmt.Println(label, t.Format(time.RFC822Z))
	}

	// Choose your layout MM YYY DD HH:SS
	layout := "02 Jan 06 15:04"
	// Format this date the Location TimeZone
	date := "09 Jun 22 19:30"
	// Load the locations to Parse
	india, inerr := time.LoadLocation("Asia/Kolkata")
	hongkong, hkerr := time.LoadLocation("Asia/Hong_Kong")
	newyork, nyerr := time.LoadLocation("America/New_York")
	if inerr == nil && hkerr == nil && nyerr == nil {
		nolocation, _ := time.Parse(layout, date)
		indiaTime, _ := time.ParseInLocation(layout, date, india)
		hongkongTime, _ := time.ParseInLocation(layout, date, hongkong)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		PrintTime("No location: ", &nolocation)
		PrintTime("India: ", &indiaTime)
		PrintTime("Hong Kong: ", &hongkongTime)
		PrintTime("New York: ", &newyorkTime)
	} else {
		fmt.Println(inerr.Error(), nyerr.Error())
	}
}

func parseDuration() {
	printfln := func(template string, values ...interface{}) {
		fmt.Printf(template+"\n", values...)
	}

	d, err := time.ParseDuration("1h30m")
	if err == nil {
		printfln("Hours: %v", d.Hours())
		printfln("Mins: %v", d.Minutes())
		printfln("Seconds: %v", d.Seconds())
		printfln("Milliseconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}
