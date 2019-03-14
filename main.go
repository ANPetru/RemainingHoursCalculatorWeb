package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("horas hechas")
	hor, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		horas := getDoubleFromString(hor)
		rest := 400 - horas
		tDias := int(rest / 8.15)
		rHoras := int(rest - float64(tDias)*8.15)
		fmt.Printf("Quedan %d dias y %d horas\n", tDias, rHoras)
	}
}

func getDoubleFromString(s string) float64 {
	str := ""
	for i := range s {

		num := int(s[i])
		if num >= 48 && num <= 57 {
			str += string(num)
		} else if s[i] == '.' {
			str += "."
		}
	}
	hor, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		return hor
	}
	return 0
}
