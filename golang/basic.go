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
func personA(web chan string, receive <-chan bool) {
	for <-receive{
	}
}
func personB(web chan string, receive <-chan bool) {
	for <-receive{
	}
}
func personC(web chan string, receive <-chan bool) {
	for <-receive{
	}
}

func center_check(from <-chan string, to chan<- string, toA,toB,toC chan<- bool) {

}

func run_time() {
	web := make(chan string)
	flagA, flagB, flagC := make(chan string)
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
}