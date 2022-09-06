package unit4

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*

--------- Leason 16 --------

  - Arrays are aordered collection of elements with aa dixed lenth of the same type

Declaring arrays and accessing their elements

  - following plaents array constins exactly eight elements
    - var planets [8]string
  - individual elements of an array can be accessed by using square brackets []
    - with index begining at 0
  - Go has a handful of buit in functions that don't require an import statement.
    the len function can determine the length of a cariety of types/ in this case it returns size of array

Don't fo out of bounds

  - an eight elemnt arrays has indices from 0 - 7.
    - go compiler will report an errror when it detects access to an element outside of this range
  - if the go compiler is unable to derect the error you program may panic while running
  - a panic will crash your program, which is still better than modifyung memory that doesn't belong to the planets array

Initialize arrays with composite literals

  - a composite literal is a concise syntax to initialize any composite type with the values you want
    rather than declare an array and assing elements one by one Go's compositge literal syntax will declare and initialize an array in a single step
    - swafs := [5]string{"1","2","3","4","5"}
  - as a convenience you can ask the fo compiler to count the numner of elements in the composite literal by specyfing the ellipsis ... instead of a number

Iterating through arrays
  - loop
  - the range keyword provides an index and value for each element of an array with less code and less chance for mistakes
    - can use the blank identifier _

Arrays are copied
  - Assigning an aray to a new variable or passing it to a function makes a complete copy of its contents
  - arrays are values and functuos pass by value

Arrays of arrays

Experiment chess.go

---- Leason 17 Slices windows into arrays --------

  - slicing doesn't alter the array
  - it jusst creates a window or a view into the array
  - this view is a type called a slice

Slicing is and array

  - slicing is expressed with a hal-open range
    - planets [0:4]
    - beings with the planet at index 0 and continues up to but no including the planet at index 4
    - can slice an Slices
    - assaigning a new value to an element of a slice modigied the underlyinh planet array
      the change will be visible throuh the toher slices

Default indices for slicing
  - when slicing an array, omitting the first index defaults to the begining of the array.
    ommiting the last index defaults to the length of the array
  - Slice indices may not be negative
  - omiting both gives you the whole array
  - the slicing syntazx for arrays also works on strings
    - the result of slicing a string is another string. However ssing a new value to neptune wont change the value of tune or vicec versa
    - be aware that the indices indicate number of bytes not runes

Composite literals for slices
  - slicing an array is on way to create a slice, but you can also declare a slice directly. A lice of string has the type []string with no value between the brackets
    differs from n array declaration which specifies a dixed length or ellipsis between the brackets
  - d := []string{"1","2"}
    - ther is a still an underlying array behinf the scenes go declares a five element arrau then makes aslice that view all of its elements

The Power of slices
  - passing a slice to function copies the slice but the underlying arra is used

Slices with methods
  - in go you can degined a type with an underlying slice or array. Once you hace a tyoe you can attach methods to its

experiment terraform.go

---- Lesson 8 A bigger slice ------
  - Arrays have a fixed number of elements and silces just view into thos fixed arrays
  - we need variable length array that frows as needed
  - comnined slices and a bultin function anmed append
  - go proivdes the capabilities of variable length arrays

The append function
  - to add to slice use append function
  - the append function is variadic, so you vcan pass multiple elements to append in one go

Length and capacity

  - The number of elements that are visible throough a slice determines its length. if a slice has an underlying array that is larger the slice may still have capacity to grow

Investigating the append function
  - the array backking the slice doesnn't have enough room (capacity) to append "a", so appemd copues the contents of the slice to a frreshly allocated array with twice the capacity.

Theree index slicing
  - Go introduced three index slicing to limit the capacity of the resulting slice. In the next listing terrestrial has a length and capcity of 4 ceres causes a new array to be
    allocated leacing the planets array unaltered
    - planets := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
      terresrial := panets[0:4:4]
      worlds := append(terrestrial, "9")
      fmt.Println(planets)
      "Prints [1 2 3 4 5 6 7 8]"
    - iF the thrid index isn't specified terrestrial will have a capacity of 8 appending 9 doesn't allocate a new array but instead overites 5

Preallocate slices with make
  - When there isn't enough capacity for append, Go must allocate a new array and coy the  contents of the old array. You can acoid extra allocation
    and copues by preallocation a slice with the bulitin make function
  - the make function in the listing specifies both the length (0) and capcity of (10) up to 10 elements can be appended before runs out od capacity
    causing append to allocate a new arrays
    - d := make([]string, 0, 10)
  - the capacity argument is optional. to strt with length and capacity of 10 you can use make([]string, 10) Each of the 10 elements will contain
    zeri valye fir their type. the append function would add the 11th element

Declaring variadi functions
  - Printf and append are cariadic function they aaccepts a variable number of arguments. To declare a variadi function use the ellipsis ... with the alat parameter
    - func terraform(perfix string, worlds ...string) []string{}
  - to pass a slice instread of multipme arguments expand the slice with n allipsis
    - palents := []string{"1", "2", "3"}
      newPlanets = terrafirn("new", planets)
  - if terrafirn wre ti niduufy or murat elemtns of the worlds parameter palnets slice would also see those cahanges


Experiment capacity.go

------ Lesson 19 the ever versatile map -------
  - Maps come in hadny when you are searching for something go provides a map collection with jey that map to values
    where as arrays and slices are indexed by sequetial integers, map keys can be nearly any type
  - usefull for unstructured data wwhere keys are determined while program is running

