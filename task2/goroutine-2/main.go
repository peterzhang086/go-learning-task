package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型
type Task func()

// TaskResult 存储任务执行结果
type TaskResult struct {
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Success   bool
	Error     error
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks     []Task
	results   []TaskResult
	wg        sync.WaitGroup
	mu        sync.Mutex
	startTime time.Time
	endTime   time.Time
}

// NewScheduler 创建新的调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务到调度器
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Run 并发执行所有任务
func (s *Scheduler) Run() {
	s.startTime = time.Now()
	s.results = make([]TaskResult, len(s.tasks))

	for i, task := range s.tasks {
		s.wg.Add(1)
		go s.executeTask(i, task)
	}

	s.wg.Wait()
	s.endTime = time.Now()
}

// executeTask 执行单个任务并记录结果
func (s *Scheduler) executeTask(id int, task Task) {
	defer s.wg.Done()

	result := TaskResult{
		TaskID:    id,
		StartTime: time.Now(),
	}

	defer func() {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)

		// 捕获可能的panic
		if r := recover(); r != nil {
			result.Success = false
			result.Error = fmt.Errorf("task panicked: %v", r)
		}

		// 安全地保存结果
		s.mu.Lock()
		s.results[id] = result
		s.mu.Unlock()
	}()

	// 执行任务
	task()
	result.Success = true
}

// PrintResults 打印任务执行结果
func (s *Scheduler) PrintResults() {
	fmt.Printf("任务执行完成，总耗时: %v\n", s.endTime.Sub(s.startTime))
	fmt.Println("==========================================")

	for _, result := range s.results {
		status := "成功"
		if !result.Success {
			status = "失败"
		}

		fmt.Printf("任务 %d: %s, 耗时: %v",
			result.TaskID, status, result.Duration)

		if result.Error != nil {
			fmt.Printf(", 错误: %v", result.Error)
		}
		fmt.Println()
	}
}

// 示例任务函数
func sampleTask1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("任务1执行完成")
}

func sampleTask2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("任务2执行完成")
}

func sampleTask3() {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("任务3执行完成")
}

func failingTask() {
	time.Sleep(50 * time.Millisecond)
	panic("任务执行出错!")
}

func main() {
	// 创建调度器
	scheduler := NewScheduler()

	// 添加任务
	scheduler.AddTask(sampleTask1)
	scheduler.AddTask(sampleTask2)
	scheduler.AddTask(sampleTask3)
	scheduler.AddTask(failingTask)

	// 执行所有任务
	fmt.Println("开始执行任务...")
	scheduler.Run()

	// 打印结果
	scheduler.PrintResults()
}
