package address

import "strings"

func addressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	firstWordAddress := strings.ToLower(strings.Split(address, " ")[0])

	for _, word := range validTypes {
		if word == firstWordAddress {
			return firstWordAddress
		}
	}
	return "Invalid Address"
}
