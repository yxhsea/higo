package decoretor

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	person := &Person{"John"}
	person.show()

	fmt.Println("========================")

	ts := new(TShirts)
	ts.SetDecorator(person)
	ts.show()

	fmt.Println("========================")

	bt := new(BigTrouser)
	bt.SetDecorator(person)
	bt.show()

	fmt.Println("========================")

	sk := new(Sneakers)
	sk.SetDecorator(person)
	sk.show()
}
