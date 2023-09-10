package units

import (
	"fmt"
	"testing"
)

func TestGetBACnetUnitByValue(t *testing.T) {
	fmt.Println(unitsList.percent)
	fmt.Println(GetBACnetUnitByValue(unitsList.percent))
	fmt.Println(BACnetUnits())
}
