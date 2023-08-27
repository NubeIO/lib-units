package units

import "github.com/martinlindhe/unit"

// https://www.unitconverters.net/power-converter.html

type PowerUnit struct {
	unitCommon
	power unit.Power
	to    func(power unit.Power) float64
}

// FromFloat implements SimpleUnit
func (mu *PowerUnit) FromFloat(f float64) UnitVal {
	return PowerVal{unit.Power(f) * mu.power, mu}
}

var (
	Watt     = &PowerUnit{"w", unit.Watt, unit.Power.Watts}
	Kilowatt = &PowerUnit{"kw", unit.Kilowatt, unit.Power.Kilowatts}
	Megawatt = &PowerUnit{"mw", unit.Megawatt, unit.Power.Megawatts}
)

type PowerVal struct {
	V unit.Power
	U *PowerUnit
}

func (mv PowerVal) String() string {
	return simpleUnitString(mv.U.to(mv.V), mv.U)
}

func (mv PowerVal) AsFloat() float64 {
	return mv.U.to(mv.V)
}

// Convert implements UnitVal conversion
func (mv PowerVal) Convert(to UnitType) (UnitVal, error) {
	if to, ok := to.(*PowerUnit); ok {
		mv.U = to
		return mv, nil
	}
	return nil, ErrorConversion{mv.U, to}
}
