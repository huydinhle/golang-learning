package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	time := readTime()
	result := convertTimeStr(time)
	fmt.Printf("%+v\n", result)
}

func convertTimeStr(time string) string {
	hour, minute, second, m := breakTime(time)
	result := convertTime(hour, minute, second, m)
	return result
}

func convertTime(hour, minute, second int, m string) string {
	var newHour int
	var newMinute int
	var newSecond int
	if hour == 12 && m == "AM" {
		newHour = 00
	} else if hour == 12 && m == "PM" {
		newHour = 12
	} else if m == "PM" {
		newHour = hour + 12
	} else {
		newHour = hour
	}
	newMinute = minute
	newSecond = second

	return fmt.Sprintf("%02d", newHour) + ":" + fmt.Sprintf("%02d", newMinute) + ":" + fmt.Sprintf("%02d", newSecond)
}

func breakTime(time string) (int, int, int, string) {
	sliceString := strings.Split(time, ":")
	sub1 := sliceString[2][:2]
	sub2 := sliceString[2][2:]

	hour, _ := strconv.ParseInt(sliceString[0], 10, 64)
	minute, _ := strconv.ParseInt(sliceString[1], 10, 64)
	second, _ := strconv.ParseInt(sub1, 10, 64)

	return int(hour), int(minute), int(second), sub2
}

func readTime() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	time := scanner.Text()
	return time
}
