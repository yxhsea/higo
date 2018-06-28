package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	array := []string{
		"",
		"",
		"",
		"",
	}

	for _, v := range array {
		str := fmt.Sprintf("govendor add %s", v)
		out, err := exec.Command("cmd", "/C", str).Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}
}
