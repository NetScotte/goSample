package mylock

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 一个读写锁的示例， 读的场景多于写的场景
type User struct {
	name string
	count float64
	lock sync.RWMutex

	depositTotal float64
	depositLock sync.RWMutex

	drawTotal float64
	drawLock sync.RWMutex
}


func (u *User) GetName() string {
	return u.name
}


func (u *User) GetCount() float64 {
	defer u.lock.RUnlock()
	u.lock.RLock()
	return u.count
}


func (u *User) AddCount(money float64) {
	defer u.lock.Unlock()
	defer u.depositLock.Unlock()
	u.lock.Lock()
	u.depositLock.Lock()
	u.count += money
	u.depositTotal += money

}

func (u *User) SubCount(money float64) (err error){
	if u.GetCount() < money {
		err = fmt.Errorf("not enough money")
		return
	}
	defer u.lock.Unlock()
	defer u.drawLock.Unlock()
	u.lock.Lock()
	u.drawLock.Lock()
	u.count -= money
	u.drawTotal += money
	return
}


func (u *User) GetDrawTotal() float64 {
	defer u.drawLock.RUnlock()
	u.drawLock.RLock()
	return u.drawTotal
}

func (u *User) GetDepositTotal() float64 {
	defer u.depositLock.RUnlock()
	u.depositLock.RLock()
	return u.depositTotal
}


func (u *User) Monitor() (err error){
	defer u.depositLock.RUnlock()
	defer u.drawLock.RUnlock()
	defer u.lock.RUnlock()
	u.depositLock.RLock()
	u.drawLock.RLock()
	u.lock.RLock()
	if u.drawTotal + u.count != u.depositTotal {
		err = fmt.Errorf("error drawTotal + count != depositTotal")
	}
	return
}


func Deposit(u *User) {
	for {
		time.Sleep(2 * time.Second)
		money := rand.Float64() * 1000
		u.AddCount(money)
		fmt.Printf("Deposit normal: User: %v deposit money: %v, balance: %v\n", u.GetName(), money, u.GetCount())
	}
}


func Draw(u *User) {
	for {
		time.Sleep(500 * time.Millisecond)
		money := rand.Float64() * 500
		err := u.SubCount(money)
		if err != nil {
			fmt.Printf("Draw warning: User: %v want draw %v, but %v, now balance: %v\n", u.GetName(), money, err, u.GetCount())
		} else {
			fmt.Printf("Draw normal: User: %v draw %v, balance: %v\n", u.GetName(), money, u.GetCount())
		}
	}
}


func monitor(u *User, wg *sync.WaitGroup) {
	for {
		time.Sleep(10 * time.Second)
		err := u.Monitor()
		if err != nil {
			fmt.Printf("Monitor [Error]: %v\n", err)
			break
		} else {
			fmt.Printf("Monitor [Normal]\n")
		}
	}
	wg.Done()
}



