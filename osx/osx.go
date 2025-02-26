package osx

type ExitCode = int

const (
	ExitSuccess       ExitCode = 0
	ExitFailure       ExitCode = 1
	ExitUsageError    ExitCode = 2
	ExitConfigError   ExitCode = 3
	ExitDatabaseError ExitCode = 4
	ExitIOError       ExitCode = 5
	ExitTimeout       ExitCode = 6
	ExitNetworkError  ExitCode = 7
	ExitAuthError     ExitCode = 8
	ExitPermission    ExitCode = 126
	ExitNotFound      ExitCode = 127
	ExitSignalBase    ExitCode = 128
	ExitInterrupted   ExitCode = ExitSignalBase + 2
	ExitKilled        ExitCode = ExitSignalBase + 9
	ExitSegmentFault  ExitCode = ExitSignalBase + 11
)
