package cliflag

import (
	"strings"
)

// Parse parses command line args tokens, and extracts the first flag, and stores the
// ‘key’, and ‘value’ of that flag, and stores that ‘key’, and ‘value’ in ‘target’.
//
// Overview
//
// For example, if the ‘tokens’ were:
//
//	[]string{"--when=now", "-m", "This is a message.", "filename.txt"}
//
// Then the first token ("--when=now") is the first flag.
//
// And the ‘key’:
//
//	"when"
//
// And the ‘value’ is:
//
//	"now"
//
// Example
//
//	var tokens []string
//	
//	// ...
//	
//	var target cliflag.KeyValue
//	// var target map[string]string
//	// var target map[string]interface{}
//
//	remainingTokens, err := cliflag.Parse(&target, tokens...)
//	
//	if nil != err {
//		switch casted := err.(type) {
//		case cliflag.EndOfFlags:
//			fmt.Fprintf(os.Stderr, "END OF FLAGS: %q is not a flag.\n", casted.Token())
//			return
//		default:
//			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
//			return
//		}
//	}
func Parse(target interface{}, tokens ...string) ([]string, error) {
	if nil == target {
		return tokens, errNilTarget
	}

	switch casted := target.(type) {
	default:
		return tokens, errUnsupportedTarget

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
		return tokens, errNoTokens
	}

	head, tail := tokens[0], tokens[1:]

	if !isFlag(head) {
		return tokens, internalEndOfFlags{token:head}
	}

	{
		c0 := head[0]
		c1 := head[1]

		switch {
		default:
			return tokens, internalEndOfFlags{token:head}

		// --name=value
		// --name:value
		// --name
		case '-' == c0 && '-' == c1:
			var str string = head[2:]

			index := strings.IndexAny(str, "=:")

			var key   string
			var value string
			switch index {
			case -1:
				key   = str
				value = ""
			default:
				key   = str[:index]
				value = str[1+index:]
			}

			if err := storer.Store(key, value); nil != err {
				return tokens, err
			}

			return tail, nil

		// -n
		// -n value
		// -n -- value
		case '-' == c0 && '-' != c1 && 2 == len(head):

			var key string = string(c1)

			switch {
			default:
				var value string = ""

				if err := storer.Store(key, value); nil != err {
					return tokens, err
				}

				return tail, nil

			case 2 <= len(tail) && "--" == tail[0]:

				var value string = tail[1]

				if err := storer.Store(key, value); nil != err {
					return tokens, err
				}

				return tail[2:], nil

			case 0 < len(tail) &&  isFlag(tail[0]):

				var value string = ""

				if err := storer.Store(key, value); nil != err {
					return tokens, err
				}

				return tail, nil

			case 0 < len(tail) && !isFlag(tail[0]):

				var value string = tail[0]

				if err := storer.Store(key, value); nil != err {
					return tokens, err
				}

				return tail[1:], nil
			}

		// -n=value
		// -n:value
		case '-' == c0 && '-' != c1 && 2 < len(head) && ('=' == head[2] || ':' == head[2]):

			var key   string = string(c1)
			var value string = head[3:]

				if err := storer.Store(key, value); nil != err {
					return tokens, err
				}

				return tail, nil

		// -nvalue
		case '-' == c0 && '-' != c1 && 2 < len(head) && ('=' != head[2] && ':' != head[2]):

			var key   string = string(c1)
			var value string = head[2:]

				if err := storer.Store(key, value); nil != err {
					return tokens, err
				}

				return tail, nil

		// +name=value
		// +name:value
		// +name
		case '+' == c0:

			var str string = head[1:]

			index := strings.IndexAny(str, "=:")

			var key   string
			var value string
			switch index {
			case -1:
				key   = str
				value = ""
			default:
				key   = str[:index]
				value = str[1+index:]
			}

			if err := storer.Store(key, value); nil != err {
				return tokens, err
			}

			return tail, nil
		}
	}
}
