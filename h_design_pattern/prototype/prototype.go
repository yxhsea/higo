/*
 Prototype原型模式：
        用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象
 个人想法：注意浅复制、深复制
*/
package prototype

import "fmt"

type Resume struct {
	name     string
	sex      string
	age      string
	timeArea string
	company  string
}

func (r *Resume) setPersonInfo(name, sex, age string) {
	if r == nil {
		return
	}
	r.name = name
	r.sex = sex
	r.age = age
}

func (r *Resume) setWorkExperience(timeArea, company string) {
	if r == nil {
		return
	}
	r.timeArea = timeArea
	r.company = company
}

func (r *Resume) display() {
	if r == nil {
		return
	}
	fmt.Println("个人信息: ", r.name, r.sex, r.age)
	fmt.Println("工作经历: ", r.timeArea, r.company)
}

func (r *Resume) clone() *Resume {
	if r == nil {
		return nil
	}

	//深复制
	//return r

	//浅复制
	newObj := (*r)
	return &newObj
}

func NewResume() *Resume {
	return &Resume{}
}
