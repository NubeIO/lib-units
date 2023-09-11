package units

import (
	"fmt"
	"testing"
)

func TestGetBACnetUnitByValue(t *testing.T) {
	fmt.Println(unitsList.percent)
	GetBACnetUnitByValue(unitsList.percent)

	fmt.Println(BACnetUnitsNames())
}
