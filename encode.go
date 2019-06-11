package cli

import (
	"io"
	"unicode/utf8"
)

// encode is used by canonicalCommand() to turn a ‘[]string’ into a ‘string’.
//
// Basically canonicalCommand() uses “\x1F” (ASCII/Unicode ‘unit separator’) as
// a special character.
//
// So encode() escapes “\x1F” by turning it into “\x1B\x1F”.
//
// And, since “\x1B” is the escape character....
// also encode() escapes “\x1B” by turning it into “\x1B\x1B”.
func encode(writer io.Writer, str string) error {
	s := str

	for {
		if "" == s {
			break
		}

		r, size := utf8.DecodeRuneInString(s)
		if utf8.RuneError == r && 0 == size {
			break
		}
		if utf8.RuneError == r && 1 == size {
			return errInvalidUTF8
		}

		switch r {
		case '\x1B':
			var buffer [2]byte = [2]byte{'\x1B', '\x1B'}
			_, err := writer.Write(buffer[:])
			if nil != err {
				return err
			}
		case '\x1F':
			var buffer [2]byte = [2]byte{'\x1B', '\x1F'}
			_, err := writer.Write(buffer[:])
			if nil != err {
				return err
			}
		default:
			var buffer [utf8.UTFMax]byte
			n := utf8.EncodeRune(buffer[:], r)

			_, err := writer.Write(buffer[:n])
			if nil != err {
				return err
			}
		}

		s = s[size:]
	}

	return nil
}
