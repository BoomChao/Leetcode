package interview

import (
	"fmt"
	"sync"
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

// 下面这种写法不是很完美,因为最后协程会hang住，程序没有优雅退出

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

// 参考这种写法,下面这种会优雅退出

var ch = make(chan int)

var ch1 = make(chan bool)

var ch2 = make(chan bool)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			<-ch2
			num, ok := <-ch
			if !ok {
				close(ch1)
				return
			}
			fmt.Println(num)
			ch1 <- true
			time.Sleep(time.Millisecond * 300)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			<-ch1
			num, ok := <-ch
			if !ok {
				close(ch2)
				return
			}
			fmt.Println(num)
			ch2 <- true
			time.Sleep(time.Millisecond * 300)
		}
	}()

	ch2 <- true

	wg.Wait()
}

/*
 Q:当select监控多个chan同时到达就绪态时，如何先执行某个任务？
 A:可以在子case再加一个for select语句
*/

// 这样如果ch1,ch2有数据,则永远是先从ch1里面先拿出数据
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

/*
	并发协程实现求和
*/

// 下面这种方式是使用channel来进行同步的
func add(wg *sync.WaitGroup, ch chan int, receiveCh chan int) {
	defer wg.Done()
	sum := 0

	for {
		select {
		case val, ok := <-ch:
			if ok {
				sum += val
			} else {
				receiveCh <- sum
				return
			}
		}
	}
}

func main() {
	numOfTask := 10
	wg := &sync.WaitGroup{}
	ch := make(chan int, 20)
	receiveCh := make(chan int, numOfTask)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < numOfTask; i++ {
		wg.Add(1)
		go add(wg, ch, receiveCh)
	}
	wg.Wait()

	close(receiveCh)

	sum := 0
	for res := range receiveCh {
		sum += res
	}

	fmt.Println(sum)
}
