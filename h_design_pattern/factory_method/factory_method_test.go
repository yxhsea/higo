package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	var tmp *OperationAdd = nil
	if val, ok := tmp.Result(); ok == true {
		t.Error(val)
	}

	var of OperationFunc
	cAdd := COperationAdd{}
	of = cAdd.CreateOperation("+")
	of.SetNumA(10)
	of.SetNumB(110)
	if val, ok := of.Result(); ok == true && val != 120 {
		t.Error("Add Error")
	}

	cSub := COperationSub{}
	of = cSub.createOperation("-")
	of.SetNumA(10)
	of.SetNumB(110)
	if val, ok := of.Result(); ok == true && val != -100 {
		t.Error("Sub Error")
	}
}
