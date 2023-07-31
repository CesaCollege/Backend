package main

import (
	"fmt"
	"sync"
	"time"
)

type MainPage struct {
	WG              *sync.WaitGroup
	statusMessageCH chan string
}

func (mp *MainPage) Authenticate(userid int) {
	// auth logic
	time.Sleep(50 * time.Millisecond)
	fmt.Println("user is authenticated")
}
func (mp *MainPage) UserInfo(userid int) {
	// auth logic
	time.Sleep(100 * time.Millisecond)
	mp.statusMessageCH <- fmt.Sprint("user info is ready")
	mp.WG.Done()
}
func (mp *MainPage) UserHistory(userid int) {
	// auth logic
	time.Sleep(100 * time.Millisecond)
	mp.statusMessageCH <- fmt.Sprint("user history is ready")
	mp.WG.Done()
}
func (mp *MainPage) Recommendations(userid int) {
	// auth logic
	time.Sleep(2 * time.Second)
	mp.statusMessageCH <- fmt.Sprint("recoms are ready")
	mp.WG.Done()
}
func main() {
	start_time := time.Now()
	userID := 1
	wg := &sync.WaitGroup{}
	mainPage := &MainPage{
		WG:              wg,
		statusMessageCH: make(chan string, 3),
	}
	mainPage.Authenticate(userID)
	// flag := 3
	wg.Add(3)
	go mainPage.UserInfo(userID)
	go mainPage.UserHistory(userID)
	go mainPage.Recommendations(userID)
	wg.Wait()
	close(mainPage.statusMessageCH)
	for msg := range mainPage.statusMessageCH {
		fmt.Println(msg)
	}
	// time.Sleep(2*time.Second)
	// for flag != 0 {

	// }
	fmt.Println("page completely loaded")

	duraion := time.Since(start_time)

	fmt.Println(duraion.Milliseconds())

}
