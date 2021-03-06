package cli

import (
	"fmt"
)

type ExitCode uint8

func (receiver ExitCode) Code() uint8 {
	return uint8(receiver)
}

func (receiver ExitCode) GoString() string {
	return fmt.Sprintf("cli.ExitCode(%d)", receiver)
}

func (receiver ExitCode) String() string {
	switch receiver {
	case ExitCodeOK:
		return "OK"
	case ExitCodeError:
		return "Error"
	case ExitCodeBadRequest:
		return "Bad Request"
	case ExitCodeBadInput:
		return "Bad Input"
	case ExitCodeNoInput:
		return "No Input"
	case ExitCodeNoUser:
		return "No User"
	case ExitCodeNoHost:
		return "No Host"
	case ExitCodeUnavailable:
		return "Unavailable"
	case ExitCodeInternalError:
		return "Internal Error"
	case ExitCodeOSError:
		return "OS Error"
	case ExitCodeOSFileError:
		return "OS File Error"
	case ExitCodeNoOutput:
		return "No Output"
	case ExitCodeIOError:
		return "I/O Error"
	case ExitCodeTemporaryFailure:
		return "Temporary Failure"
	case ExitCodeProtocolError:
		return "Protocol Error"
	case ExitCodeUnauthorized:
		return "Unauthorized"
	case ExitCodeBadConfiguration:
		return "Bad Configuration"

	default:
		return fmt.Sprintf("Unrecognized Error: ‘%d’", receiver)
	}
}

	// This is the “success” or “OK” ‘exit code’.
	//
	// This is the only ‘exit code’ that does not represent an error.
	//
	// If everything went OK with the execution of your software, then you should return this ‘exit code’.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		return cli.ExitCodeOK
	//	}
	//
	// On other systems, this is known as: EX_OK
	const ExitCodeOK = ExitCode(0)

	// This is a generic “error” ‘exit code’.
	//
	// NOTE that there are more specific “error” ‘exit codes’.
	// For example: “bad request”, “internal error”, “OS error”, “I/O error”, etc.
	//
	// It is recommended that for most situations that if you could accurately use a more specific “error” ‘exit codes’
	// (such as “bad request”, “internal error”, “OS error”, “I/O error”, etc) then you should use that “error” ‘exit code’
	// rather than this generic “error” ‘exit codes’.
	//
	// It is recommended that you should only use this generic “error” ‘exit code’ if you do not want to provide details
	// about the nature of the error, or if an “error” ‘exit code’ does not exist to accurately describe the error that happened.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeError
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_ERROR
	const ExitCodeError = ExitCode(1)

	// You would return this error — “bad request” — if the request the user makes via the command line has an error in it.
	//
	// For example, you would return this “error” ‘exit code’ if your software is used incorrectly by the user by:
	//
	//• providing the wrong number or command line arguments, or
	//
	//• providing a bag flag, or
	//
	//• providing bad syntax with a flag, or
	//
	//• etc.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeBadRequest
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_USAGE
	const ExitCodeBadRequest = ExitCode(64)

	// You would return this error — “bad input” — if the input data was bad in some way.
	//
	// For example, you would return this “error” ‘exit code’ if the bad input data was given to your software via:
	//
	//• Stdin (i.e., os.Stdin), or
	//
	//• a user file (ex: has a syntax error), or
	//
	//• input the user types in (maybe via a form), or
	//
	//• etc.
	//
	// NOTE that if the input problem relates to an operating system (OS) file (such as “/etc/os-release”, “/etc/passwd”, “/var/run/utmp”, etc)
	// then cli.ExitCodeOSFileError should be returned instead of cli.ExitCodeBadInput.
	//
	// Also NOTE that if the input user file does not exist, or cannot be opened, then cli.ExitCodeNoInput
	// should be returned instead of cli.ExitCodeBadInput.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeBadInput
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_DATAERR
	const ExitCodeBadInput = ExitCode(65)

	// You would return this error — “no input” — if the input data does not exist, or cannot be opened.
	//
	// For example, you would return this “error” ‘exit code’ if:
	//
	//• a user file that your software expected did not exist, or
	//
	//• a user file (existed) by could not be opened, or
	//
	//• etc.
	//
	// NOTE that if the input problem relates to an operating system (OS) file (such as “/etc/os-release”, “/etc/passwd”, “/var/run/utmp”, etc)
	// then cli.ExitCodeOSFileError should be returned instead of cli.ExitCodeNoInput.
	//
	// Also NOTE that if the input user file (exists, and can indeed be opened, bu) has a syntax error,
	// then cli.ExitCodeBadInput should be returned instead of cli.ExitCodeNoInput.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeNoInput
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_NOINPUT
	const ExitCodeNoInput = ExitCode(66)

	// You would return this error — “no user” — if the specified user does not exist.
	//
	// For example, you would return this “error” ‘exit code’ if:
	//
	//• a user for a remote login did not exist, or
	//
	//• an e-mail address (which represents a user) did not exist, or
	//
	//• etc.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeNoUser
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_NOUSER
	const ExitCodeNoUser = ExitCode(67)

	// You would return this error — “no host” — if the specified host does not exist.
	//
	// For example, you would return this “error” ‘exit code’ if:
	//
	//• the host for a remote login did not exist, or
	//
	//• the host of an e-mail address did not exist, or
	//
	//• the host in a URL did not exist, or
	//
	//• etc.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeNoHost
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_NOHOST
	const ExitCodeNoHost = ExitCode(68)

	// You would return this error — “unavailable” — if something is Unavailable.
	//
	// For example, you would return this “error” ‘exit code’ if:
	//
	//• a service is unavailable, or
	//
	//• etc.
	//
	// You can also use this as a catchall “error” ‘exit code’ when something you wanted does not work,
	// and you don't know why.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeUnavailable
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_UNAVAILABLE
	const ExitCodeUnavailable = ExitCode(69)

	// You would return this error — “internal error” — if an internal error has been detected in your software.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeInternalError
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_SOFTWARE
	const ExitCodeInternalError = ExitCode(70)

	// You would return this error — “OS error” — if an error with the operating system (OS) has been detected by your software.
	//
	// I.e., that something is wrong with the operating system (OS) itself.
	//
	// So, for example, if a ‘temp directory’ does not exist on the operating system (OS) your software is running on
	// (and your software notices this problem via os.TempDir() returning an empty string ("")), then that could be
	// an appropriate situation to return cli.ExitCodeOSError (if your software doesn't have a backup plan
	// to deal that error situation).
	//
	// Another example of something being wrong with your operating system (OS) is if your operating system (OS)
	// has ran out of available file descriptors.
	//
	// Etc.
	//
	// However, for example, if your software was trying to change directories (maybe by calling os.Chdir())
	// to a sub-directory in the user's home directory that does not exist, that would not be an error with
	// the operating system (OS), even though the operating system (OS) reported the error. Since there is
	// nothing wrong with the operating system (OS) in this case.
	//
	// NOTE that if the problem with your operation system (OS) relates to an operating system (OS) file (such as
	// “/etc/os-release”, “/etc/passwd”, “/var/run/utmp”, etc) such that it does not exist, or cannot be opened, or has a syntax error,
	// then cli.ExitCodeOSFileError should be returned instead of cli.ExitCodeOSError.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeOSError
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_OSERR
	const ExitCodeOSError = ExitCode(71)

	// You would return this error — “OS file error” — if an operating system (OS) file that you expect to be there
	// (such as “/etc/os-release”, “/etc/passwd”, “/var/run/utmp”, etc) does not exist, or cannot be opened, or has a syntax error.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeOSFileError
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_OSFILE
	const ExitCodeOSFileError = ExitCode(72)

	// You would return this error — “no output” — if a user specified output file cannot be created.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeIOError
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_CANTCREAT
	const ExitCodeNoOutput = ExitCode(73)

	// You would return this error — “I/O error” — if an error occurred when doing some input, or output on a file.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeIOError
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_IOERR
	const ExitCodeIOError = ExitCode(74)

	// You would return this error — “temporary failure” — if a temporary failture occurred.
	// A temporary failure is not really an error.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeTemporaryFailure
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_TEMPFAIL
	const ExitCodeTemporaryFailure = ExitCode(75)

	// You would return this error — “protocol error” — if a remote system returned something that is not possible
	// according to a protocol exchange.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeProtocolError
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_PROTOCOL
	const ExitCodeProtocolError = ExitCode(76)

	// You would return this error — “unauthorized” — if there is something trying to be done for which
	// you are unauthorized because you have insufficient permissions to do whatever you want to do.
	//
	// Note that this is not meant to be used for file system permissions. Either cli.NoInput, or cli.NoOutput
	// should be used for those instead.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeUnauthorized
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_NOPERM
	const ExitCodeUnauthorized = ExitCode(77)

	// You would return this error — “bad configuration” — if a bad configuration is encountered.
	// Perhaps something is unconfigured. Perhaps something is misconfigured.
	//
	// This could be due to a configuration file; but doesn't have to be.
	// For example, this could be a misconfigured, or unconfigured service of some kind.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) ExitCode {
	//	
	//		// ...
	//	
	//		if nil != err {
	//			return cli.ExitCodeBadConfiguration
	//		}
	//	
	//		// ...
	//	
	//	}
	//
	// On other systems, this is known as: EX_CONFIG
	const ExitCodeBadConfiguration = ExitCode(78)
