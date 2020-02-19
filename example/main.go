package main

import (
	"fmt"
	"sync"

	timer "github.com/tiniyo/timer"
)

var testTimer timer.TiniyoTimer

type testTimerData struct {
	message string
}

var wg sync.WaitGroup

func testTimerCallbackHandler(data interface{}) error {
	fmt.Println("testTimerCallbackHandler : ", data.(testTimerData).message)
	wg.Done()
	return nil
}

func main() {
	wg.Add(1)
	testTimer.InitializeTiniyoTimer(testTimerCallbackHandler)
	testTimer.Run()
	testTimer.StartTimer(30, "1234", testTimerData{message: "This is test"})
	// testTimer.CancelTimer("1234")
	wg.Wait()
}
