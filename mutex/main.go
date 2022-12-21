package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var BALANCE float32 = 12000

type User struct {
	name           string
	depositAmount  float32
	withdrawAmount float32
}

func (u *User) deposit(wg *sync.WaitGroup, mu *sync.Mutex) {
	fmt.Printf("%s is deposting USD %f\n", u.name, u.depositAmount)
	mu.Lock()
	BALANCE += u.depositAmount
	defer mu.Unlock()
	wg.Done()
}

func (u *User) withdraw(wg *sync.WaitGroup, mu *sync.Mutex) {
	fmt.Printf("%s is withdrawing USD %f\n", u.name, u.withdrawAmount)
	mu.Lock()
	BALANCE -= u.withdrawAmount
	defer mu.Unlock()
	wg.Done()
}

func main() {
	// raceCondition()
	// withRWMutex()
	embeddedInStruct()
}

func raceCondition() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	users := []User{
		{name: "Marco Lazerri", withdrawAmount: 1300, depositAmount: 1000},
		{name: "Paige Wilunda", withdrawAmount: 1400, depositAmount: 123},
		{name: "Gerry Rivers", withdrawAmount: 900, depositAmount: 25},
		{name: "Sean Bold", withdrawAmount: 200, depositAmount: 5432},
		{name: "Mike Wegner", withdrawAmount: 5600, depositAmount: 2344},
	}
	rand.Seed(time.Now().UnixNano())
	for i := range users {
		wg.Add(2)

		i = rand.Intn(len(users))
		go users[i].deposit(&wg, &mu)
		go users[i].withdraw(&wg, &mu)
		time.Sleep(time.Second)
	}

	wg.Wait()
	fmt.Printf("New account BALANCE is %f\n", BALANCE)
}

var (
	rwLock sync.RWMutex
)

func read(wg *sync.WaitGroup) {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Reading locking...")
	time.Sleep(time.Second)
	fmt.Println("Reading unlocking...")
	wg.Done()
}

func write(wg *sync.WaitGroup) {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write locking...")
	time.Sleep(time.Second)
	fmt.Println("Write unlocking...")
	wg.Done()
}

func readerWriter(wg *sync.WaitGroup) {
	wg.Add(5)

	go write(wg)
	go read(wg)
	go read(wg)
	go read(wg)
	go write(wg)

	time.Sleep(time.Second)
	fmt.Println("Done...")
	// wg.Done()
}

func withRWMutex() {
	var wg sync.WaitGroup
	readerWriter(&wg)
	wg.Wait()
	fmt.Println("Main thread is done")
}

type User2 struct {
	depositAmount float32
	sync.Mutex
}

var BALANCE2 float32 = 12000

func (u *User2) deposit(wg *sync.WaitGroup) {
	u.Lock()
	BALANCE2 += u.depositAmount
	u.Unlock()
	wg.Done()
}

func embeddedInStruct() {
	var wg sync.WaitGroup
	user := User2{depositAmount: 1200}
	for i := 0; i < 10; i++ {
		wg.Add(3)

		go user.deposit(&wg)
		go user.deposit(&wg)
		go user.deposit(&wg)
	}
	wg.Wait()
	fmt.Printf("New account BALANCE is %f\n", BALANCE2)
}
