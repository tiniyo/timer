package main

import (
	"math/rand"
	"sync"

	timer "github.com/tiniyo/timer"
)

var testTimer timer.TiniyoTimer

type testTimerData struct {
	message string
}

var wg sync.WaitGroup

func testTimerCallbackHandler(data interface{}) error {
	//fmt.Println("testTimerCallbackHandler : ", data.(testTimerData).message)
	wg.Done()
	return nil
}

func main() {
	testTimer.InitializeTiniyoTimer(testTimerCallbackHandler)
	testTimer.Run()
	min := 1
	max := 5
	wg.Add(1)
	for i := 0; i < 1000000; i++ {
		testTimer.StartTimer(uint64(rand.Intn(max-min)+min), string(i), testTimerData{message: "This is test"})
		wg.Add(1)
	}
	wg.Wait()
}
