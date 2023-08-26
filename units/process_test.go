package units

import (
	"fmt"
	"github.com/NubeIO/lib-units/pprint"
	"testing"
)

func TestProcess(t *testing.T) {
	process, u, err := Process(25.1, "c", "f")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(process)
	fmt.Println(u.String())
	fmt.Println(u.AsFloat())
}

func TestSupportedUnits(t *testing.T) {
	// uu := UnitsMap()
	unitsMap := SupportedUnits()
	// fmt.Println(unitsMap)

	pprint.PrintJSON(unitsMap)
}
