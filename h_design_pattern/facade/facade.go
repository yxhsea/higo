/*
  Facade 外观模式：
        为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，
		这个接口使得这一子系统更加容易使用（投资：基金，股票，房产）
 个人想法：中介者模式、外观模式：每个对象都保存一份中介者对象，
        在和其他对象交互时，通过中介者来完成，外观模式：外观中保存了一堆对象，
		这些对象或者是组成某个子系统的，将其封装在外观对象中，给客户端一种只有一个对象的感觉，
		一个是结构型模式，一个是行为性模式
*/

package facade

import "fmt"

type FuncOne struct {
	str string
}

func (f FuncOne) Out() {
	fmt.Println("funcone", f.str)
}

type FuncTwo struct {
	i int
}

func (f FuncTwo) Out() {
	fmt.Println("functwo", f.i)
}

type FuncThree struct {
	f float32
}

func (f FuncThree) Out() {
	fmt.Println("functhree", f.f)
}

type Facade struct {
	One   FuncOne
	Two   FuncTwo
	Three FuncThree
}

func (f Facade) OutOne() {
	f.One.Out()
	f.Three.Out()
}

func (f Facade) OUtTwo() {
	f.Two.Out()
	f.Three.Out()
}

func NewFacade(i int, f float32, str string) *Facade {
	return &Facade{FuncOne{str}, FuncTwo{i}, FuncThree{f}}
}
