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

func main() {

	/*
	testb := 13
	testba := testb++
	testbb := ++testb
	fmt.Println("we can find b++", testba)
	fmt.Println("we can find ++b", testbb)
	*/
	// with err, b++ can't as expressment

	stringsTest()

}