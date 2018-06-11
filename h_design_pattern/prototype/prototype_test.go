package prototype

import (
	"fmt"
	"testing"
)

func TestProtoType(t *testing.T) {
	resume := NewResume()

	resume.setPersonInfo("John", "男", "22")
	resume.setWorkExperience("2", "Apple")
	resume.display()

	fmt.Println("==========================================")

	cloneResume := resume.clone()
	cloneResume.display()

	fmt.Println("==========================================")

	cloneResume.setPersonInfo("Tom", "女", "22")
	cloneResume.setWorkExperience("2", "HW")
	cloneResume.display()

	fmt.Println("==========================================")
	resume.display()
}
