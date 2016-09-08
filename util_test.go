package dvb

import "testing"

func TestNormalize(t *testing.T) {
	uao := "üäö"
	expected := "uao"
	normalize(&uao)
	if uao != expected {
		t.Error("Expected", expected, "got", uao)
	}
}
