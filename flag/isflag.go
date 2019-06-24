package cliflag

// IsFlag returns true if ‘value’ is cli flag, else returns false otherwise.
//
// Flags can begin with either a ‘-’ character, or a ‘+’ character.
//
// Ex:
//
//	-n
//
//	-n value
//
//	-n=value
//
//	-n:value
//
//	-nvalue
//
//	--name
//
//	--name=value
//
//	--name:value
//
//	+name
//
//	+name=value
//
//	+name:value
//
// But note that no everything that begins with a ‘-’ character, or a ‘+’ character is a flag.
//
// Ex:
//
//	-
//
//	--
//
//	---
//
//	----
//
//	+
//
//	++
//
//	+++
//
//	++++
//
//	---n
//
//	---name
//
//	---name=value
//
//	---name:value
//
//	++n
//
//	++name
//
//	++name=value
//
//	++name:value
func IsFlag(value string) bool {
	if 2 > len(value) {
		return false
	}

	c0 := value[0]
	c1 := value[1]

	if 2 == len(value) {
		switch {

		// -n
		case '-' == c0 && '-' != c1:
			return true

		// +n
		case '+' == c0 && '+' != c1:
			return true

		default:
			return false
		}
	}

	c2 := value[2]

	switch {

	// -n
	// -nvalue
	// -n=value
	// -n:value
	case '-' == c0 && '-' != c1:
		return true

	// --name
	// --name=value
	// --name:value
	case '-' == c0 && '-' == c1 && '-' != c2:
		return true

	// +name
	// +name=value
	// +name:value
	case '+' == c0 && '+' != c1:
		return true

	default:
		return false
	}
}
