package unit6

import (
	"errors"
	"fmt"
	"net/url"
)

/*
------- Lesson 26 A few pointers -------
  - pointers directs you to a address
  - a pounter is a variable that points to the address of another variable. In computer science, ointers are a form of inderiection

The ampersand and the asterisk
  - there are two synsbos the ampersand (&) ansd the asterisk (*) tough the asterisk server sual purpose
  - & is the address operator determines the address of a variable inmemory. Variable store values in computer ram and the lcoation where a values is stored is knows as its memory address
    - a := 42
    fmt.Println(&a)
    - " prints 0x1030c319"
    - this is the location imemory where the computer sotred 42
    - you cant get the address of a luteral
  - the address operator (&) provides the memory address of a value The reverse operation is know as dereferencing which provides the value that a memory address refers to
    - a := 42
    fmt.Println(&a)
    - " prints 0x1030c319"
    ad := &a
    fmt.Println(*ad)
    - " prints 42"

pointer types

  - poubters store memory addresses
  - the address variable is a pointer of type *int as the %T format verb reveal
  - the asterisk in *int denotes that the type is a pointer in this case it can point to other variables of type int
  - pointers types can appear anywhere types are used including in variable declaraiton function parameters return types structure field typs and so ointers

Pointers are for pointing
  - is possible to derences a pointer to cahnge value of variable indirectly
  - *administrator = "hola"
  - bolden now equals "hola"
  - assiging the derenced value of pointer to another variable makkes a copu of the string

pointing to structures
  - unlike strings and numbers, composite lieterals can be prefixed with an address operator
  - it isnt necessary to dereference structures to access their fields the following listing is preferable to writing (*a).siperpower
    - a.superpower
  - go automatically derederences pointer for fields

pointing to arrays
  - as with structures composute literals for arrays  can be prefices with the address operator to create a new pointer to an arrays.
    Arrays also provide automatic dereferencing
  - composite lierals for slices and maps can slso be prefixed with the address operator but there's no automatic dereferencing

Enabling mutationo
  - Pointers are used to enable mutatuon across function and method boundaries

Pointers as parameters
  - when a pointer is passed to a function the function recieves a copy of the memory address by dereferencing the memory address, a dunction can mutate the value pointer points to

Pointer receivers
  - method reveivers are similar to parameters
  - no need create the variable with addres operator with the composite literals
    Go will automatically determine the address of a variable when aclling methods with dot n ottaion so we dont need to write
    - (&a).birthday()
  - birthday must specify a pointer receiver otherwise age wouldn't increment

Interior pointers
  - ysed to determine the memory address of a field inside a structure
  - the addres soperator in go can be used to point to field inside as tructure
    - &player.stats

Mutating arrays
  - slices tend to be preferred over arrays using rarays can be appropriate when there's no need to change their length

pointer in disguise
  - not all mutations require explicit use of a pointer. Go uses pointers behinf the scenes for some fo the built in collections

Maps are Pointers
  - leasson 19 state that maps aren't copied when assigned or passed as arguments
    maps are pointers in disguise so pointing to a map is redundant dont do this
    - func demplish(palents *map[string]string)
      - unnecessary pointer
  - it;s prefectly fine fot the jey or value of a map to be a pointer but there's rarely a reason to point to a map

Slices point at arrays
  - lesson 17 describes a slice as a window into an array to point at an elment of the array slices use a Pointers
  - a slice is represented internally as a structure with three elementsa pinter to an array the capacity of the slice and the length
  - the internal pointer allows the underlying data to be mutated when a slice is passed directly to a function or methods
  -sn rcpliviy pinyrt yo s dlivr id onu udrful ehrn mofifyin the slice itself: the length capacity or starting offset.

pointer and interfaces
  - depdens how the method is implemented

experiment turtle.go

----- Lesson 27 Much ado about nil -----
  - nil means nothing or zero
  - nil is a zero value
  - a pointer with nowhere to point has the value nil
  - nil is the zero value for slices maps and interfaces too

Nil Leads to panic
  - if pointer isn't pointer aywhere attempting to derefeecne the pointer wont work
  - derefence a nil pointer and program will crash

Nil function values
  - when a variable is dclared as a function type is value is nil by default

Nil slices
  - a slice thats delcared withou a composite lirate or the make built in will have a value nil
  - len append all wark with nil slices
  - an empty slice and a nil slice aren't quivalent but they can be used interchangeably

Nil maps
  - as wtih slices a map declared without a composite literal or the make built in function has value of nil
  - maps can be read even when nil
  - if a function oly reads from a map it fine to pass nil insteasd of making an empty mpa

Nil interfaces
  - when a variable is declared to be an interface type without an assigment the zero value is nil

An alternative to nil
 - use structs in a creative way

experiment knight.go

----- Lesson 28 To err is human -------
  - Consider the err that could occur
  - How to commiunicatge those errors
  - how to handle them
  - go keeps erro handling front and center

Handeling errors
  - function have  multiple return
    - one of them could be an errors
    - if ther was no error then error would be nil
  - when error occurs the other reutn values generally shouldn't be realid on
    - they may be set to the zero valus of their type but some function may return partial data or something else entirely

Elegant error handling
  - isolate an error free sunset of a program from the inherently error prone code
  - Error are values
  - Don't just check errors, handle them gracefully
  - Don't panic
  - make the zero value useful
  - the biffer the interface the weaker the abstraction
  - interface{} says nothing

Writing a file
  - numnber of things can go wrong
    - invalid path
    - permissions
    - creating file may file before writing
    - once writng device could run out of disk space or be unplugged
    - file must be closed when done

The deger keyword
  - To ensure that the file is closed correctly you can make use of the defer keyword. Fo ensure that all deffered actions take palce before containg function return

Creative erro handling
  - declare a new type called saffeWeirer
    - is writng to a file it sores the erro insead of retuning it subssequent attempt to write to the same filewill be skipped if writeln sees that an error ocurred previusly
  - error handeling goes inside the method writeln

New Errors
  - if something foes wrong tou can create and return new error calues to infom the caller
  - error package contqins a constructor function that acepts a string for an error message. Uing it the Set method mat create and return an out of bounds error

which error is which
  - declare and export variable for the errors that they could return

custum error types
  - Error is an interface
  - any type that implents an Error() method that returns a string will implcitly sastigu the error interface
  - type assetion comoa ok resukt check if is of types

Dont panic
  - Go has panic instead of exceptions
  - panic will run any deferred functions

keep calm and carry on
  - to keep panic from crashing the program go provides recover function
  - deferred function are excuted before a function return even in the cas of panic if a dfered dunction calls recover the panic will stop and the program will
    continue running as such revocer serves a similar purpose to catch except acnd rescue in other languges

Experiment: url.go
*/

