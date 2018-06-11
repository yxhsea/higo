package abstract_factory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	uData := User{1, "u"}
	dData := Department{1, "d"}
	data := DataAccess{}

	iU := data.createUser("access")
	iU.insert(&uData)
	gU := iU.getUser(1)
	fmt.Println(gU)

	fmt.Println("==============================")

	iD := data.createDepartment("sqlServer")
	iD.insert(&dData)
	gD := iD.getDepartment(1)
	fmt.Println(gD)

	if iS := data.createDepartment("a"); iS != nil {
		t.Error()
	}
}
