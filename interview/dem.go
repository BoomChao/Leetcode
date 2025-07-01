package interview

import (
	"fmt"
	"time"
)

/*
	有三个函数，分别打印"cat", "fish","dog"要求每一个函数都用一个goroutine，按照顺序打印100次。
*/

var cat = make(chan struct{})
var dog = make(chan struct{})
var fish = make(chan struct{})

func Dog() {
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
}

func Fish() {
	<-cat
	fmt.Println("fish")
	fish <- struct{}{}
}

func Cat() {
	<-dog
	fmt.Println("cat")
	cat <- struct{}{}
}

func main() {
	for i := 0; i < 100; i++ {
		go Dog()
		go Fish()
		go Cat()
	}
	fish <- struct{}{}
	time.Sleep(time.Second * 10)
}

/*
	两个协程交替打印10个字母和数字
*/

var word = make(chan struct{}, 1)
var num = make(chan struct{}, 1)

func printNums() {
	for i := 0; i < 10; i++ {
		<-word
		fmt.Println(i)
		num <- struct{}{}
	}
}

func printWords() {
	for i := 0; i < 10; i++ {
		<-num
		fmt.Printf("%c\n", 'a'+i)
		word <- struct{}{}
	}
}

func main() {
	num <- struct{}{}
	go printNums()
	go printWords()
	time.Sleep(time.Second * 2)
}

/*
 Q:当select监控多个chan同时到达就绪态时，如何先执行某个任务？
 A:可以在子case再加一个for select语句
*/

// 这样如果ch1,ch2有数据,则永远是先从ch1里面拿出数据
func priority(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}
