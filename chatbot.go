package main

import (
	"fmt"
	"strconv"
)

func main() {
	greetUser("Curtis", 1999)
	readName()
	guessAge()
	countNumber()
	quiz()
	fmt.Println("Congratulations, have a nice day!")
}

func greetUser( bot_name string, birth_year uint16) {
	fmt.Println("Hello! My name is ", bot_name, ".")
	fmt.Println("I was created in ", birth_year, ".")
}

func guessAge(){
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")

	// reading all remainders
	var remainder3 int
	var remainder5 int
	var remainder7 int
	fmt.Scan(&remainder3, &remainder5, &remainder7)

	//calculation for "guessing" age provided by hyperskill
	age := (remainder3*70 + remainder5*21 + remainder7*15) % 105

	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

func readName(){
	fmt.Println("Please, remind me your name.")

	var name string
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
}	

func countNumber () {
fmt.Println("Now I will prove to you that I can count to any number you want.")

var number uint
fmt.Scan(&number)
var i uint = 0

for i <= number {
	fmt.Println(i, " !")
	i++
	}
}

func quiz() {
	fmt.Println("Let's test your programming knowledge.")
	var continueQuiz bool = true
	var answer int8
	fmt.Print(`Why do we use methods?
1. To repeat a statement multiple times.
2. To decompose a program into several small subroutines.
3. To determine the execution time of a program.
4. To interrupt the execution of a program.
`)
	for continueQuiz {
		fmt.Scanln(&answer)
		if answer != 2 {
			fmt.Println("Please, try again.")
		} else {
			fmt.Println("Completed, have a nice day!")
			continueQuiz = false
		}
	}
}