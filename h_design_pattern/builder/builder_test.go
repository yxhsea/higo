package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	thin := Thin{}
	fat := Fat{}

	director := Director{&thin}
	director.CreatePerson()

	fmt.Println("================================")

	director = Director{&fat}
	director.CreatePerson()
}
