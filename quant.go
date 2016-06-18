package quant

import "fmt"

type Power int64

const (
	Watt     Power = 1
	Kilowatt       = 1000 * Watt
	Megawatt       = 1000 * Kilowatt
	Gigawatt       = 1000 * Megawatt
)

type Energy int64

const (
	WattSecond   Energy = 1
	KilowattHour        = 3.6e+6 * WattSecond
)

func (p Power) String() string {
	precision := 2
	gw := float64(p) / float64(Gigawatt)
	mw := float64(p) / float64(Megawatt)
	kw := float64(p) / float64(Kilowatt)

	if gw >= 1 {
		return unitString(gw, "GW", precision)
	}

	if mw >= 1 {
		return unitString(mw, "MW", precision)
	}

	if kw >= 1 {
		return unitString(kw, "kW", precision)
	}

	return unitString(float64(p), "W", precision)
}

func (e Energy) String() string {
	precision := 2
	kwh := float64(e) / float64(KilowattHour)

	if kwh >= 1 {
		return unitString(kwh, "kWh", precision)
	}

	return unitString(float64(e), "Ws", precision)
}

func unitString(val float64, units string, precision int) string {
	i := int64(val)
	if val == float64(i) {
		return fmt.Sprintf("%d %s", i, units)
	}

	formatStr := fmt.Sprintf("%%0.%df %%s", precision)
	return fmt.Sprintf(formatStr, float64(val), units)
}
