package cliflags

import (
	"github.com/reiver/go-cli/flag"
)

func Parse(target interface{}, tokens ...string) ([]string, error) {
	if nil == target {
		return tokens, errNilTarget
	}

	switch casted := target.(type) {
	default:
		return tokens, internalUnsupportedTarget{target}

	case Storer:
		return parse(casted, tokens...)

	case *map[string]string:
		storer := &storerMapStringString{p: casted}

		return parse(storer, tokens...)

	case *map[string]interface{}:
		storer := &storerMapStringInterface{p: casted}

		return parse(storer, tokens...)
	}
}

func parse(storer Storer, tokens ...string) ([]string, error) {
	if nil == storer {
		return tokens, errNilTarget
	}

	if 1 > len(tokens) {
		return tokens, nil
	}

	for 0 < len(tokens) {

		var err error

		before := len(tokens)

		tokens, err = cliflag.Parse(storer, tokens...)
		if nil != err {
			switch err.(type) {
			case cliflag.EndOfFlags:
				return tokens, nil
			default:
				return tokens, err
			}
		}

		after := len(tokens)

		if after >= before {
			return tokens, errInternalError
		}
	}

	return tokens, nil
}
