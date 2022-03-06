package tool

import "strings"

func GetRealDecimalValue(value string,decimal int) string{
	if strings.Contains(value,".") {
		arr :=strings.Split(value, ".")
		if len(arr) != 2 {
			return ""
		}
		num := len(arr[1])
		left := decimal - num
		return arr[0]+arr[1]+strings.Repeat("0",left)
	}else{
		return value+strings.Repeat("0",decimal)
	}
}

