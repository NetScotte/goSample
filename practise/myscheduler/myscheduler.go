package scheduler

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"time"
)

type EasyJob struct {
	Error error
}

func (j *EasyJob) Run() {
	defer func() {
		if e := recover(); e != nil {
			j.Error = fmt.Errorf("%v", e)
		}
	}()
	fmt.Println("run job")
	panic("手动触发了panic")
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
	entryId, err := c.AddJob("* * * * * *", &EasyJob{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(entryId)
		c.Start()
	}

	time.Sleep(5 * time.Second)
	entry := c.Entry(entryId)
	job := entry.Job.(*EasyJob)
	fmt.Println(job.Error)

	fmt.Println("wait sig to exit")
	<-sig
	c.Stop()
	fmt.Println("end")
}
