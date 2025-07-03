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

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t task) do() {
	sum := 0
	for i := t.begin; i < t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func buildTask(taskChan chan<- task, resultChan chan<- int, count int) {
	group := count / 10
	mod := count % 10
	if mod != 0 {
		group += 1
	}
	for i := 0; i < group; i++ {
		end := (i + 1) * 10
		if end > count {
			end = count
		}
		tsk := task{
			begin:  i * 10,
			end:    end,
			result: resultChan,
		}
		taskChan <- tsk
	}
	close(taskChan)
}

func distributeTask(taskChan <-chan task, workers int, done chan<- struct{}) {
	for i := 0; i < workers; i++ {
		go workerTask(taskChan, done)
	}
}

func workerTask(taskChan <-chan task, done chan<- struct{}) {
	for v := range taskChan {
		v.do()
		done <- struct{}{}
	}
}

func closeResult(done chan struct{}, resultChan chan<- int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultChan)
}

func processResult(resultChan <-chan int) int {
	sum := 0
	for v := range resultChan {
		sum += v
	}
	return sum
}

func main() {
	workers := 5
	count := 100

	taskChan := make(chan task, 10)
	resultChan := make(chan int, 10)
	done := make(chan struct{}, 10)

	go buildTask(taskChan, resultChan, count)
	go distributeTask(taskChan, workers, done)
	closeResult(done, resultChan, workers)

	sum := processResult(resultChan)
	fmt.Println(sum)
}
