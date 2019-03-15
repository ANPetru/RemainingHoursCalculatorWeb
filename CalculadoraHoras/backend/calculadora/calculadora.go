package calculadora

import (
	"fmt"
	"strconv"
)

func GetHoursAndMinutes(hoursStr string) (days int, hour int, min float64) {
	hours := getDoubleFromString(hoursStr)
	rest := 400 - hours
	totalDays := int(rest / 8.15)
	totalHours := int(rest - float64(totalDays)*8.15)
	totalMins := 60 * (rest - float64(totalDays)*8.15 - float64(totalHours))
	return totalDays, totalHours, totalMins
}

func getDoubleFromString(s string) float64 {
	if hora, err := strconv.ParseFloat(s, 64); err != nil {
		fmt.Println(err)
		return 0
	} else {
		return hora
	}
}
