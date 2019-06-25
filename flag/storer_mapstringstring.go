package cliflag

type storerMapStringString struct {
	p *map[string]string
}

func (receiver *storerMapStringString) Store(key string, value string) error {
	if nil == receiver {
		return errNilReceiver
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

	(*p)[key] = value

	return nil
}
