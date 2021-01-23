package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)

	valor1 := parsear(entradaLimpia[0])
	valor2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		return valor1 + valor2
	case "-":
		return valor1 - valor2
	case "*":
		return valor1 * valor2
	case "/":
		return valor1 / valor2
	default:
		fmt.Printf("Operacion %s incorrecta\n", operador)
		return 0
	}
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parsear(entrada string) int {
	valor1, _ := strconv.Atoi(entrada)
	return valor1
}
