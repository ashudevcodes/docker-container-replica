## Container From Scratch

This Go code helps create a container similar to Docker using the `os` module and namespaces. This code will change the hostname of the user but in a separate user space. It will not affect the main root of the Linux kernel and will run any command you execute on your PC in a separate space.

## The Difficulties I Face

- **Sudo Error**: The program will not run as a non-root user; it requires root privileges.
- **User Application or Root User Application**: User applications do not run when executed in root mode because user and root applications are stored in separate bin folders.

### Error I Face
```go
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
```
- This code will not run without using `sudo` or being the root user.
- Ensure that Go is installed for the root user.

## Command to Run This Code
```bash
    sudo go run main.go run /bin/bash
```


## Potential Issues that were fixed

- **Argument Checking**: The program assumes that os.Args[1] always exists, which can cause a panic if no arguments are provided.
  ```if len(os.Args) < 2 {
    fmt.Printf("Usage: %s [run|child] [command]\n", os.Args[0])
    os.Exit(1)
}
```