type Sudoku [9][9]int8

func (s *Sudoku) add(val int8, i int, j int) {
	if !validRow(*s, val, i, j) {
		return
	}
	if !validColumn(*s, val, i, j) {
		return
	}
	if !validBox(*s, val, i, j) {
		return
	}

	s[i][j] = val
}

type Box struct {
	i int
	j int
}

func validBox(s [9][9]int8, val int8, i int, j int) bool {

	boxes := [9]Box{
		Box{i: 2, j: 2},
		Box{i: 2, j: 5},
		Box{i: 2, j: 8},
		Box{i: 5, j: 2},
		Box{i: 5, j: 5},
		Box{i: 5, j: 8},
		Box{i: 8, j: 2},
		Box{i: 8, j: 5},
		Box{i: 8, j: 8},
	}

	for index, box := range boxes {
		if i <= box.i && j <= box.j {
			startI := 0
			startJ := 0

			if index == 0 {
				return checkIfValid(s, startI, startJ, box.i, box.j, val)
			} else {
			  if(boxes[index-1].i == 8){
          startI = boxes[index-1].i
        }else{
          startI = boxes[index-1].i+1
        }
        if(boxes[index-1].j == 8){
          startJ = boxes[index-1].j
        }else{
          startJ = boxes[index-1].j +1
        }
				return checkIfValid(s, startI, startJ, box.i, box.j, val)
			}
		}
	}
	return false
}

func checkIfValid(s [9][9]int8, startI int, startJ int, i int, j int, val int8) bool {

        fmt.Println("aqui aqui aqui")
	for indexI := startI; indexI <= i; indexI++ {
		for indexJ := startJ; indexJ <= j; indexJ++ {

      fmt.Println(i, j)
			if s[indexI][indexJ] == val {
				return false
			}
		}
	}

	return true
}

func validRow(s [9][9]int8, val int8, i int, j int) bool {
	for index, _ := range s {
		if s[index][j] == val {
			return false
		}
	}
	return true
}

func validColumn(s [9][9]int8, val int8, i int, j int) bool {
	for index, _ := range s[i] {
		if s[i][index] == val {
			return false
		}
	}
	return true
}

func TestSudoku() {
	a := Sudoku([9][9]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
  printSudoku(a)
  a.add(1, 0, 6)
  printSudoku(a)

  fmt.Println("asodihovfiuhaerviouaehfviuhb")
  a.add(1, 0, 7)
  printSudoku(a)

}

func printSudoku(a Sudoku){
count := 1
  countj:= 1
	for _, row := range a {
		for _, v := range row {
			if count == 3 {
				fmt.Print(v, " | ")
				count = 1
			} else {
				fmt.Print(v, " ")
        count++
			}
		}
    if(countj ==3){
      fmt.Println()
      fmt.Println("---------------------------------------------------------------")
      countj = 1
    }else{
      countj++
      fmt.Println()
    }
	}



}

func Url() {
	_, err := url.Parse("https://a b.com")
	fmt.Println(err)
}

type Character struct {
	leftHand *Item
}

func (c *Character) pickup(i *Item) {
	c.leftHand = i
}

func (c *Character) give(to *Character) {
	to.leftHand = c.leftHand
	c.leftHand = nil
}

type Item struct {
	i string
}

func Knights() {
	i := Item{i: "sword"}
	a := Character{}
	f := Character{}

	e := errors.New("hello")
	fmt.Printf("%T", e)

	a.pickup(&i)
	a.give(&f)
	fmt.Println(f.leftHand.i)
}

type Turtle struct {
	X int
	Y int
}

func (t *Turtle) up() {
	t.Y++
}

func (t *Turtle) down() {
	t.Y--
}

func (t *Turtle) left() {
	t.X--
}

func (t *Turtle) right() {
	t.X++
}

func TurtleMain() {
	t := Turtle{X: 10, Y: 10}

	t.up()
	fmt.Printf("%+v", t)
}
