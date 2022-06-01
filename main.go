package main

import (
	"fmt"
	"os"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		check(err)
	}
	anydeskId, err := exec.Command("anydesk", "--get-id").Output()
	if err != nil {
		check(err)
	}
	createObject, err := os.Create("/home/user/Documents/" + hostName + ".txt")
	if err != nil {
		check(err)
	}
	defer createObject.Close()

	createObject.WriteString(string(anydeskId))
	fmt.Println(anydeskId, createObject)

}
