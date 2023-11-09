package FlowUnitTransform

import (
	"fmt"
	"testing"
)

func TestTransformUnit(t *testing.T) {
	a := TransformUnit("648124")
	fmt.Println(a)
}
