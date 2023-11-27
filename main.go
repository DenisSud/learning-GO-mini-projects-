package main

import "fmt"

func main() {

	// //stritngs
	// var hi string = "hi"
	// var hello = "hello"
	// var ok string
	// ok = "ok"

	// fmt.Println(hi, hello, ok)

	// nameour := "I'm good"
	// fmt.Println(nameour)

	//ints
	// var age1 int = 20
	// var age2 int8 = 12
	// var int3 uint8 = 255
	// fmt.Println(age1, age2, int3)

	//floats
	// var float1 float32 = 3.0120398420987418273401326457634771265781234642134
	// var float2 float64 = 3.0120398420987418273401326457634771265781234642134
	// fmt.Println(float1, float2)
	// print
	// age := 20
	// name := "Nick"
	// // fmt.Printf("hi, I'm %v \nI'm %v\n", name, age)
	// var str = fmt.Sprintf("hi, I'm %v \nI'm %v", name, age)
	// fmt.Println(str)

	// arrays
	var array = [3]int{1, 2, 1}
	names := [3]string{"Nick", "John", "Bob"}
	fmt.Println(array, names)

	//Slices
	var score = []int{1, 2, 3}
	score = append(score, 4)
	fmt.Println(score, len(score))
	r1 := score[1:3] //including score[1] but not score[3]
	r2 := score[2:]
	r3 := score[:3]
	fmt.Println(r1, r2, r3)
	
}