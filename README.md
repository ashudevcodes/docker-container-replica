## Container From Scratch

So basicacally this os goLang Code that halp to maka Container Like Docker Using os Module NameSpace this code will cahng the hostnme of the user but in seperate user space it will not effect in to mainroot of linux kernal and run any commad you will run i to your pc in seperate space

## the Deficulti i face is

- Sudo error
TLDR; the Non root user progam will not run the main root program
- user Application Or Root User Application
user Application are not run when you run them in to root mode why because of the user and root Application will store in to seperate bin foldel

### error i Face
```go
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
```
- this code will not run without in sudo or root user
- and also you will sure that the Go land will install in the root user
## Commad that will Run is
```bash
    sudo run go main.go run <CMD>

```
