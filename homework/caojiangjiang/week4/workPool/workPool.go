package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 创建一个工作池
type Task struct {
	ID      int
	randnum int
}

type Result struct {
	task    Task
	result  int
}

var tasks = make(chan Task, 10)
var results = make(chan Result, 10)
var mutex sync.Mutex

func process(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	time.Sleep(2 * time.Second)
	// 加锁
	mutex.Lock()
	fd_time := time.Now().Format("2006-01-02 15:04:05");
	sumStr := strconv.Itoa(sum)
	fd_content:=strings.Join([]string{"======",fd_time,"=====",sumStr,"\n"},"")
	log.Println("test-caojiangjiang-" + fd_content)
	mutex.Unlock()
	return sum
}

func createWorkerPool(numOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	// 这里是否需要关闭results通道，是由稍后的range迭代这个通道决定的，不关闭这个通道会一直阻塞range，最终导致死锁
	close(results)
}

//
// 最后需要关闭tasks通道，因为所有任务都分配完之后，没有任务再需要分配。
// 当然，这里之所以需要关闭tasks通道，是因为worker()中使用了range迭代tasks通道
// 如果不关闭这个通道，worker将在取完所有任务后一直阻塞，最终导致死锁。
func allocate(numOfTasks int) {
	for i := 0; i < numOfTasks; i++ {
		rand.Seed(time.Now().UnixNano()) // 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
		randnum := rand.Intn(999)
		task := Task{i, randnum}
		tasks <- task
	}
	close(tasks)
}

// 创建工作进程
func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		result := Result{task, process(task.randnum)}
		results <- result
	}
}

func getResult(done chan bool)  {
	for result := range results {
		fmt.Printf("Task id %d, randnum %d , sum %d\n", result.task.ID, result.task.randnum, result.result)
	}
	done <- true
}

func main()  {
	// 记录起始终止时间，用来测试完成所有任务耗费时长
	startTime := time.Now()

	numOfWorkers := 20
	numOfTasks := 100
	// 创建任务到任务队列中
	go allocate(numOfTasks)
	// 创建工作池
	go createWorkerPool(numOfWorkers)
	// 取得结果
	var done = make(chan bool)
	go getResult(done)
	<- done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
