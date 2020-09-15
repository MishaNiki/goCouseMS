package tasks

import (
	"errors"
	"math"
	"regexp"
)

// StrToInt - converting the first number in a string to a number
func StrToInt(str string) (int, error) {
	var sr []rune
	re := regexp.MustCompile(`(?m)^\D*(\d+)`)
	if strNumber := re.FindAllStringSubmatch(str, -1); len(strNumber) != 0 {
		sr = []rune(strNumber[0][1])
	} else {
		return 0, errors.New("not number")
	}
	number := 0
	for i, v := range sr {
		number += (int)(math.Pow(10, (float64)(len(sr)-i-1))) * (int(v) - 48)
	}
	return number, nil
}
