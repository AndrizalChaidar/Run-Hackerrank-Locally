package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Should add runtime type argument ("rust" | "go")!!`)
		return
	}
	defer timer("exec")()
	data, err := os.ReadFile("./hackerrank_test.txt")
	handleErr(err)
	var process *exec.Cmd
	runtimeType := os.Args[1]
	if runtimeType == "rust" {
		process = exec.Command("cargo", "run", "--manifest-path", "./hackerrank_exercise/rust/Cargo.toml", "--release")
	} else if runtimeType == "go" {
		process = exec.Command("go", "run", "./hackerrank_exercise/golang/tmpl.go")
	} else {
		fmt.Printf("Runtime type %s is not found\n", runtimeType)
		return
	}
	piped_stdin, err := process.StdinPipe()
	handleErr(err)
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	err = process.Start()
	handleErr(err)

	_, err = piped_stdin.Write(data)
	handleErr(err)
	err = piped_stdin.Close()
	handleErr(err)
	process.Wait()
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
