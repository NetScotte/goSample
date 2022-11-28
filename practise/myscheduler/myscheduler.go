package scheduler

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"time"
)

// Job 通用简单的Job， 其他job应该
type Job struct {
	Error  error
	Params string // 任何参数，通过json序列化和反序列化
	Fn     func() error
}

func NewEasyJob(params string, fn func() error) *Job {
	return &Job{
		Params: params,
		Fn:     fn,
	}
}

func (j *Job) Run() {
	defer func() {
		if e := recover(); e != nil {
			j.Error = fmt.Errorf("%v", e)
		}
	}()
	if j.Fn != nil {
		err := j.Fn()
		if err != nil {
			j.Error = err
		}
	} else {
		j.Error = fmt.Errorf("无法执行job， fn未指定！！！")
	}
}

func Basic() {
	fmt.Println("start")
	sig := make(chan os.Signal, 1)
	c := cron.New(cron.WithSeconds())
	entryId, err := c.AddFunc("* * * * * *", func() { fmt.Println(time.Now().Format("2006-01-02 15:04:05")) })
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(entryId)
		c.Start()
		<-sig
	}
	// 等待一段时间后，移除任务
	time.Sleep(10 * time.Second)
	fmt.Printf("cancel task: %v\n", entryId)
	c.Remove(entryId)

	fmt.Println("wait sig to exit")
	<-sig
	c.Stop()
	fmt.Println("end")
}

func JobSample() {
	fmt.Println("start")
	sig := make(chan os.Signal, 1)
	c := cron.New(cron.WithSeconds())
	entryId, err := c.AddJob("* * * * * *", &Job{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(entryId)
		c.Start()
	}

	time.Sleep(5 * time.Second)
	entry := c.Entry(entryId)
	job := entry.Job.(*Job)
	fmt.Println(job.Error)

	fmt.Println("wait sig to exit")
	<-sig
	c.Stop()
	fmt.Println("end")
}
