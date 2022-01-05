package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
	"sync"
	"sync/atomic"
)

func speech_dealy(word string, times int, sign chan<- string) {
	time.Sleep(time.Duration(times) * time.Second)
	sign<- word
}

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

// goroutine buffer
func test_buffer_A() {
	info := make(chan string, 2)

	info <- "hello"
	info <- "world"

	fmt.Println(<-info)
	fmt.Println(<-info)
}

// will err, it can't save two string
func test_buffer_B() {
	info := make(chan string)

	info <- "hello"
	info <- "world"

	fmt.Println(<-info)
	fmt.Println(<-info)
}

func test_buffer() {
	test_buffer_A()
	fmt.Println("------------------")
	test_buffer_B() // err
}
// end goroutine buffer

// select
func use_select() {
	// waiting all sign return at same time
	info1 := make(chan string)
	info2 := make(chan string)
	go speech_dealy("will one second", 1, info1)
	go speech_dealy("will two second", 2, info2)

	// will wating both of them
	// one of them come it will run and to next
	for i:=0; i<2;i++{
	// there change i < 2, it will check one second at first, two second after
		select{
		case msg:= <-info1:
			fmt.Println("info1: ", msg)
		case msg:= <-info2:
			fmt.Println("info2: ", msg)
		case <-time.After(time.Second * 3):
			// out of time 3s, will not run
			// a way for timer
			fmt.Println("out of time")
		}
	}
}

// select all time
func select_player(name string, target_ip chan<- int, info <-chan int) {
	keep_talking := true
	for keep_talking{
		select{
		case i, more := <-info: // more false mean channel close
			if (!more) {
				fmt.Println(name, "\t quit chattingroom")
				keep_talking = false
			} else{
				fmt.Println(name, "\t get ", i)
				target_ip <- (i * 10086 + 23) % 37
			}
		default:
			fmt.Println(name, "\t watting...")
			time.Sleep(time.Duration(500) * time.Millisecond)
		}
	}
}

func chatting_room(a_ip, b_ip <-chan int, a_info, b_info chan<- int, work_done chan<- bool, seed int) {
	a_info <- seed // start
	keep_chatting := true
	for keep_chatting {
		select{
		case a_send:= <-b_ip:
			if (a_send % 7 == 0) {
				close(a_info)
				close(b_info)
				fmt.Println("ending...")
				keep_chatting = false
				work_done<-true
			} else {
				b_info <- a_send
			}
		case b_send:= <-a_ip:
			if (b_send % 7 == 0) {
				close(a_info)
				close(b_info)
				fmt.Println("ending...")
				keep_chatting = false
				work_done<-true
			} else {
				a_info <- b_send
			}
		default:
			fmt.Println("watting charring")
			time.Sleep(time.Second * 1)
		}
	}
}

func test_async_chatting() {
	a_ip := make(chan int)
	b_ip := make(chan int)
	a_info := make(chan int)
	b_info := make(chan int)

	wd := make(chan bool)

	go select_player("person", b_ip, a_info)
	go select_player("genshi", a_ip, b_info)

	go chatting_room(a_ip, b_ip, a_info, b_info, wd, 2)

	<-wd
}

// chatting ending

// range test
func boss_send_work(all_work int, jobs chan<- int) {
	for i := all_work; i>0; i-- {
		jobs<-i
		time.Sleep(time.Millisecond * time.Duration(100))
	}
	close(jobs)
}

func woker_due_work(jobs <-chan int, work_done chan<- bool, woker_id int) {
	// use close() close the channel make range stop
	for work := range jobs {
		fmt.Printf("woker%d\t fix work %d\n", woker_id, work)
	}
	work_done<-true
}

func test_working_range() {
	jobs := make(chan int, 20)
	wd1 := make(chan bool)
	wd2 := make(chan bool)
	go boss_send_work(20, jobs)
	go woker_due_work(jobs, wd1, 1)
	go woker_due_work(jobs, wd2, 2)

	<-wd1
	<-wd2
}
// range test for queue end

// timer
func timer_test() {
	timer1 := time.NewTimer(time.Second * 2)
	<- timer1.C
	fmt.Println("timer1 end")

	// timer different to time.sleep is you can cancel
	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("timer2 end")
	}()

	stop2 := timer2.Stop()
	if (stop2) { // success stop
		fmt.Println("timer2 stop")
	}
}

// ticker
func ticker_test() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func () {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(time.Second * 4)
	ticker.Stop()
	fmt.Println("ticker stop")
}

// worker pool
func worker_pool_singel_worker(id int, work <-chan int, result chan<- int) {
	for w := range work {
		fmt.Printf("woker%d start \t work %d\n", id, w)
		time.Sleep(time.Second)
		fmt.Printf("woker%d ending\t work %d\n", id, w)
		result <- w*2
	}
}
func worker_pool_center() {
	works := make(chan int, 100)
	results := make(chan int, 100)

	var wokers_number, works_number = 3, 7

	for i:=0; i< wokers_number; i++ {
		go worker_pool_singel_worker(i, works, results)
	}

	for i:=0; i < works_number; i++ {
		works <- i
	}
	close(works)
	for i:=0; i < works_number; i++ {
		<-results
	}
}
// end worker pool

// rate limiting
func normal_rate_limiting() {
	requests := make(chan int, 5)
	for i:=0;i<5;i++ {
		requests <- i
	}
	close(requests)

	ticker := time.Tick( time.Millisecond * 500)
	for req := range requests {
		<-ticker
		fmt.Println("request ", req," handle at ", time.Now())
	}
}

func short_bursts_rate_limiting() {
	wd := make(chan bool)
	wwd := make(chan bool)
	burstyLimiter := make(chan time.Time, 3)
	// will find work 0-2 is run nearly same time
	for i:=0; i<3; i++ {
		burstyLimiter<-time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 500) {
			select {
			case <-wd:
				wwd <- true
				return
			default:
				burstyLimiter <- t
			}
		}
	}()

	burstyRequests := make(chan int ,10)
	for i:=0;i<10;i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request ", req," handle at ", time.Now())
	}
	wd<-true
	<-wwd
}
// end rate limiting

// add num safe
func automic_counter() {
	var t uint64 = 0
	for i:=0; i<50; i++{
		// auto kill when main thread end
		go func (){
			for {
				atomic.AddUint64(&t, 1)
				time.Sleep(time.Millisecond * 1)
			}
		}()
	}

	time.Sleep(time.Second * 2)
	// the t still update, so we get the copy
	// use func LoadUint64
	tFinal := atomic.LoadUint64(&t)
	fmt.Println("final: ", tFinal)
}

// state write read
func state_write_read() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps, writeOps uint64 = 0, 0
	
	for i:=0; i<100; i++ {
		go func() {
			// read, close at main end
			total := 0 // meanless
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i:=0; i<5; i++ {
		go func() {
			// write, close at main end
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second * 2)
	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Printf("total read: %d, total write: %d \n", readOpsFinal, writeOpsFinal)
	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()
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

	state_write_read()
}