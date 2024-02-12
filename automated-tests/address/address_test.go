package address_test

import (
	"automated-tests/addres"
	"testing"
)

type testCase struct {
	addressInsert  string
	expectedReturn string
}

// sempre começar com Test + inicio da palavra com letra ^
func TestTypeAddress(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{"Rua abc", "Rua"},
		{"Avenida abc", "Avenida"},
		{"Rodovia xpto", "Rodovia"},
		{"Praça xpto", "Invalid Address"},
	}

	for _, caset := range testCases {
		typeAddresReceived := addres.AddressType(caset.addressInsert)
		if typeAddresReceived != caset.expectedReturn {
			t.Errorf("The type %s received is different from what was expected %s",
				typeAddresReceived,
				caset.expectedReturn)
		}
	}
}

func TestAny(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("test broken.")
	}
}
