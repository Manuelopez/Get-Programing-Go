package unit5

/*

------ Leason 21 A little structure ---------
  - Go provides a structure typ

Declaring a structure
  - location structure represent a location
    - var curiosity struct {
        lat float64
        long float64
      }
    - To assign a flied or access the value of a filed use dot notation with variable anme dot field name
      - curiosity.lat = -4.5
    - The pRint famuly of dunction will display the contents of structures for you

reusing structures with typs
  - if you neeed multiple structures with the same fields you can degine a type much like thecelsius type. The location typ declared in the following listing
    - type location struct{
        lat float64
        long float64
      }
initialize structures with composite literals
  - composite literals for initializeing strutures come in two different forms
    - field value pairs, Flieds may be in order and fields hat arent listd will retain the zero value for their type
    this form tolarates change and will continue to work correctly even if fields are added to tjhe structure or if fields are reodderd
    - if locationgained an altitude field, both the set var would defaul to an altitude of zero
    - opportunity := location{lat: -1, long 354}
  - The composite literal can be without names, instread a value must be provided for each dield in the same order which thy're listed in the sturcture
    this form works best for types that are stable and only have a dew fields. if the location type gains altitude field you must specify altitude for the program to compile
    - spirit := location{-1, -2}
  - %+v to print out the field names

Structure are copied
  - when assaigning a variable to a struct variable the item are copied so the values change idepently

A slice of structures
  - []struct is a collection of zero or more values where each valus is baed on a structure

Encoding structures to json
  - Marshal function from the json package to encode the data in location Json format  marshal returns json data as bytes which can be sent over the wire or
    converted t0 a string for display it ay also return an error
  - if properties began with loweracae the output {}

Customizing JOSOn with struct tags
  - Go json package requires that fields have an intial uppercase letter and multiword field names use CamelCase by convetion.
  - you may want to use snake_case
  - the fields structure can be tagged with field names you want the kson package to use
  -  struct tags to make it json `json:"latitutde:`
  - json.Marshal makes it to a []bytes use string() to convert to string

Experiment landing.go


------ Lesson 22 Go's got no class ----------
  - Go has no classes
  - go still provides what you need to apply ideas from object oriented deseing

Attaching methods to sturctures
  - can attach methods to sturct or any other type

Constructur functions
  - classicual languages orovide as a specual feature to construct object.
  - Go doesn't have a language feature for ocnstructors uinstread just use and ordinary ufnction
  - if needed crate multiple constuctor functions with difeferent parameters

-----  Leasson 23 Composition and Forwarding ------
composing structures
  - group related fields together with structures
  - fill a sturcture with those structures
  - can foward a method by putting a method in the parent structure that class the inner struct method

Forwardding methods
  - Go will do method forwarding for you with struct embedding
  - to emberd a type in a structure specify the type without a field name,
    - type report struct {
        sol   int
        temp
        loc
      }
    - temperature has a average method
    - you can callit with r.average
    - is the same as r.temo.average
    - works on proterties also

name collisions
  - if theres is a collison compiler will tell you
  - implement it in the parent struct and call the the wanted method by fowarding

experiment gps.go

---- Leasson 24 Interfaces -----
  - Interfaces, con can express abstract concepts

The interface type
  - interfaces are concerned with what a type can do not the value it holds 
  - methods express the behaviou a type provides  so interfaces are declared with a set of methods that a type must satisgy
     var t interface {
      talk() string
     }
    - variable t can hold an valu of any type that sastfies the interface. Or a type if it declares a method named talk that accpts no arguments and returns a string
  - typically interfdaces are declared as named types that can be resued/ Ther's a cpnvetion of naming interface tyes with an =er suffox  a ta;ler is anything that talks
  - interface type cane be used anywhere other types are used
  - function shout(t talker){}
    - ther argument you pass to shout  function mus sastyfy the talker interface
  - go refuess to compule the program if the passed argument does not have a talk function
  - interfaces can be ysed with struct embedding 
  - starship{
      laser
    }
  -Embedding laser gives startship a talk method that forwards to the laser. now startship also satisfies tal;ker interface allowing it to be used with shout

Satisfying interfaces
  - 

*/

import (
	"encoding/json"
	"fmt"
	"math"
)

type Site struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Landing() {
	s := Site{"1", 1, 1}
	// s2 := Site{"2", 2, 2}
	//s3 := Site{"3", 3, 3}
	//sA := []Site{s,s2,s3}
	bytes, err := json.Marshal(s)

	if err != nil {
		fmt.Println("Not work")
	} else {
		fmt.Println(string(bytes))
	}

}

type World struct {
  radius float64
}

func rad(deg float64) float64{
  return deg * math.Pi /180
}
func (w *World) distance(p1, p2 Location) float64 {
  s1, c1 := math.Sincos(rad(p1.lat)) 
  s2, c2 := math.Sincos(rad(p2.lat))
  clong := math.Cos(rad(p1.lon - p2.lon))
  return w.radius * math.Acos(s1*s2 +c1+c2+clong)
}

type Location struct{
  name string
  lat float64
  lon float64
}

func (l *Location) description() string{
  return fmt.Sprintf("%v\nlat: %v\nlon: %v", l.name, l.lat, l.lon)
}

type Gps struct{
  curr Location
  dest Location
}

func (g *Gps) distance() float64 {
  s1, c1 := math.Sincos(rad(g.curr.lat)) 
  s2, c2 := math.Sincos(rad(g.dest.lat))
  clong := math.Cos(rad(g.curr.lon - g.dest.lon))
  return  math.Acos(s1*s2 +c1+c2+clong)
}
 type Rover struct {
  Gps
}

func GpsMain(){
  curiosity := Rover{
    Gps: Gps{
      curr: Location{ lat: -4.5, lon: 137.4},
      dest: Location{lat: -4.5, lon: 135.9},
    },
  }
   curiosity.distance()
}


