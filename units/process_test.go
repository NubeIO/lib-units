package units

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {
	u, err := Process(120, "c", "f")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.String())
	fmt.Println(u.AsFloat())
}

func TestSupportedUnits(t *testing.T) {
	// uu := UnitsMap()
	names, symbols := SupportedUnitsNames()
	// fmt.Println(unitsMap)
	for i, name := range names {
		fmt.Println(name)
		fmt.Println(symbols[i])
	}

	//fmt.Println(symbols)
}
