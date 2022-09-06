package unit7

/*
------------------------ Concurrent Programing --------------------------------
  - any go code can be made concurent by startin it in a goroutine
  - goroutines use channels for communication and coordination

----- Lesson 30 Foroutines and concurrency--------
  - gorutines are simillar to threads in other languges

Starting a goroutine
  - all you need is the go keyword in front of the call
  - sleepygopher.go

More than one goruine
  - each time we use go keyword a new go ruines is tarted
  - all gorutines appear to run at the same time.
    - thecnically they mufht not because computer only hace a limited number of processing units
  - processors spend some time on one fo routine before proceesing to nother. thecnique n=known as time sharing
  -if the goroutines are foung more than just sleeping we weont know how long the're going to take to do their work
    -we need some way for the code to know when all the goruitnes have finished. fourunatly go provides us with exactly what we need channels

Channels
  - can be used to send values safely from one forutine to another
  - channels can be usaed as variables, passed to function stored in structure and do alomost anything else you want them to do
  - channles are a type
  - to create a channel ue make
  - channles have a type when you make them
    - c:= make(chan int)
      - can send and recieve int
  - you can send and recieve values
  - you send or recieve on a channel with the left arrow operator <-
  - to send a value point the arrow toward the channels expression as
  - following sends the value 99
    - c <- 99
  - to receive a value from a channel the arrow points away from the channel
    - r:= <-c
      - in the following code we recieve a value from channel c and assign it to variable r. simarly to sending on a chanel the reciever will wait
        until another gorutine tries to send on the same channel

Channel surfin with select
  - we want to continue recivering values from the go rutines unitl they finnish or we are done waiting that means we need to wait on both timer and the other channel at the same tieme
  - the slect satement allows sus to do this
  - select statement looks like a switch
  - each case insde a selectholds a channel recive or send
  - slect wait until one case is ready and then runs it and its associated statement

Nil channels do nothing
  - channels default value is nil
  - if tou try to use a nil channel it won't panic
    - opetation send or recieve will blcok foreever

Bloacking and deadlock
  - when waoting to send or recive we sa it vlocked
  - a blcoked rrotuine takes no resources other than memory
  - when gortuines are bloked for something that can never happen it;s called dealock

a gopher assembly line
  - A PIPELINE SEE testPipeline
  - we could close the channel
    - close(c)
  - when a channel is closed tou can't write any more values to it you'll get a panic if tou try
  - and any read will return immediatly with the zero value for the type not of the channel of the type it sends and reives
  - reading from a closed channel the loop will spin forever
  -check if channel closed
    - v, ok := <c
      - ok will be true id sucessfully read form the channel
      - it's false when the channel has been closed
  - with these we can easily closed the pipeline
  - if we use channel in a reange statement it will read values from the channel until the channel is closed

Experiment remove-identival.go
Experiment SplitWords.go

----- Lesson 31 Concurrent State ------
  - shared values
    - unless we explictly know that it's okay to use the specific kind of value in question concurrently we must assume that is's not okay
      this kinf of sitation is known as race condition because it's as if the gorutins are racing to use the value
  - mutex
    - mutual exclusion
    - fo ruitnes can use mutex to exculde each other from doing something at the same time/ Th someting is queestion is up to the programmer to decde
    - two methods
      - lock
        - is like taking a token from a jar
        - we put back the token by calling unlck
      - unlock
    - if any fgorutine call lock while the mutex is locked it'll wait unti it's unlocked brefore locking again
    - unlikke channles fo mutexes aren't niolt into the language rahter that are available in sync package
    - scrape.go

Mutex pitfalls
  - may blok a gorutine
  - may cause a deadlok

LongLive workers
  -
*/

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func TestRover() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
  r.Stop()

	time.Sleep(3 * time.Second)
  r.Start()

	time.Sleep(3 * time.Second)
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

type RoverDriver struct {
	commandc chan command
}

func (r RoverDriver) Left() {
	r.commandc <- left
}

func (r RoverDriver) Right() {
	r.commandc <- right
}

func (r RoverDriver) Stop() {
	r.commandc <- stop
}

func (r RoverDriver) Start() {
	r.commandc <- start
}

func (r RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	prevDirection := direction
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case stop:
				prevDirection = direction
				direction = image.Point{X: 0, Y: 0}
      case start:
        direction = prevDirection
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func TestVisited() {
	v := Visited{visited: map[string]int{}}
	go v.VisitLink("a.com")
	go v.VisitLink("a.com")
	go v.VisitLink("a.com")
	go v.VisitLink("a.com")
	go v.VisitLink("b.com")
	time.Sleep(3 * time.Second)
	fmt.Println(v.visited)
}

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func SplitWords() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceSplitWords(c1)
	go splitWords(c1, c2)
	printSplitWords(c2)
}

func sourceSplitWords(downstream chan string) {
	for _, s := range []string{"a b", "c d"} {
		downstream <- s
	}
	close(downstream)
}

func splitWords(downstream, upstream chan string) {
	for s := range downstream {
		split := strings.Split(s, " ")
		for _, word := range split {
			upstream <- word
		}
	}
	close(upstream)
}

func printSplitWords(upstream chan string) {
	for s := range upstream {
		fmt.Println(s)
	}
}

func RemoveIdentical() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceDuplicate(c1)
	go filterDuplicates(c1, c2)
	printRemovedDuplicates(c2)
}

func sourceDuplicate(downstream chan string) {
	for _, s := range []string{"a", "b", "b"} {
		downstream <- s
	}
	close(downstream)
}

func filterDuplicates(downstream chan string, upstream chan string) {
	a := ""
	for s := range downstream {
		if a != s {

			upstream <- s
		}
		a = s
	}
	close(upstream)
}

func printRemovedDuplicates(upstream chan string) {
	for s := range upstream {
		fmt.Println(s)
	}
}

func TestSleepyGopherChanels() {
	c := make(chan int) // Makes a channel to communicate ovaer
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		gopherId := <-c // recieves value from channel

		fmt.Println("gopher ", gopherId, " has finnished") //
	}
}

func TestSleepyGopherSelect() {
	c := make(chan int) // Makes a channel to communicate ovaer
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	timeout := time.After(2 * time.Second) //make a timeout channel
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c:
			fmt.Println("gopher ", gopherId, " has finnished")
		case <-timeout:
			fmt.Println("no patirnece")
			return
		}

	}
}

func sleepyGopher(i int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond) // gopher sleeps c <- i
	c <- i
}

func TestPipeline() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)

}
func sourceGopher(downstream chan string) {
	s := []string{"hello world", "a bad apple", "goodbye all"}
	for _, v := range s {
		downstream <- v
	}

	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
