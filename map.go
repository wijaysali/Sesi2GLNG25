package main

import "fmt"

func main(){
	// var person map[string] string // deklarasi map

	// person = map[string]string{} // inisialisasi map

	// person["name"] = "Hacktiv8"
	// person["age"] = "6"
	// person["address"] = "jalan lorem ipsum"

	// fmt.Println("name:", person["name"])
	// fmt.Println("age", person["age"])
	// fmt.Println("address", person["address"])

    // cara ke 2 memberikan value beserta key secara langsung

    // var person = map[string]string{
	// 	"name": "hacktiv8",
	// 	"age" : "6",
	// 	"address" : "jalan lorem ipsum",
	// }

	// fmt.Println("name:", person["name"])
	// fmt.Println("age", person["age"])
	// fmt.Println("address", person["address"])

	//other : looping with map
	// for key, value := range person{
	// 	fmt.Println(key, ":", value)
	// } 

	// delete value

	// fmt.Println("Before deleting:", person)

	// delete(person, "age")

	// fmt.Println("After deleting:", person)

	// detect value

	// value, exist := person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("value doesn't exist")
	// }

	// delete(person, "age")

	// value, exist = person["age"]

	// if exist{
	// 	fmt.Println(value)
	// }else{
	// 	fmt.Println("Value has been deleted")
	// }


	// combine slice with map using range loop

	var people = []map[string]string{
		{"name": "Hacktiv8", "age": "6"},
		{"name": "Hacktivkidz", "age": "2"},
		{"name": "KODE", "age": "5"},
	}

	for i, person := range people{
		fmt.Printf("Index: %d, name: %s, age:%s\n", i, person["name"], person["age"])
	}

}