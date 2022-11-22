package numbers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	hex = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	bin = []string{"0", "1"}
)

func Convert(ns NumberSystem, value string) (int, error) {
	switch ns {
	case BinaryNumberSystem:
		return binary(value)
	case DecimalNumberSystem:
		return decimal(value)
	case HexadecimalNumberSystem:
		return hexadecimal(value)
	case UnknownNumberSystem:
		fallthrough
	default:
		return -1, fmt.Errorf("unknown conversion")
	}
}

func BinaryString(value int) string {
	var calc []int
	remainder := 0
	for {
		dividend := value / 2
		remainder = value % 2
		calc = append(calc, remainder)
		if dividend == 0 {
			break
		}

		value = dividend
	}

	result := make([]string, len(calc))
	for index, s := range calc {
		result[len(calc)-index-1] += bin[s]
	}

	return fmt.Sprintf("0b%v", strings.Join(result, ""))
}

func DecimalString(value int) string {
	return fmt.Sprintf("%v", value)
}

func HexadecimalString(value int) string {
	var calc []int
	remainder := 0
	for {
		dividend := value / 16
		remainder = value % 16
		calc = append(calc, remainder)
		if dividend == 0 {
			break
		}

		value = dividend
	}

	result := make([]string, len(calc))
	for index, s := range calc {
		result[len(calc)-index-1] += hex[s]
	}

	return fmt.Sprintf("0x%v", strings.Join(result, ""))
}

func binary(value string) (int, error) {
	value = strings.TrimPrefix(value, "0b")

	var converted []string
	converted = append(converted, strings.Split(value, "")...)

	total := 0
	for index, s := range converted {
		if s != "0" && s != "1" {
			return 0, fmt.Errorf("value is not binary")
		}
		d, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}

		pow := math.Pow(float64(2), float64(len(converted)-index-1))
		total += int(d) * int(pow)
	}

	return total, nil
}

func decimal(value string) (int, error) {
	return strconv.Atoi(value)
}

func hexadecimal(value string) (int, error) {
	value = strings.TrimPrefix(value, "0x")
	if len(value)%2 != 0 {
		value = "0" + value
	}

	var converted []string
	converted = append(converted, strings.Split(value, "")...)

	total := 0
	for index, s := range converted {
		d, err := strconv.Atoi(s)
		if err != nil {
			if strings.EqualFold(s, "A") {
				d = 10
			} else if strings.EqualFold(s, "B") {
				d = 11
			} else if strings.EqualFold(s, "C") {
				d = 12
			} else if strings.EqualFold(s, "D") {
				d = 13
			} else if strings.EqualFold(s, "E") {
				d = 14
			} else if strings.EqualFold(s, "F") {
				d = 15
			} else {
				return -1, fmt.Errorf("value is not hexadecimal")
			}
		}

		pow := math.Pow(float64(16), float64(len(converted)-index-1))
		total += int(d) * int(pow)
	}

	return total, nil
}
