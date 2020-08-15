package main

import (
	"flag"
	"fmt"
	"greeting"
	"time"
)

func main() {
	name := flag.String("name", "", "please enter your name")
	flag.Parse()

	if *name == "" {
		fmt.Println("please enter your name", *name)
		return
	}

	greetee := greeting.Person{
		Name: *name,
	}
	fmt.Printf("hello %s\n", greetee.Name)
	fmt.Printf("the time is %s\n", time.Now().Format("2006-01-02 15:04:05"))
}
