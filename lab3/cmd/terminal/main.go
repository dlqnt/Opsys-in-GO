package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"io/ioutil"
)

// Task 7: Simple Shell
//
// This task focuses on building a simple shell that accepts
// args that run certain OS functions or programs. For OS
// functions refer to golang's built-in OS and ioutil packages.
//
// The shell should be implemented through a command line
// application; allowing the user to execute all the functions
// specified in the task. Info such as [path] are command arguments
//
// Important: The prompt of the shell should print the current directory.
// For example, something like this:
//   /Users/meling/Dropbox/work/opsys/2020/meling-stud-labs/lab3>
//
// We suggest using a space after the > symbol.
//
// Your program should be able to at least the following functions:
// 	- exit
// 		- exit the program
// 	- cd [path]
// 		- change directory to a specified path
// 	- ls
// 		- list items and files in the current path
// 	- mkdir [path]
// 		- create a directory with the specified path
// 	- rm [path]
// 		- remove a specified file or folder
// 	- create [path]
// 		- create a file with a specified name
// 	- cat [file]
// 		- show the contents of a specified file
// 			- any file, you can use the 'hello.txt' file to check if your
// 			  implementation works
// 	- help
// 		- show a list of available args
//
// You may also implement any number of optional functions, here are some ideas:
// 	- help [command]
// 		- give additional info on a certain command
// 	- ls [path]
// 		- make ls allow for a specified path parameter
// 	- rm -r
// 		WARNING: Be aware of where you are when you try to execute this command
// 		- recursively remove a directory
// 			- meaning that if the directory contains files, remove
// 			  all the files within the directory first, then the
// 			  directory itself
// 	- calc [expression]
// 		- Simple calculator program that can calculate a given expression
// 			- example expressions could be + - * \ pow
// 	- ipconfig
// 		- show ip interfaces
// 	- history
// 		- show command history
// 		- Alternatively implement this together with pressing up on your
// 		  keypad to load the previous command
// 		- clrhistory to clear history
// 	- tail [n]
// 		- show last n lines of a file
// 	- head [n]
// 		- show first n lines of a file
// 	- writefile [text]
// 		- write specified text to a specified file
//
// 	Or, alternatively, implement your own functionality not specified as you please
//
// Additional notes:
// 	- If you want to use colors in your terminal program you can see the package
// 		"github.com/fatih/color"
//
// 	- Helper functions may lead to cleaner code
//

// Terminal contains
type Terminal struct {
}

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")

// Execute executes a given command
func Execute(command string) error {

	args := strings.Split(command, " ")

	switch args[0] {
	case "cd":

		err := os.Chdir(args[1])
		if err != nil {
			panic(err)
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	case "ls":
		files, err := ioutil.ReadDir("./")
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	case "rm":
		os.Remove(args[1])
	case "cat":
		data, err := os.ReadFile(args[1])
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(data)
	case "create":
		_, err := os.Create(args[1])
		if err != nil {
			panic(err)
		}

	case "mkdir":
		err := os.Mkdir(args[1], 0777)
		if err != nil {
			panic(err)
		}
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command return the error.
	return cmd.Run()
}

// Split the command separate the command and the arguments.

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		path, err := os.Getwd()
		if err != nil {
			print("Error main")
			panic(err)
		}
		fmt.Print(path + "> ")
		// Read the keyboad command.
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Remove the newline character.
		command = strings.TrimSuffix(command, "\n")

		// Skip an empty command.
		if command == "" {
			continue
		}

		// Handle the execution of the command.
		if err = Execute(command); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
