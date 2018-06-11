package logeric

// Level represents the logging level of an entry.
type Level uint

const (
	// Panic is the highest priority critical log entry and will cause the application to panic.
	Panic Level = iota
	// Fatal is a high priority critical log entry and will cause the application to quit.
	Fatal
	// Error is a medium priority log entry which should be noted and fixed.
	Error
	// Warn is a serious log entry that should be noted and may be important.
	Warn
	// Info is a low priority informational log entry and may be ignored.
	Info
	// Debug is a very spammy log entry that should probably be ignored / filtered out.
	Debug
)

// String returns the name of a Level.
func (l Level) String() string {
	switch l {
	case Panic:
		return "panic"
	case Fatal:
		return "fatal"
	case Error:
		return "error"
	case Warn:
		return "warn"
	case Info:
		return "info"
	case Debug:
		return "debug"
	}
	return ""
}
