package main

import (
	common "Practica1/calculadora/pkg/common"
	ops "Practica1/calculadora/pkg/ops"
	"fmt"
	"strconv"
)

//f, err := strconv.ParseFloat("3.1415", 64)

func main() {
	for {

		menu()
		ope := common.ReadValue()
		execCop(ope)
	}
}

func menu() {
	fmt.Println("Elige una operacion a realizar")
	fmt.Println(`	1.- Suma
	2.- Resta
	3.- Multiplicacion
	4.- Division`)
}

/*func readValue() string {
	var value string
	fmt.Scanf("%s", &value)
	return value
}*/

func execCop(value string) {
	fmt.Println("Ingresa el primer numero:")
	x, err := strconv.ParseFloat(common.ReadValue(), 32)
	fmt.Println("Ingresa el segundo numero:")
	y, err1 := strconv.ParseFloat(common.ReadValue(), 32)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	if err1 != nil {
		fmt.Println(err1.Error())
		panic(err1)
	}

	var res float32
	switch value {
	case "1":
		res = ops.Suma(float32(x), float32(y))
		break
	case "2":
		res = ops.Resta(float32(x), float32(y))
		break
	case "3":
		res = ops.Multiplicacion(float32(x), float32(y))
		break
	case "4":
		res = ops.Division(float32(x), float32(y))
		break
	}
	fmt.Println("El resultado es: ", res)
}
