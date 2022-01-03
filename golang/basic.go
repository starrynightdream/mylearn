package main

import (
	"fmt"
	"strings"
)

func outputPoint() {
	testb := 13
	ptestb := &testb

	fmt.Println((&testb))
	fmt.Printf("pointer is type %T \n", ptestb)
	fmt.Println("show as ", ptestb)
	fmt.Println("we can find b", testb)
}

func type_switch() {
	var k interface{}
	switch k.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float")
	default:
		fmt.Println("none")
	}
}

	// use like
	// sq := seq()
	// fmt.Println(sq())
	// fmt.Println(sq())
	// fmt.Println(sq())
func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// like class, right?
type Person struct {
	name string
}

func (person Person) sayhello() string {
	ans := person.name + ": hello"
	fmt.Println(ans)
	return ans
}

	// use like
	// p := Person{name:"Jack"}
	// p.sayhello()

// end class

func stringsTest() {
	var aword = "become water, my friend"
	listofstring := []string {"stop", "download", "catch"}
	fmt.Println(aword)
	fmt.Println(strings.Join(listofstring, " "))
}

func slice_test() {
	numbers := []int {1, 1, 2, 3, 5, 8, 13}
	numbers_temp := append(numbers, 21, 34)

	show_slice(numbers)
	show_slice(numbers_temp)

	numbers_test := numbers_temp[2:5]

	show_slice(numbers_test)
}

func show_slice(sli []int) {
	fmt.Printf("size: %d cap: %d \n", len(sli), cap(sli))
	fmt.Println(sli)
}

func use_range() {
	numbers := [10]int {0,1,2,3,4,5,6,7,8,9}

	for i := range numbers {
		fmt.Println(i)
	}
}

func use_map() {
	var mymap map[string] string
	mymap = map[string] string {"chino":"coffee", "cocoa":"cocos", "molika":"moseka"}

	show_map(mymap)

	delete(mymap, "molika")

	show_map(mymap)
}
func show_map(m map[string] string) {
	fmt.Printf("info: size: %d \n", len(m))
	for k := range m {
		fmt.Println("key: ", k, " val: ", m[k])
	}
}

// interface
	// if i miss a func in one struct, i can't use it with this interface

	// use like
	// aki := Otaku{name:"**##!", love:"vetoreka"}
	// nanami := Component{belong:"last standdust"}

	// join_work(aki)
	// join_work(nanami)
type Speaker interface {
	success()
	speak(word string)
}

type Otaku struct {
	name, love string
}

type Component struct {
	belong string
}

func (otaku Otaku) success() {
	fmt.Println("my name is ", otaku.name," I love ", otaku.love)
}
func (otaku Otaku) speak(word string) {
	fmt.Println("(<ゝω·)☆: ", word)
}

func (comp Component) success() {
	fmt.Println("think to ", comp.belong)
}
func (comp Component) speak(word string) {
	fmt.Println("working and say: ", word)
}

func join_work(speaker Speaker) {
	speaker.speak("HEhe")
	speaker.success()
}
// end interface

// goroutine
// todo: finish logical
func person(web chan int, seed int, finish chan<- bool) {
	var info int 
	info_num := seed
	for true {
		info = <-web
		if (info < 0) {
			fmt.Println("player", seed, "\t Bad info! ", info)
			web <- seed * info_num
			break
		}
		info_num = info_num * 79 % 101
		web<-info_num
		fmt.Println("player", seed, "\t send: ", info_num)
	}
	finish<- true
}

func random_check (web chan int, time int) {
	var info int
	for true {
		info = <- web
		if (info % 7 ==0 || info < 0) {
			time--
			fmt.Println("checker:\t bad info! ", info)
			web <- -1
			if (time == 0) {
				break
			}
		} else {
			fmt.Println("checker:\t ok~ ", info)
			web <- info
		}
	}
}

func run_time_system() {
	var p1, p2, p3 = make(chan bool), make(chan bool), make(chan bool)
	web := make(chan int)
	go person(web, 11, p1)
	go person(web, 38, p2)
	go person(web, 93, p3)

	go random_check(web, 10000)

	web <- 5
	<-p1
	<-p2
	<-p3

	fmt.Println("over------------")
}

// end goroutine


func main() {

	/*
	testb := 13
	testba := testb++
	testbb := ++testb
	fmt.Println("we can find b++", testba)
	fmt.Println("we can find ++b", testbb)
	*/
	// with err, b++ can't as expressment
	run_time_system()
}