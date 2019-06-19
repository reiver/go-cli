package cliflag

type storerMapStringString struct {
	p *map[string]string
}

func (receiver *storerMapStringString) Store(key string, value interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	castedValue, casted := value.(string)
	if !casted {
		return errUnsupportedValue
	}

	p := receiver.p
	if nil == p {
		return errNilTarget
	}

	if nil == *p {
		*p = map[string]string{}
	}

	_, found := (*p)[key]
	if found {
		return errFound
	}

	(*p)[key] = castedValue

	return nil
}

