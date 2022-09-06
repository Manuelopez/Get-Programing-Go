package unit3

/*
------ Leason 12 Functions-------
  - In Go  the functions, variables and ther identifiers that begin with an uppercase letter are expported and become available to other packages
  - a function accepts parameter
  - a function is invoked with arguments
  - function can accept multople parameters with a comma separated list
  - elliosis (...) Is can make afunction take a indefinite amout of variables 1 or more
    - variadic functions
    - interface{} known as the empty interface
    - ...interface means you cans pass any number of arguments of any type
expertiment celsiusTOfahrenheit

------- Leason 13 Methods ----------

Declarin new types
  - type keywards declares a new type with a name and a underlying type as shown
    - type celsisus float64

  - celsius tpye is a new numeric tpye with the same behavior and representation as a float64
  - you can also add values to temperature and generally use it as though it wrea a float64
  - celsius is a unique type not a type alias like rune or byte
  - if you try to use it with a float64 you'll get a msmatched types error

adding behaviour to types wth methods
  - declaring a method is as easy as declaring a function they both begin with gunc keyword and the gunction body is identical to the method body\
    - func (k kelvin) celsius() celsius {
        return celsius(k - 273.15)
      }
      - the celsisu method doesn't accept any paramtehr, but it has something like a parameter befor the name
        it's called ta reciever Method and functions can both accept multiple parameters but, but methods must have exactly one receiver
  - methods are called with dot notation
    - a celsius function that returns a celsius type isn't possible


experiment methods.go
  kelvin celsius farenheit

------- Leason 14 First Class Functions -----
  - In Go you can assign functios to variables pass functions to functions and evenwrite functions that return function.

Assigning functions to variables
  - to assign function to a variable just
    - var a = funcName
    - with no parenthesees

passing functions to other functions
  - variables refer to functon and variables can be passed to functions which means go allows you to pass function to other functions
  - to pass a function to to a function just
  - func measureTemperature(samples int, sensor, func() kelvin){
      for i := 0; i < samples; i++{
        k := sensore()
        fmt.Printf("%v K/\n", k)
        time.Sleep(time.Second)
      }
    }
  - func fakeSensor() kelvin{
       return kelvin(rand.Intn(151) + 150)
    }
  - func main(){
      measuretemperature(4, fakeSensor)
    }

Declaring function types
  - type sensor func() kelvin

Closures and anonymus functions
  - anonymus function also called a function literal in Fo is a function whothout a name unlike refular function literals are clusures because they
    keep references to variables in the surrounding scope
  - you can assign an anonymus function to a variable
    - var f - func{
        fmt.Println("hola")
      }
  - you can ecen declare and invoke an anonymous function in one step
    - func(){
      fmt.Println("hola")
    }()

  - anonymous functions can come in handy you need to create a function onthe fly. one sircumstance is when a function return from another function
    - func calibrate(s sensor, offset kelvin) sensor {
        return func() kelvin{
          return s() + offset
        }
      }
  - the anonymous function in the precedin listing makes use of closures. It references the s and offset variables that te calibrate function accepts as paramentes
    and even adter the caibrate functions returns the variables captured by closure survive so calls to sensor still have access to thos variables.
       th anonymous function encloses the variables in scopes which explains the term closure
  - Because a closure keeps a reference to surrounding variables rather that a copy of their values chages to those variables are reflected in calls to the anonymous function

experiment calibrate.go

Capstone Temperature tables
*/

import (
	"fmt"
)

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9.0 / 5.0) + 32
}

type celsius float64

func (c celsius) farenheit() farenheit {
	return farenheit((c * (9 / 5)) + 32)
}
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

type farenheit float64

func (f farenheit) celsius() celsius {
	return celsius((f - 32) * (5 / 9))
}
func (f farenheit) kelvin() kelvin {
	return kelvin(f.celsius() + 273.15)
}

type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) farenheit() farenheit {
	return farenheit(k.celsius()*(9/5) + 32)
}

func Calibrate() {
	offset := 5
	s := fakeSensor
	rfunc := aCalibrate(offset, s)
	fmt.Println(rfunc())
	offset = 10

	fmt.Println(rfunc())
	s = func() int {
		return 1
	}

	fmt.Println(rfunc())

}

func aCalibrate(offset int, s func() int) func() int {
	return func() int {
		return s() + offset
	}
}

func fakeSensor() int {
	return 5
}

func TemperatureTables() {
	drawTable(createData)
}

func createData() ([]celsius, []farenheit) {
	c := make([]celsius, 0)
	f := make([]farenheit, 0)
	c1 := celsius(-40)
	for c1 < celsius(100) {
		c = append(c, c1)
		f = append(f, c1.farenheit())
		c1 += celsius(5)
	}

	return c, f
}

func drawTable(data func() ([]celsius, []farenheit)) {
	fmt.Println("===== | C | |F| =======")
	c, f := data()
	for _, d := range c {
		fmt.Print("|", d, "|")
	}
	fmt.Print("FFFFF")
	for _, d := range f {
		fmt.Print("|", d, "|")
	}
}
