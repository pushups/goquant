package quant

import "testing"

func testString(s string, expected string, t *testing.T) {
	if s != expected {
		t.Error("Expected", expected, "got", s)
	}
}

func TestWattsToString(t *testing.T) {
	w := 4 * Watt

	testString(w.String(), "4 W", t)
}

func TestKWToString(t *testing.T) {
	w := 1000 * Watt

	testString(w.String(), "1 kW", t)
}

func TestMWToString(t *testing.T) {
	w := 3000 * 1000 * Watt

	testString(w.String(), "3 MW", t)
}

func TestGWToString(t *testing.T) {
	w := 3000 * 1000 * 1000 * Watt

	testString(w.String(), "3 GW", t)
}

func TestWattConversion(t *testing.T) {
	a := 1 * Kilowatt
	b := 1000 * Watt
	if a != b {
		t.Error("Expected", a, "to equal", b)
	}
}

func TestEnergy(t *testing.T) {
	ws := 240 * WattSecond
	testString(ws.String(), "240 Ws", t)
}

func TestKilowattHours(t *testing.T) {
	kwh := 12 * KilowattHour
	testString(kwh.String(), "12 kWh", t)
}

func TestEnergyConversion(t *testing.T) {
	kwh := 5.6e+6 * WattSecond
	ws := 2.5e+3 * WattSecond
	testString(kwh.String(), "1.56 kWh", t)
	testString(ws.String(), "2500 Ws", t)
}
