package cliflags

import (
	"sync"
)

// KeyValues represents key-value pairs.
//
// It stores key-value pairs.
//
// You would use a ‘&cliflags.KeyValues’ (i.e., a pointer to a cliflags.KeyValues) with cliflags.Parse(),
// as the target, if you didn't want to use a ‘&map[string]string’, a ‘&map[string]interface{}’, or
// something that fits the ‘cliflags.Storer’ interface, as the target to cliflags.Parse().
//
// Example
//
//	var keyvalues cliflags.KeyValues
//	
//	remainingTokens, err := cliflags.Parse(&keyvalue, tokens...)
//	
//	kevalues.For(func(key string, value string){
//		fmt.Printf("key = %q\n", keu)
//		fmt.Printf("value = %q\n", value)
//		fmt.Println()
//	})
type KeyValues struct {
	mutex sync.RWMutex
	data map[string]string
}

func (receiver *KeyValues) For(fn func(string,string)) {
	if nil == receiver {
		return
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	for k, v := range receiver.data {
		fn(k,v)
	}
}

func (receiver *KeyValues) Len() int {
	if nil == receiver {
		return 0
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return len(receiver.data)
}

func (receiver *KeyValues) Load(key string) Value {
	if nil == receiver {
		return NoValue()
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	data := receiver.data
	if nil == data {
		return NoValue()
	}

	value, found := data[key]
	if !found {
		return NoValue()
	}

	return SomeValue(value)
}

func (receiver *KeyValues) Store(key string, value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.data {
		receiver.data = map[string]string{}
	}

	foundValue, found := receiver.data[key]
	if found {
		return internalKeyFound{
			key:key,
			value:value,
			foundValue:foundValue,
		}
	}

	receiver.data[key] = value

	return nil
}
