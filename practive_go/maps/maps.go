package main

import "fmt"

func main() {
	//                                 ------- MAPS -------------

	studentGrades := make(map[string]int)

	studentGrades["sahil"] = 100
	studentGrades["rahul"] = 90
	studentGrades["vijay"] = 60
	studentGrades["laxmi"] = 62

	// fmt.Println(studentGrades)

	// delete(studentGrades,"laxmi")
	// fmt.Println(studentGrades["laxmi"])
	// fmt.Println(studentGrades["sahil"])

	//                                    -----------check if key present --------------

	// grades,exist := studentGrades["sahil"]
	// fmt.Println(grades,exist)

	// grades1,exist1 := studentGrades["laxmi"]
	// fmt.Println(grades1,exist1)

	
	// for name , grades := range studentGrades{
	// 	fmt.Println(name,grades)
	// }

// 	for student := range studentGrades {
// 	grades, exists := studentGrades[student]
// 		fmt.Println(grades, exists)
	
// }

//                                       ----------------- another way of writing ---------------

person:= map[string]int{
	"sahil": 23,
	"bob" : 35,
	"raj": 18,
}

for name , age:= range person{
	fmt.Println(name , age)
}


}