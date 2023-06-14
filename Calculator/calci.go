package main

import (
	"Calculator/calops"
	"fmt"
)

func main() {

	var v_num1 float32
	var v_num2 float32
	var v_opr  string

	for {
		fmt.Print("\nEnter Operand 1: ")
		fmt.Scanln(&v_num1)

		fmt.Print("Enter Operand 2: ")
		fmt.Scanln(&v_num2)

		fmt.Print("Enter Operation[+ - * / % or Ctrl-C]: ")
		fmt.Scanln(&v_opr)

		v_res := calops.Math(v_num1, v_opr, v_num2)

		fmt.Println("\n", v_num1, " ", v_opr, " ", v_num2, " = ", fmt.Sprint(v_res))
	}

}
