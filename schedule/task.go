package schedule

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func Info() {
	fmt.Println("I am running task.")
}

// https://github.com/jasonlvhit/gocron/issues/50
func Task() {
	gocron.Every(1).Second().Do(Info)
	gocron.Start()
}
