package cliflags

import (
	"fmt"
)

type storerMapStringInterface struct {
	p *map[string]interface{}
}

func (receiver *storerMapStringInterface) Store(key string, value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	p := receiver.p
	if nil == p {
		return errNilTarget
	}

	if nil == *p {
		*p = map[string]interface{}{}
	}

	foundValueInterface, found := (*p)[key]
	if found {
		var foundValue string
		switch casted := foundValueInterface.(type) {
		case string:
			foundValue = casted
		case fmt.Stringer:
			foundValue = casted.String()
		case fmt.GoStringer:
			foundValue = casted.GoString()
		default:
			foundValue = fmt.Sprintf("%#v", casted)
		}

		return internalKeyFound{
			key: key,
			value: value,
			foundValue: foundValue,
		}
	}

	(*p)[key] = value

	return nil
}

