package main

import (
	"bufio"
	"fmt"
	"os"

	calculadora "github.com/Project1/CalculadoraHoras/backend/calculadora"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Horas hechas:")
	hours, _, err := buf.ReadLine()
	if err != nil {
		fmt.Println(err)
	} else {
		totalDays, totalHours, totalMin := calculadora.GetHoursAndMinutes(string(hours))
		fmt.Printf("Quedan %d dias, %d horas y %.2f minutos\n ", totalDays, totalHours, totalMin)
	}
}
