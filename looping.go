package main

import "fmt"

func main(){
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// var i = 0

	// for i < 3{
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// var i = 0

	// for i < 150{
	// 	fmt.Println("Angka", i)
	// 	i++
		
	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 10;i++{
	// 	if i%2 == 1{
	// 		continue
	// 	}
	
	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// for i := 0; i < 5; i++{
	// 	for j := i; j<5; j++{
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

// 	output : 

// 0 1 2 3 4
// 1 2 3 4
// 2 3 4
// 3 4
// 4


outerLoop :
		for i := 0; i < 3; i++ {
			fmt.Println("Perulangan ke -", i+ 1)
			for j := 0; j < 3 ;j++{
				if i == 2{
					break outerLoop
				}
				fmt.Print(j, " ")
			}
			fmt.Print("\n")
		}
}


