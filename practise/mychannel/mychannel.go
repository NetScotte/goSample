package mychannel

import "time"

func server1(ch chan string){
	time.Sleep(6*time.Second)
	ch <- "server1"
}

func server2(ch chan string) {
	time.Sleep(2*time.Second)
	ch <- "server2"
}


