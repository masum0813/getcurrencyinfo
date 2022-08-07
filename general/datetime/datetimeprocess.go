package datetime

import (
	"fmt"
	"time"
)

func GetDate(varDatetime string) string {
	varTime, _ := time.Parse(time.RFC3339, varDatetime)
	dayDiff := 0

	if varTime.Weekday() == time.Saturday {
		dayDiff = -1
	} else if varTime.Weekday() == time.Sunday {
		dayDiff = -2
	}

	varTime = varTime.AddDate(0, 0, dayDiff)

	return varTime.Format(time.RFC3339)

}

func GetTime(varDateTime string) string {
	varTime, _ := time.Parse(time.RFC3339, varDateTime)

	if varTime.Hour() < 10 {
		return "1000"
	} else if varTime.Hour() > 15 {
		return "1500"
	} else {
		return fmt.Sprintf("%d00", varTime.Hour())
	}

}

func GetCustomRequestDate(vDate string) string {
	varTime, _ := time.Parse(time.RFC3339, vDate)
	retVal := fmt.Sprintf("%02d.%02d.%d", varTime.Day(), varTime.Month(), varTime.Year())
	return retVal
}

func GetCustomRequestTime(vDate string) string {
	varTime, _ := time.Parse(time.RFC3339, vDate)
	retVal := fmt.Sprintf("%02d:%02d", varTime.Hour(), varTime.Minute())
	return retVal
}
