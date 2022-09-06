package unit1

/*

-------Leason 1---------

What is go?
  - Compiled Lang

  what are two benegits of the Go compiler?
    - arge programs compile in seconds and the go compuler catch typos and mistakes begore unning a programs

Packages and Functions
  - Three Go Keywords
    - package
      "The package keyword declatres this code to belogns to. All code in GO is oprganized into packages. Go Providfes a standard library"
    - import
      "Specifies packages this code will uses. Packages contains any number of function and variables"
    - func
      "declaes a function the body of each function is enclosed in curly braces which is how fo know where each function begins and ends"

  - main identifier is special when you run a program written in Go execution begins at the main function in the main package. WIthout main, the Go Compiler will report an error
    Because it doesnt know where the program should start
  - Println function to print a line of text
    - prefixed with fmt followd by dot because it is probided by the fmt package
  - Where does a Go program start
    - at the main function in the main package
  - What does fmt package Providfes
    - The fmt package provides function for formatted input and output

The one true brace style
  - an opening brace must be on the same line as func rather than on an sperate line

- we used 3 of the 25 go reserved keywords

--------Leason 2------------

Arithmetic Operator
  - +, -, *, /, %

  - comment is //
  - Print function can pass list of arguments separated by commas
    - can be text a number or a mathematical expression

Formated Print
  - Printf  you can inservalues aywhere in the text
    - Thtext cintains the format verb %v which is substituted for th value the expression provided by the second argument
    - does not put a new line needs \n
    - if multiple format verbes are specified will sunstitute multiple values in order
    - %4v give the 4 haracter of width

- how do tou print a new line?
  - use \n anywhere in tet youre printing to insert a new line or use fmt.Println()j

- What does printf do when it encounters the %v format verb
  - The %v is sunstituted for a value from the following arguments

Constants and variables
  - const
    - create a constant variable that cannot be reassigned
  - var
    - creates a variable that can be reassigned

  - Declare multiple variables at once
    - var (
      distance = 5
      speed = 4
    )
    - var distance, speed = 5, 4
Increment and assigment operators
  - ++
  - +=
  - --
  - /=

Think of a number
  - can generate oseudorandom numbers using the rand package
  - rand.Intn(10) return a nuber fom 0 to 9

  - write a program that generates a random distance from 56 t0 401,000,000
    - look at func RandomNumber

Sumary
  - Used 5 of the 25 go jetwords, package, import, func ,const, and var

Experiment
  malacandra.go write a pgroram to determine how fast a ship would need to travel in km/h in order to reach malacandra in 28 days assume a distance fo 56,000,000

----- Leason 3 --------
  Loops and branches

True or False
  - Tue and false are boolean values

COmparisons
  - another way to arrive to a true or galse values is by comparing two values go provides a comparison operatoes show
    - ==
    - <
    - <=
    - !=
    - >
    - >=

Branching with if
  - A computer can use boolean values or comparisons to choose between different paths with an if statement as show in th followin listing
    func main(){
      var command = "go easgt"
      if command == "go east"{
        fmt.Println("you head further up the mountain")
      } else if command == "go inside"{
        fmt.Println(?you enter the cave where you liv e out the rest of your lige)
      }else{
        fmt.Println("didnt quite get that")
      }
    }

    - Prints out You head further up the mountain
    - Both else if and else are optional. when there are sevberal paths to consider you can repeat else if as many time as needed

Logical operators
  - in go the logical operator || means "or"
  - the logical operator && means "and"
  - logical ! not operator

  - go uses short circuit logi if the fitst condition is true thers no need to follows the || operator so it is ifnored
  - the && operator is gust the opposite the result is false unless both condition are true

Branching with switch
  - Go Provides thw swithc statment to compare one value to serveral others
  - Switch has the fallthorugh keywarod whic is used to execute the body of the next case

repetion with loops
  - The for jeywird reoaest code for you given the codition is  meth
  - for can be used as a while loop
  - for can be used as an infinite loop that uses break

---- leason 4 -----
  Variable Scope

looking into Scope
  - when craible si declared it comes into Scope
  func main(){
    var count = 0
    for count < 10 {
      var num = rand.Intn(10) +1
      fmt.Println(num)
       count++
    }
  }
  - Thc ount variable is declared within the funciton scope and is visible until the end of the main function
  - the num cariable is declared within the scope of the for loop after the loop ends the num vairbal foes out of scope

Short declaration
  - short declaration provres an alternative syntax for the var keyword
  - var count = 10
  - count := 10
  - makes it posoble to declare a new variable in an if statement or in a for loop or switch statment

capstone tickeet to mars



*/

import (
	"fmt"
	"math/rand"
)

func Playgorund() {
	fmt.Println("Hello, world")
	fmt.Println("Hello, 世界")
}

func RandomNumber() {
	fmt.Println(rand.Intn(401000000-56000000) + 56000000)
}

func Malacandra() {
	//distance/time = speed

	distance := 56000000
	time := 28 * 24

	fmt.Println(distance / time)
}

func GuessTheNumber() {

	var s string
	var max = 100
	var min = 0
	var guess int
	var size string
	for {
		guess = rand.Intn(max-min) + min

		fmt.Println("Is this your number?:", guess)

		fmt.Scanln(&s)
		if s == "y" {
			break
		} else {
			fmt.Println("Bigger or smaller?")
			fmt.Scanln(&size)
			if size == "b" {
				min = guess
			} else {
				max = guess
			}
		}

	}
}

func TiketToMars() {
	fmt.Printf("%20v %5v %12v $ %7v\n", "SpaceLine", "Days", "Trip Type", "Price")
	fmt.Println("================================================")
	for count := 0; count < 10; count++ {
		option := rand.Intn(3)
		spaceline, days, roundOrOne, cost := getValues(option)
		trip := "One-Way"
		if roundOrOne {
      trip = "Round-Trip"
		}
		fmt.Printf("%20v %5v %12v $ %7v\n", spaceline, days, trip, cost)

	}
}

func getValues(option int) (string, int, bool, int) {

	spaceline := "Space Adventures"
	if option == 1 {
		spaceline = "Space X"
	} else if option == 2 {
		spaceline = "Virgin Galatic"
	}
	distance := 62100000
	speed := rand.Intn(30-16) + 16
	var roundOrOne bool
	if a := rand.Intn(2); a == 0 {
		roundOrOne = true
	} else {
		roundOrOne = false
	}
	cost := speed + 20
	days := distance / speed / 86400 
	if roundOrOne {
		cost += 20
	}

	return spaceline, days, roundOrOne, cost

}
