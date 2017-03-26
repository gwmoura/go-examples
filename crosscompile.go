package main

import "fmt"
import "runtime"

func main() {
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}

// GOOS=darwin GOARCH=386 go build crosscompile.go
/*
$ ./crosscompile
OS: darwin
Architecture: 386
*/

// GOOS=windows GOARCH=386 go build crosscompile.go
/*
$ ./crosscompile
OS: windows
Architecture: 386
*/

// GOOS=linux GOARCH=386 go build crosscompile.go
/*
$ ./crosscompile
OS: linux
Architecture: 386
*/
