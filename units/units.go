package units

import (
	"fmt"
	"sort"
)

type Units struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	LongName string `json:"long_name"`
	Symbol   string `json:"symbol"`
}

func SupportedUnitsNames() ([]string, []string) {
	var names = []string{"disabled"}
	var symbols = []string{"disabled"}
	for _, units := range SupportedUnits() {
		names = append(names, fmt.Sprintf("%s -> %s", units.Category, units.LongName))
		symbols = append(symbols, units.LongName)
	}
	return names, symbols
}

func SupportedUnits() []Units {
	var out []Units
	for unitType, units := range supportedUnits() {
		unitsLen := len(units)
		var category string
		var unitName string
		var longName string
		var symbol = unitType.String()
		if unitsLen > 0 {
			category = units[0]
		}
		if unitsLen > 1 {
			unitName = units[1]
		}
		if unitsLen > 2 {
			longName = units[2]
		}

		newUnit := Units{
			Category: category,
			Name:     unitName,
			LongName: longName,
			Symbol:   symbol,
		}
		out = append(out, newUnit)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Category < out[j].Category
	})
	return out
}

var supportedUnitsList = map[UnitType][]string{
	// Length
	Meter:        {"length", "m", "meter"},
	Kilometer:    {"length", "km", "kilometer"},
	Millimeter:   {"length", "mm", "millimeter"},
	Centimeter:   {"length", "cm", "centimeter"},
	Nanometer:    {"length", "nm", "nanometer"},
	Inch:         {"length", "in", "inch"},
	FootInch:     {"length", "ft", "feet+inches"},
	Foot:         {"length", "foot", "feet"},
	Yard:         {"length", "yd", "yard"},
	Mile:         {"length", "mi", "mile"},
	Lightyear:    {"length", "ly", "lightyear"},
	NauticalMile: {"length", "nmi", "nautical mile"},

	// Power
	Watt:     {"power", "w", "watt"},
	Kilowatt: {"power", "kw", "kilowatt"},
	Megawatt: {"power", "mw", "megawatt"},

	// Mass
	Gram:     {"mass", "g", "gram"},
	Kilogram: {"mass", "kg", "kilogram"},
	Pound:    {"mass", "lb", "lbs", "pound"},
	Stone:    {"mass", "st", "stone"},

	// Temperature
	Celsius:    {"temperature", "c", "celsius"},
	Fahrenheit: {"temperature", "f", "fahrenheit"},
	Kelvin:     {"temperature", "k", "kelvin"},

	// Speed
	MilesPerHour:      {"speed", "mph", "mp/h"},
	KilometersPerHour: {"speed", "kmh", "km/h"},

	// Volume
	Liter:      {"volume", "l", "liter"},
	Centiliter: {"volume", "cl", "centiliter"},
	Milliliter: {"volume", "ml", "milliliter"},
	Gallon:     {"volume", "gal", "gallon"},
	Quart:      {"volume", "qt", "quart"},
	Pint:       {"volume", "pt", "pint"},
	Cup:        {"volume", "cup", "cups"},
	FlOunce:    {"volume", "oz", "ounce"},
	Tablespoon: {"volume", "tbsp", "tablespoon"},
	Teaspoon:   {"volume", "tsp", "teaspoon"},

	// Duration
	Second: {"duration", "s", "second"},
	Minute: {"duration", "min", "minute"},
	Hour:   {"duration", "hr", "hour"},
	Day:    {"duration", "day", "days"},
	Week:   {"duration", "wk", "week"},
	Month:  {"duration", "month", "months"},
	Year:   {"duration", "yr", "year"},
}

// SupportedUnits returns all the supported unit types mapped to a list of aliases
func supportedUnits() map[UnitType][]string {
	unitLock.RLock()
	defer unitLock.RUnlock()
	supported := make(map[UnitType][]string, len(supportedUnitsList))
	for unit, aliases := range supportedUnitsList {
		supported[unit] = aliases
	}
	return supported
}
