package interview

import (
	"fmt"
	"sync"
	"time"
)

/*
	协程交叉打印123
*/

// 下面这种写法缺点就是没有优雅退出

var first = make(chan bool)
var second = make(chan bool)
var third = make(chan bool)

func main() {
	go func() {
		for {
			<-first
			fmt.Println(1)
			time.Sleep(time.Second)
			second <- true
		}
	}()
	go func() {
		for {
			<-second
			fmt.Println(2)
			time.Sleep(time.Second)
			third <- true
		}
	}()
	go func() {
		for {
			<-third
			fmt.Println(3)
			time.Sleep(time.Second)
			first <- true
		}
	}()

	first <- true
	time.Sleep(time.Second * 10)
}

// 优雅退出的写法

var first = make(chan bool)
var second = make(chan bool)
var third = make(chan bool)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-first
			fmt.Println(1)
			second <- true
		}
		close(second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-second
			fmt.Println(2)
			third <- true
		}
		close(third)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-third
			fmt.Println(3)
			first <- true
		}
		close(first)
	}()

	first <- true

	time.Sleep(time.Second)

	<-first

	wg.Wait()
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

var chNum = make(chan bool)

var chChar = make(chan bool)

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			_, ok := <-chNum
			if !ok {
				break
			}
			fmt.Println(i)
			chChar <- true
		}
		close(chChar)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			_, ok := <-chChar
			if !ok {
				break
			}
			fmt.Printf("%c\n", i+'A')
			chNum <- true
		}
		close(chNum)
	}()

	chNum <- true

	// 这里一定要从channel里面读出数据,因为这是无缓冲的channel,否则上面的 chNum<-true写入会阻塞,因为没有下游来读
	<-chNum

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

func main() {
	wg := &sync.WaitGroup{}

	numOfTasks := 20

	ch := make(chan int)
	sum := make(chan int, numOfTasks)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < numOfTasks; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			subSum := 0
			for num := range ch {
				subSum += num
			}
			sum <- subSum
		}()
	}
	wg.Wait()

	// 这里需要先关闭
	close(sum)

	res := 0
	for n := range sum {
		res += n
	}

	fmt.Println(res)
}
