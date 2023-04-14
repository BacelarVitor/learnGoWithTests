package main

import (
	"os"
	"time"

	"github.com/BacelarVitor/learnGoWithTests/mocking"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
