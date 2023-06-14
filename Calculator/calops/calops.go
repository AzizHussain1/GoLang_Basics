package calops

import (
	"fmt"
)

// calulates addition for 2 operands
func Math(p_num1 float32, p_opr string, p_num2 float32) float32 {

	a_symbols := [5]string{"+", "-", "*", "/", "%"}

	if !contains(a_symbols[:], p_opr) {

		fmt.Println("Invalid Operator " + p_opr)
		return 0

	} else {

		switch p_opr {
		case a_symbols[0]:
			{
				return p_num1 + p_num2
			}
		case a_symbols[1]:
			{
				return p_num1 - p_num2
			}
		case a_symbols[2]:
			{
				return p_num1 * p_num2
			}
		case a_symbols[3]:
			{
				if p_num2 == 0 {
					fmt.Println("Error: Zero Division Error.")
					return 0
				}

				return p_num1 / p_num2
			}
		case a_symbols[4]:
			{
				mod_res := int(p_num1) % int(p_num2)
				return float32(mod_res)

			}
		default:
			{
				return 0
			}
		}
	}

}

// confirms if p_opr contains desired operation symbol[+ - * /]
func contains(p_arr []string, p_symbol string) bool {

	for _, v_val := range p_arr {

		if v_val == p_symbol {

			return true
		}

	}

	return false
}

//fmt.Println(a_symbols)
