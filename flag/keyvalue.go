package cliflag

// KeyValue represents a single key-value pair.
//
// It stores a single key-value pair.
//
// If more than one key-value pair is attempted to be stored in it, it will return an error.
//
// You would use a ‘&cliflag.KeyValue’ (i.e., a pointer to a cliflag.KeyValue) with cliflag.Parse(),
// as the target, if you didn't want to use a ‘&map[string]string’, a ‘&map[string]interface{}’, or
// something that fits the ‘cliflag.Storer’ interface, as the target to cliflag.Parse().
//
// Example
//
//	var keyvalue cliflag.KeyValue
//	
//	remainingTokens, err := cliflag.Parse(&keyvalue, tokens...)
//	
//	key, value, err := keyvalue.Unwrap()
type KeyValue struct {
	key string
	value string
	loaded bool
}

// Store makes cliflag.KeyValue fit the cliflag.Storer interface.
//
// Note that if you try to store more than one key-value pair inside
// of cliflag.KeyValue, it will return an error.
//
// Example
//
//	var keyvalue cliflag.KeyValue
//	
//	remainingTokens, err := cliflag.Parse(&keyvalue, tokens...)
//	
//	key, value, err := keyvalue.Unwrap()
func (receiver *KeyValue) Store(key string, value interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	if receiver.loaded {
		return errLoaded
	}

	castedValue, casted := value.(string)
	if !casted {
		return errUnsupportedValue
	}

	receiver.loaded = true
	receiver.key    = key
	receiver.value  = castedValue

	return nil
}

// Unwrap lets you get the ‘key’, and the ‘value’ of the key-value pair out of the cliflag.KeyValue.
//
// Note that if the cliflag.KeyValue has not been loaded, then it returns an error.
//
// Example
//
//	var keyvalue cliflag.KeyValue
//	
//	remainingTokens, err := cliflag.Parse(&keyvalue, tokens...)
//	
//	key, value, err := keyvalue.Unwrap()
func (receiver KeyValue) Unwrap() (key string, value string, err error) {
	if !receiver.loaded {
		return "", "", errNotLoaded
	}

	return receiver.key, receiver.value, nil
}
