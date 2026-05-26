package main

func main() {
		 Operations()

	// =========================
	// Learned:
	// - Variables in Go
	// - Data types
	// - Constants
	// - Short variable declaration
	// - Printing output
	// =========================

	// var name = "hi sahil"
	// var number = 54
	// var number1 = 25.123
	// var number2  = "C"
	// var boy = true
	// const game = 34

	// fmt.Println(number)
	// fmt.Println(name)
	// fmt.Println(number1)
	// fmt.Println(number2)
	// fmt.Println(boy)
	// fmt.Println(game)

	// number = 98

	// SAHIL :=67
	// fmt.Println(SAHIL)

	// =========================
	// Learned:
	// - Taking input using Scanln
	// - Multiple input values
	// // =========================

	// fmt.Println("what is your name?")
	// var name,name2 string
	// '&' gives memory address of variables
	// fmt.Scanln(&name,&name2)

	// fmt.Println("I am",name,"and",name2)

	// =========================
	// Learned:
	// - Using bufio package
	// - Reading full line input
	// - Using os.Stdin
	// - Ignoring errors using '_'
	// // =========================

	// reader := bufio.NewReader(os.Stdin)

	// // ReadString reads input until '\n'
	// name, _ := reader.ReadString('\n')

	// fmt.Println("hello", name)

///////////////////////////////////////////////////////////using functions
// ans :=add(2.2, 4.54)
// fmt.Println(ans)

// ans,k:=divide(10,0)
// fmt.Println(ans , k)

}

// func add(a,b float32)(result float32){  // here the output format is written after the fucntion 
// 	return a+b
// }

// func divide(a,b float32)(float32 , string){  // can be returned as (return type , error) or (return type , string/or anything)
// 	if(b==0){
// 		return 0 , "Not divisible by 0"
// 	}
// 	return a/b ,""
// }