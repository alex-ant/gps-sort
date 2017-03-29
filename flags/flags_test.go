package flags

import "testing"

func TestFlags(t *testing.T) {
	if Values.InputMode == "" ||
		Values.InputFile == "" ||
		Values.MySQLHost == "" ||
		Values.MySQLPort == 0 ||
		Values.MySQLUser == "" ||
		Values.MySQLPass == "" ||
		Values.MySQLDatabase == "" ||
		Values.TopAmount == 0 ||
		Values.ComparisonPointLat == 0 ||
		Values.ComparisonPointLng == 0 {
		t.Error("failed to set default flag values")
	}
}
