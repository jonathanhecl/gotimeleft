package main

import (
	"fmt"
	"github.com/jonathanhecl/gotimeleft"
	"time"
)

func main() {

	timeleft := gotimeleft.Init(100)

	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond) // Simulate a long process

		timeleft.Step(1)
		fmt.Printf("Time left: %s - %s - %s\n", timeleft.GetTimeLeft().String(), timeleft.GetProgressValues(), timeleft.GetProgress())
	}

	fmt.Printf("Done! in %s\n", timeleft.GetTimeSpent().String())

}
