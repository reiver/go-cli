package cliflag

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

	_, found := (*p)[key]
	if found {
		return errFound
	}

	(*p)[key] = value

	return nil
}

