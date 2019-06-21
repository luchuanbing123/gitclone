package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	isForGo := len(os.Args) > 2 && strings.ToLower(os.Args[2]) == "go"
	url := os.Args[1]
	arr := strings.Split(url, "/")
	if isForGo {
		arr[1] = os.Getenv("GOPATH") + "/src"
	} else {
		if runtime.GOOS == "windows" {
			arr[1] = "c:\\code"
		} else {
			arr[1] = "~/code"
		}
	}

	dirPath := strings.Join(arr[1:len(arr)-1], "/")
	_ = os.MkdirAll(dirPath, os.ModeDir)

	repName := strings.Split(arr[len(arr)-1], ".")[0]
	repPath := dirPath + "/" + repName
	_ = os.RemoveAll(repPath)
	cmd := exec.Command("git", "clone", url, repPath)
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	fmt.Printf("already cloned %s to %s", url, repPath)
}
