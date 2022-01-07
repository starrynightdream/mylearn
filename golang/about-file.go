package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writing_file(file_name string) {
	p:= fmt.Printf

	data1 := []byte("abcdefg\n")
	err := ioutil.WriteFile(file_name, data1, 0644)
	check(err)

	// os.Open reanonly
	// f, err := os.Open(file_name)
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer f.Close()

	data2 := []byte("hijklmm\n")
	data3 := []byte{111, 112, 113, 114, 115, 116}
	n2, err := f.Write(data2)
	check(err)
	p("write data2 %d byte\n", n2)
	n3, err := f.Write(data3)
	check(err)
	p("write data3 %d byte\n", n3)

	n4, err := f.WriteString("\nuvw\n")
	check(err)
	p("write string %d byte\n", n4)

	f.Sync()

	w := bufio.NewWriter(f)
	defer w.Flush()
	n5, err := w.WriteString("xyz\n")
	check(err)
	p("write string %d byte\n", n5)
}

func reading_file(file_name string) {
	p := fmt.Printf

	dat, err := ioutil.ReadFile(file_name)
	check(err)
	p(string(dat))

	// readonly
	f,err := os.Open(file_name)
	check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	p("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(3, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	p("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(3, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	p("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 倒带
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	p("5 bytes: %s\n", string(b4))
}

func main() {
	file_name := "playground.dat"
	writing_file(file_name)
	reading_file(file_name)
}