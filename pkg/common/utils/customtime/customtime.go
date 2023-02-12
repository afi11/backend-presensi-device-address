package customtime

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func TimeNowHuman() string {
	now := time.Now()
	var timeNow = fmt.Sprint(now.Hour(), ":", now.Minute(), ":", now.Second())
	return timeNow
}

func SplitInputTime(timeInput string) []string {
	result := strings.Split(timeInput, ":")
	return result
}

func CalculateTime(timeInput1 string, timeInput2 string) time.Duration {
	now := time.Now()
	getTime1 := SplitInputTime(timeInput1)
	getTime2 := SplitInputTime(timeInput2)

	hour1, _ := strconv.Atoi(getTime1[0])
	minute1, _ := strconv.Atoi(getTime1[1])
	second1, _ := strconv.Atoi(getTime1[2])

	hour2, _ := strconv.Atoi(getTime2[0])
	minute2, _ := strconv.Atoi(getTime2[1])
	second2, _ := strconv.Atoi(getTime2[2])

	var time1 = time.Date(now.Year(), now.Month(), now.Day(), hour1, minute1, second1, 0, time.Local)
	var time2 = time.Date(now.Year(), now.Month(), now.Day(), hour2, minute2, second2, 0, time.Local)

	var duration = time1.Sub(time2)

	return duration

}

func FormatCorrectTime(timeInput time.Duration) string {
	hour := int(timeInput.Minutes() / 60)
	minute := int(math.Mod(timeInput.Abs().Minutes(), 60))

	var result = fmt.Sprint(hour, ":", minute, ":00")
	return result
}
