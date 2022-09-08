package main

import (
	"fmt"
	"github.com/jonathanhecl/gotimeleft"
	"time"
)

func main() {

	timeleft := gotimeleft.Init(110)

	for i := 0; i < 110; i++ {
		time.Sleep(100 * time.Millisecond) // Simulate a long process

		timeleft.Step(1)
		fmt.Printf("Time left: %s - %s - %s - Speed: %.2f/s\n", timeleft.GetTimeLeft().String(), timeleft.GetProgressValues(), timeleft.GetProgress(1), timeleft.GetPerSecond())
	}

	fmt.Printf("Done! in %s\n", timeleft.GetTimeSpent().String())

}
