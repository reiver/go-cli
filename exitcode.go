package cli

type ExitCode struct {
	value uint8
}

const (
	// This is the “success” or “OK” ‘exit code’.
	//
	// This is the only ‘exit code’ that does not represent an error.
	//
	// If everything went OK with the execution of your software, then you should return this ‘exit code’.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode {
	//	
	//		// ...
	//	
	//		return cli.ExitCodeOK
	//	}
	ExitCodeOK = ExitCode{0}

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
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode {
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
	ExitCodeError = ExitCode{1}

	// You would return this error — “bad request” — if the request the user makes via the command line has an error in it.
	//
	// For example, you would return this “error” ‘exit code’ if your software is used incorrectly by the user by:
	//	• providing the wrong number or command line arguments, or
	//	• providing a bag flag, or
	//	• providing bad syntax with a flag, or
	//	• etc.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode {
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
	ExitCodeBadRequest = ExitCode{64}

	// You would return this error — “bad input” — if the input data was bad in some way.
	//
	// For example, you would return this “error” ‘exit code’ if the bad input data was given to your software via:
	//	• Stdin (i.e., os.Stdin), or
	//	• a user file (ex: it does not exist, or cannot be opened, or has a syntax error), or
	//	• input the user types in (maybe via a form), or
	//	• etc.
	//
	// NOTE that if the input problem relates to an operating system (OS) file (such as “/etc/os-release”, “/etc/passwd”, “/var/run/utmp”, etc)
	// then cli.ExitCodeOSFileError should be returned instead of cli.ExitCodeBadInput.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode {
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
	ExitCodeBadInput = ExitCode{65}

	// You would return this error — “internal error” — if an internal error has been detected in your software.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode {
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
	ExitCodeInternalError = ExitCode{70}

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
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode {
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
	ExitCodeOSError = ExitCode{71}

	// You would return this error — “OS file error” — if an operating system (OS) file that you expect to be there
	// (such as “/etc/os-release”, “/etc/passwd”, “/var/run/utmp”, etc) does not exist, or cannot be opened, or has a syntax error.
	//
	// Example
	//
	//	func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) ExitCode {
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
	ExitCodeOSFileError = ExitCode{72}


)
