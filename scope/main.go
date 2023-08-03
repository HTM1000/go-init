package scope

import "fmt"

var y int = 20

// func scope() {
// 	x := 10
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// 	PrintZ()
// }

func PrintY() {
	fmt.Println(y)
	//fmt.Println(x)
}
