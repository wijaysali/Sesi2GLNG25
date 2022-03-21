package main

import (
	"fmt"
    // "strings"
)

func main(){
	// var numbers [4]int
	// numbers = [4]int{1,2,3,4}

	// var strings = [3]string{"Airell","Nanda","Mailo"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// var fruits = [3]string{"apel","pisang","mangga"}
	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	// fmt.Printf("%#v\n", fruits)


	// var fruits = [3]string{"apel","pisang","mangga"}

	//cara pertama
	// for i, v := range fruits{
	// 	fmt.Printf("Index: %d, Value: %s\n", i, v)
	// }

	// fmt.Println(strings.Repeat("#", 25))

	//cara kedua
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	// }


	// output : 

	// Index: 0, Value: apple
	// Index: 1, Value: banana
	// Index: 2, Value: mango

	// Index: 0, Value: apple
	// Index: 1, Value: banana
	// Index: 2, Value: mango

	// balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	// for _, arr := range balances {
	// 	for _, value := range arr {
	// 		fmt.Printf("%d", value)
	// 	}
	// 	fmt.Println()
	// }

	// var fruits = []string{"apple", "banana", "mango"}

	// _ = fruits

    // menggunakan function make + slice

	// var fruits = make([]string, 3)

	// _ = fruits

	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	//append

	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"grapes", "starfruit", "orange"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v", fruits1)

	// fmt.Printf("%#v", fruits)
	
   //function copy slice

    // var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"grapes", "starfruit", "orange"}

	// nn := copy(fruits1, fruits2)

	// fmt.Println("Fruits1 =>", fruits1)
	// fmt.Println("Fruits2 =>", fruits2)
	// fmt.Println("Copied elements =>", nn)

	//slicing

	// var fruits1 = []string{"apple","banana","mango","starfruit","pineapple"}

	// var fruits2 = fruits1[1:4]
	// fmt.Printf("%#v\n", fruits2)

	// var fruits3 = fruits1[0:]
	// fmt.Printf("%#v\n", fruits3)

	// var fruits4 = fruits1[:3]
	// fmt.Printf("%#v\n", fruits4)

	// var fruits5 = fruits1[:]
	// fmt.Printf("%#v\n", fruits5)

	// var fruits1 = []string{"apple","banana","mango","starfruit","pineapple"}

	// fruits1 = append(fruits1[:3],"orange")

	// fmt.Printf("%#v\n", fruits1)

	// Backing Array slice header : 
	// 1. alamat memory/address dari backing array
	// 2. panjang dari slice yang kita dapatkan fungsi len
	// 3. Kapasitas dari slice yang kita dapatkan dari fungsi cap

	// var fruits1 = []string{"apple","banana","mango","starfruit","pineapple"}
	// var fruits2 = fruits1[2:4]

	// fruits2[0] = "rambutan"

	// fmt.Println("fruits1 =>", fruits1)
	// fmt.Println("fruits2 =>", fruits2)

	// Fungsi CAP

	// var fruits1 = []string{"apple","banana","mango","starfruit","pineapple"}
	
	// fmt.Println("Fruits1 cap:", cap(fruits1))
	// fmt.Println("Fruits1 len:", len(fruits1))

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits2 = fruits1[0:3]

	// fmt.Println("Fruits2 cap:", cap(fruits2))
	// fmt.Println("Fruits2 len:", len(fruits2))

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits3 = fruits1[1:]

	// fmt.Println("Fruits3 cap:", cap(fruits3))
	// fmt.Println("Fruits3 len", len(fruits3)) 

	cars := []string{"ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Suzuki"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

}