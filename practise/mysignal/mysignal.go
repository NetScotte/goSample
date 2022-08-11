package mysignal

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SampleSignal() {
	execProcessChan := make(chan error, 1)
	go func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("run task")
			case s, _ := <-execProcessChan:
				fmt.Printf("task exit with error: %s", s)
				return
			}
		}
	}()
	ctc, cancel := context.WithCancel(context.Background())
	go func() {
		signalChan := make(chan os.Signal, 1)
		defer close(signalChan)

		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

		for {
			select {
			case s, _ := <-signalChan:
				log.Warningf("receive signal: %s\n", s)
				execProcessChan <- fmt.Errorf("%s, 退出执行\n", s)
				time.Sleep(3 * time.Second)
				cancel()
			}
		}
	}()
	<-ctc.Done()
	fmt.Println("end app")
}

// run and press ctrl + c
//func main() {
//	SampleSignal()
//}
