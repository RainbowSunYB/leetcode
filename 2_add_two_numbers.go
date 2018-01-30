package main

import "fmt"

func main() {
	fmt.Println(AddTwoNumbers([]int{2, 4, 3}, []int{5, 6, 6}))
}

func AddTwoNumbers(first, second []int) []int {
	max := first
	min := second
	if len(second) > len(first) {
		max = second
		min = first
	}

	carry := 0
	for i := 0; i < len(max); i ++ {
		if i > len(min) {
			max[i], carry = AddWithCarry(max[i], 0, carry)
		} else {
			max[i], carry = AddWithCarry(max[i], min[i], carry)
		}
	}

	if carry > 0 {
		max = append(max, carry)
	}

	return max
}

func AddWithCarry(a, b, carry int) (c, returnCarry int) {
	t := a + b + carry
	if t >= 10 {
		return t - 10, 1
	} else {
		return t, 0
	}
}
