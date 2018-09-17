package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Printf("echo1: %v\n", measure("echo1"))
	fmt.Printf("echo2: %v\n", measure("echo2"))
	fmt.Printf("echo3: %v\n", measure("echo3"))
}

func measure(subject string) time.Duration {
	start := time.Now()
	exec.Command(subject, "hoge", "fuga", "piyo", "foo", "bar", "1", "2", "3", "4", "5").Run()
	fin := time.Now()
	return fin.Sub(start)
}