Declaring a map
  - the keys of maps can be nerly any type
  - mustly specify type for the keys and valuex
  -s to declare map with keys of type sting and values of type int the suntax is
    - map[string]int
  - use brakets to look up values by key, to assing over existing values or to add values to the map
  - composite lieteral
    - temerature := map[string]int{
        "1": 1
        "2": 2
      }
      temp := 53MP34Q5U43["1"]
      fmt.Println(temp)
      "prints 1"
      temperature["1"] = 3
      fmt.Println(temperature)
      "prints map[1: 2, 2"2]"
  - if you acces a key tat doesnt exist int hte  map, the result is the zero value for the type
  - go prives the como, ok suntax which you can destinguis between the key not existing in the map versus being  present in the map
    - if t, ok := temeprature["3"]; ok{
        fmt.Println("yay")
      }
      else{
        fmt.Println("nay")
      }
      "prints nay"

  - the ok variable will be true if the key is presen or false otherwise

Maps arent copied
  - when passing maps to a function it will alter the original map
    - same with slices
  - delete function removes an element from a map
    - delete(temperature, "1")

Prealocating maps with make
  - Unless you intialize them with a composite literal, maps need to be allocated with th emake built in function
  -for maps, mae only accepts one ro two parametes, The second one preallocates pace for a number of jeys much like capcity for slices
    - map'sintial length will always be zero when using make

using maps to count things
  - frequence counter

grouping data with maps and slices
  - truncate data
  - keys pount to slice value
     add values to the trucated value like 30: [33, 31]

Reourposing maps as sets
  - a set is a collection similar to array except that each elemnt is guaranteed to occur only once
  - go doesn't privede aset collection
  - we can improcise by using maps
  - value isn't imporatant , but true is conecniet for cheking set mebership
  - map only contains one key for each temperature any duplocated are removed
  - keys have an arbitrary oeder in go
  - before can be sorted map must be converted back to slice

experiment words.go





*/

func Chess() {
	var board [8][8]rune
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'
	board[1][0] = 'p'
	board[1][1] = 'p'
	board[1][2] = 'p'
	board[1][3] = 'p'
	board[1][4] = 'p'
	board[1][5] = 'p'
	board[1][6] = 'p'
	board[1][7] = 'p'

	board[7][0] = 'r'
	board[7][1] = 'n'
	board[7][2] = 'b'
	board[7][3] = 'q'
	board[7][4] = 'k'
	board[7][5] = 'b'
	board[7][6] = 'n'
	board[7][7] = 'r'
	board[6][0] = 'p'
	board[6][1] = 'p'
	board[6][2] = 'p'
	board[6][3] = 'p'
	board[6][4] = 'p'
	board[6][5] = 'p'
	board[6][6] = 'p'
	board[6][7] = 'p'

	for _, row := range board {
		for _, p := range row {
			fmt.Printf("| %c |", p)
		}
		fmt.Println()
	}

}

type Planets []string

func (p Planets) terraform() {
	for i := 0; i < len(p); i++ {
		p[i] = "New " + p[i]
	}
}

func Terraform() {
	p := Planets{"Mars", "Uranus", "Neptune"}
	p.terraform()
	fmt.Println(p)

}

func Capacity() {
	s := []string{}
	c := cap(s)
	l := len(s)
	for {
		s = append(s, "1")
		c = cap(s)
		l = len(s)
		fmt.Println("Capcity:", c, "Lenth:", l)
		if c > 200 {
			break
		}

	}
}

func Words() {
	s := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man"

	a := strings.Split(s, " ")
	m := map[string]int{}
	for _, w := range a {
		cw := strings.ToLower(strings.Trim(w, ".;,"))
		if _, ok := m[cw]; !ok {
			m[cw] = 1
		} else {
			m[cw]++
		}

	}

	fmt.Println(m)
}

const square int = 10

func GameOfLife() {
	a := [square][square]int{}

	for i, row := range a {
		for j, _ := range row {
			if v := rand.Intn(100) + 1; v <= 25 {
				a[i][j] = 1
			}
		}
	}

	for i := 0; i < 60; i++ {
    printArray(a)
		a = calculalteNextGeneration(a)
		time.Sleep(time.Second)
	}

}

func calculalteNextGeneration(a [square][square]int) [square][square]int {
	newGen := [square][square]int{}
	for i, row := range a {
		for j, _ := range row {
			neighbors := countNeighbors(i, j, a)
			if neighbors == 2 || neighbors == 3 {
				newGen[i][j] = 1
			} else {
				newGen[i][j] = 0
			}
		}
	}
	return newGen
}

func countNeighbors(i int, j int, a [square][square]int) int {

	c := 0
	//x
	v := 0
	v2 := 0
	v = (i - 1) % square
	if v < 0 {
		v += square
    
	}
	if a[v][j] == 1 {
		c++
	}

	if v := (i + 1) % square; a[v][j] == 1 {
		c++
	}

	//y
  v = (j-1) % square
  if v < 0{
    v += square
  }
	if a[i][v] == 1 {
		c++
	}

	if v := (j + 1) % square; a[i][v] == 1 {
		c++
	}

	// diagonal
  v = (i-1) % square
  if v < 0{
    v += square
  }
  v2 = (j-1) % square
  if v2 < 0{
    v2 += square
  }
	if a[v][v2] == 1 {
		c++
	}
  v = (i+1) % square
	if a[v][v2] == 1 {
		c++
	}
  v = (i -1) % square
  if v < 0{
    v += square
  }
  v2 = (j+1)% square
	if a[v][v2] == 1 {
		c++
	}
  v = (i+1) % square 
	if a[v][v2] == 1 {
		c++
	}
	return c
}

func printArray(a [square][square]int) {
	for _, row := range a {
    fmt.Println("")
		for _, c := range row {
			if c == 0 {
				fmt.Print("| |")
			} else {
				fmt.Print("|C|")
			}
		}
    fmt.Println("")
	}
}
