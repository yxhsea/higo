package proxy

import "testing"

func TestNewProxy(t *testing.T) {
	girl := Girl{}
	girl.SetName("John")

	p := NewProxy(girl)
	p.giveDolls()
	p.giveFlowers()
	p.giveChocolate()
}
