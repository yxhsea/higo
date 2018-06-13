/*
  Composite 组合模式：
        将对象组合成树形结构，以表示“部分-整体”的层次结构。
		组合模式使得用户对单个对象和组合对象的使用具有一致性
 个人想法：
*/
package composite

import (
	"fmt"
	"strings"
)

//公司管理接口
type Company interface {
	add(Company)
	remove(Company)
	display(int)
	lineOfDuty()
}

type RealCompany struct {
	name string
}

//具体公司
type ConCreateCompany struct {
	RealCompany
	list []Company
}

func NewConCreateCompany(name string) *ConCreateCompany {
	return &ConCreateCompany{RealCompany{name}, []Company{}}
}

func (c *ConCreateCompany) add(newC Company) {
	if c == nil {
		return
	}
	c.list = append(c.list, newC)
}

func (c *ConCreateCompany) remove(delC Company) {
	if c == nil {
		return
	}
	for i, val := range c.list {
		if val == delC {
			c.list = append(c.list[:i], c.list[i+1:]...)
			return
		}
	}
	return
}

func (c *ConCreateCompany) display(depth int) {
	if c == nil {
		return
	}
	fmt.Println(strings.Repeat("-", depth), " ", c.name)
	for _, val := range c.list {
		val.display(depth + 2)
	}
}

func (c *ConCreateCompany) lineOfDuty() {
	if c == nil {
		return
	}
	for _, val := range c.list {
		val.lineOfDuty()
	}
}

//人力资源部门
type HRDepartment struct {
	RealCompany
}

func NewHRDepartment(name string) *HRDepartment {
	return &HRDepartment{RealCompany{name}}
}

func (h *HRDepartment) add(c Company)    {}
func (h *HRDepartment) remove(c Company) {}
func (h *HRDepartment) display(depth int) {
	if h == nil {
		return
	}
	fmt.Println(strings.Repeat("-", depth), " ", h.name)
}

func (h *HRDepartment) lineOfDuty() {
	if h == nil {
		return
	}
	fmt.Println(h.name, "员工招聘培训管理")
}

//财务管理
type FinanceDepartment struct {
	RealCompany
}

func NewFinanceDepartment(name string) *FinanceDepartment {
	return &FinanceDepartment{RealCompany{name}}
}

func (h *FinanceDepartment) add(c Company)    {}
func (h *FinanceDepartment) remove(c Company) {}
func (h *FinanceDepartment) display(depth int) {
	if h == nil {
		return
	}
	fmt.Println(strings.Repeat("-", depth), " ", h.name)
}

func (h *FinanceDepartment) lineOfDuty() {
	if h == nil {
		return
	}
	fmt.Println(h.name, "公司财务收支管理")
}
