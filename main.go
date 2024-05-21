package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [run|child] [command]\n", os.Args[0])
		os.Exit(1)
	}
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		fmt.Printf("bad command\n")
	}
}

func run() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
}

func child() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	fmt.Printf("Host Name Change...\n")
	if err := syscall.Sethostname([]byte("AshishNemo")); err != nil {
		fmt.Printf("Error setting hostname: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Root Change...\n")
	if err := syscall.Chroot("/path/to/new/root"); err != nil { // Change this to a valid path
		fmt.Printf("Error changing root: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Directory Change...\n")
	if err := syscall.Chdir("/home"); err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		fmt.Println("Error :) ", err)
		os.Exit(1)
	}
}
