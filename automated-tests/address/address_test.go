package address

import (
	"testing"
)

type testCase struct {
	addressInsert  string
	expectedReturn string
}

// sempre começar com Test
func TestTypeAddress(t *testing.T) {

	testCases := []testCase{
		{"Rua abc", "Rua"},
		{"Avenida abc", "Avenida"},
		{"Rodovia xpto", "Rodovia"},
		{"Praça xpto", "Invalid Address"},
	}

	for _, caset := range testCases {
		typeAddresReceived := AddressType(caset.addressInsert)
		if typeAddresReceived != caset.expectedReturn {
			t.Errorf("The type %s received is different from what was expected %s",
				typeAddresReceived,
				caset.expectedReturn)
		}
	}
}
