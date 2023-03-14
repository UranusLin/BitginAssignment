package cashier

import (
	"testing"
)

func TestChargingSystem(t *testing.T) {
	cs := &ChargingSystem{
		ChargingMap: make(map[string]ChargingMode),
	}
	cs.AddChargingMode("Normal", &NormalCharging{})
	cs.AddChargingMode("VIP", &VipCharging{Discount: 0.9})
	cs.AddChargingMode("Point", &PointCharging{Rate: 1})

	testCases := []struct {
		name        string
		price       float64
		expectedRes float64
	}{
		{"Normal", 1000, 1000},
		{"VIP", 1000, 900},
		{"Point", 100, 100},
	}

	for _, tc := range testCases {
		res := cs.Calculate(tc.name, tc.price)
		if res != tc.expectedRes {
			t.Errorf("Test case %s failed: expected %.2f, but got %.2f", tc.name, tc.expectedRes, res)
		}
	}
}

func TestPointCharging_Calculate(t *testing.T) {
	pc := &PointCharging{
		Rate: 2,
	}

	testCases := []struct {
		limit       float64
		expectedRes float64
	}{
		{1000, 2000},
		{500, 1000},
		{200, 400},
		{0, 0},
	}

	for _, tc := range testCases {
		res := pc.Calculate(tc.limit)
		if res != tc.expectedRes {
			t.Errorf("Test case failed: expected %.2f, but got %.2f", tc.expectedRes, res)
		}
	}
}
