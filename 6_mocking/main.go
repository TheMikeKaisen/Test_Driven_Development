package main

import (
	"fmt"
	"io"
	"os"
	"time"
)


func Countdown(writer io.Writer){
		finalOut := "Go!"
		initialOut := 3
		for i:=initialOut; i>0; i--{
			fmt.Fprint(writer, i);
			time.Sleep(1 * time.Second)
		}
		fmt.Fprint(writer, finalOut)
}

func main() {
	Countdown(os.Stdout);
}