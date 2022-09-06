package unit2

/*
-------Leasson 6----------
      Real Number
"Computer store mutltiple real number like 3.1415 using IEEE 754 floating point standard. Floatiing point numbers can be very large or incredibly small"
"Computer also support integers for whole numbers"

Declaring floating-point variables
  -Go compuler will infer that days is a float64
  - days := 365.2425*
  - var days = 365.2425
  - var days dlot64 = 365.2425

Single precision floating point numbers
  - Go has two floating types
    - float64 a 64 bit floating point type that uses eight bytes of memory
    - float32 type uses hal the memory of float64 but offers less precission. This type is sometimes called single precission

The zero value
  - each types has a default values, called the zero value. The degual applies when you declare a variable but don't intialize it with a values
  - var price float64
    - price is 0

Displaying floating point types
  - When uses Printf with the %f formatting verb to specify the number of digits
    - thid := 1.0 /3
    - fmt.Printf("%f\n", third)
        - prints 0.333333
    - fmt.Printf("%.3f\n", thrid)
      - prints 0.333
    - fmt.Printf("$4.2\n", third)
      - prints 0.33
    - 4.2 is width and precission
      - if width larger than number is will fill up with space
    - to a left-pad with zeros instead of spaces prefix the width with a zero as
      - fmt.Printf("$05.2\n", third)
        - prints 00.33

floating point accuracy
  - int mathematics some ratianl numbers can't be accuratly represented in decinal form. the number 0.33 is only an approximation of 1/3
  - floating-point numbers suffer from rounding erros too, except floating point uses a binary representation indested od decimal the consequnce is that computers can accuratly represented
    1/3 but have rounding erros with other numbers as the following listing illustrates
  - to minimiez rounding errors we recinnebd tgat you preform multiplication before division.

Comparing floating point numbers
  - piggyBank contained 0.3000000004, ranther than the desired 0.30
  - instead of comparing floating point numbers directly determine the absolute difference between two numbers adn then ensure the difference isn't to big
    - math.Abs
  - the upper bould for a floating point error for single operation is know as the machine epsilon which is 2^-52 float64, and 2^-23 for float32

---- Leason 7------
  Whole Numbers

Go offers 10 diffrent types for whole numbers collectively called integers. Integers don't suffer from the accuracy issues of floating point types
but they can't store fractional numbers and they have limited range.

Declaring integer variables
  - Five integer types are sigened meaning they can represent both postive and negative whole numbers. The most common one is a signed Integers
    - int
  - the other five integer types are unsigned maning the're for postive numbers only. The abbreviation for unsigend integers i uint
    - uint


the unit8 type for 8 bit colors
  - hexadecimals requrires a 0x prefix
  - %x ir $X format verbs ti dusokay bumbers in hexadecimal

Intgeres wrap around
  - unit8 ha a range 0-255. Increamenting beyond 255 will wrap back to 0
  - when the integer goes over the max it wraps around
  - math.MaxUint16 and similar constant for each architecture indepdendent type
  - %b format verb outputs integers in base2

Experimenet piggy.go


----- Lesson 8. Big numbers --------
  - 3 types
    - big>int is for big integers
    - bg.Float is for arbitrary precision floating point numbers
    - big.Rat is for fractions like 1/3
  -constant are delcated with the const keyward but every literal value in toy program is a constant too that means unusually sized numbers can be used directly 
  - calculation on constants and literals are preformed during compilation rather than while the program is running. the go compiler is written in go. Under theh hood 
    untyped numeric constatns are backed by the big package enbaling all the usual operations with number well beyond 18 quintillion 

experiment cains.go

----- Leason 9 multilingual text --------

Declaring string variables
  - Literal values wrpaped in quates are ingerred to be of type stribg
  - uf you declare a variable without providing a value it will be initialized with the zero value for its type 
    - empty string or ""

raw string literals
  -`` backtics indicat a raw string literal

Characters code points runes and bytes
  - runes is an alias for int32
  -a byte is an alias for the int uint8 type. it is inteded for binary data, otuh can be used for english characters degined by ASCII an older 128 character subset of unicode
  - an alias is another name for the same type, so rune and int int32 are interchangeable. though byte and rune have been in Fo from the beginning
  - both byte and rune behave like the integer types they are aliases
  - to dsplay charaters rather tan their numeric values, th %c format verb can be used with Printf 
  - rather than memorize unicode code points Go provides a character literal. Just enclose a character in single quotes '' fo will inger a rune
    - character lieterals can also be used with the byte alias

Pulling the strings
  - strings cant but mutated
  - orogram can acess individual character , but it can't alter the characters of a string, uses the brackets[] syntax
  - string are inmutable

Mnupulating characters with Caesa cipher

What is the result of the expression c= c - 'a' + 'A' if c is a lowercase 'g'
  - 'G'

Deoding strings into runes
  - go language provdes the range keyword iterate  over a variety of collection it can also decode encoded string UTF-8 enoded strings

experimetn caesar
experiment international.go
 

----- Leason 10 Converting between types -------
  - types don't mix
  - to convert string to number use Atoir in strconv, strconv.Ataoi
    - will return error if the string doesn't contain a valid number

Numeric type conversions
  -to convert use the type and wrap aorund a variable eg
    - a := 1
    - float64(a)
  - it always safe to convert to a type with a larger range, other integer conversion come with risks

convert types with caution
  - min/max for each type in  math package
  -

string conversions
  - strconv.Iota num to string
  - another way to convert a number to a string is to use Sprintf, a cousin printf that returns a string rather than displaying it
 
converting Boolean values
The Print famiily of functon displays the boolean values true and false as text,


Leason 11 capstone The cigenere cipher
 TODO READ THE FUKING CAPSTONE

*/

import (
	"fmt"
	"math/rand"
)

func Piggy() {
 	var piggy float32 = 0

	for {
		switch r := rand.Intn(3); r {
		case 0:
			piggy += 0.05
		case 1:
			piggy += 0.10
		case 2:
			piggy += 0.25
		}

    fmt.Printf("$%.2f\n", piggy)
    if piggy >= 20{
      break
    }


	}
}

func Canis(){
  fmt.Println(236_000_000_000_000_000 / 9.462e12)
}

func Caesar(){
  s:= "L fdph, L vdz, L frqtxhuhg."
  
  for _, c := range s {
    if c =='a'{
      c += 23
    }else{

    c -= 3
    }
    fmt.Printf("%c", c)
  }
}

func International(){
  s := "Hola Estacion Espcail Internacional"

  for _, c := range s{
    c += 13
    fmt.Printf("%c", c)
  }
}
