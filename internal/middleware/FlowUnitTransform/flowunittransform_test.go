package FlowUnitTransform

import (
	"fmt"
	"testing"
)

func TestTransformUnit(t *testing.T) {
	a := TransformUnit("12687")
	fmt.Println(a)
}
