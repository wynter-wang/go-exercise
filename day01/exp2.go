package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)

	fmt.Println("最大值： \n", ret)
}

func max(num1, num2 int) int {
	var result int

	if(num1 < num2){
		result = num2
	} else {
		result = num1
	}
	return result
}