package main

import (
	"fmt"
	"time"

	"github.com/jonathanhecl/gotimeleft"
)

func main() {

	timeleft := gotimeleft.Init(110)

	for i := 0; i < 110; i++ {
		time.Sleep(100 * time.Microsecond) // Simulate a long process

		timeleft.Step(1)
		fmt.Printf("%s Time left: %s - %s - %s - Speed: %.2f/s\n", timeleft.GetProgressBar(30), timeleft.GetTimeLeft().String(), timeleft.GetProgressValues(), timeleft.GetProgress(1), timeleft.GetPerSecond())
	}

	fmt.Printf("Done! in %s\n", timeleft.GetTimeSpent().String())

}
