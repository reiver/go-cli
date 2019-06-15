package cliautocomplete

import (
	"github.com/reiver/go-cli/autocomplete/internal/optint64"
	"github.com/reiver/go-cli/autocomplete/internal/optstring"


	"strconv"
	"strings"
)

type comp struct {
	LINE  optstring.Type
	POINT optint64.Type
	TYPE  optstring.Type
	KEY   optstring.Type
}

// It is expected that you effectively call this with:
//
//	COMP := compParse(os.Environ())
//
// The input to this is a []string, and will contain the environment variables.
//
// (I.e., what should be the returned value from os.Environ().)
//
// For example, it might look something like:
//
//	[]string{
//		"LANG=en_US.UTF-8",
//		"COMP_KEY=9",                    // <---- COMP_KEY
//		"COMP_LINE=myprogram init --na", // <---- COMP_LINE
//		"COMP_POINT=19",                 // <---- COMP_POINT
//		"COMP_TYPE=63",                  // <---- COMP_TYPE
//		"TERM=xterm-256color",
//		"SHELL=/bin/bash",
//	}
//
// (Actually, it reality it will probably have a lot more entries. But I didn't include everyting,
// to make the example more readable.)
//
// So, this functions has to parse our the value the environment variables it cares about.
//
// And the environment variable it cares about are:
//
// • COMP_LINE
// • COMP_POINT
// • COMP_TYPE
// • COMP_KEY
//
// The are the environment variables relevant for autocomplete.
func compParse(data []string) comp {

	const COMP_LINE_  = "COMP_LINE="
	const COMP_POINT_ = "COMP_POINT="
	const COMP_TYPE_  = "COMP_TYPE="
	const COMP_KEY_   = "COMP_KEY="

	var COMP comp

	var found_COMP_LINE  bool
	var found_COMP_POINT bool
	var found_COMP_TYPE  bool
	var found_COMP_KEY   bool

	for _, datum := range data {
		switch {
		case strings.HasPrefix(datum, COMP_LINE_):
			COMP.LINE = optstring.Some(datum[len(COMP_LINE_):])
			found_COMP_LINE = true

		case strings.HasPrefix(datum, COMP_POINT_):
			s := datum[len(COMP_POINT_):]

			i64, err := strconv.ParseInt(s, 10, 64)
			if nil != err {
//@TODO
				continue
			}

			COMP.POINT = optint64.Some(i64)

			found_COMP_POINT = true

		case strings.HasPrefix(datum, COMP_TYPE_):
			COMP.TYPE = optstring.Some(datum[len(COMP_TYPE_):])

			found_COMP_TYPE = true

		case strings.HasPrefix(datum, COMP_KEY_):
			COMP.KEY  = optstring.Some(datum[len(COMP_KEY_):])

			found_COMP_KEY = true
		}

		// SHORTCUT!!!
		if found_COMP_LINE && found_COMP_POINT && found_COMP_TYPE && found_COMP_KEY {
			break
		}
	}

	return COMP
}
