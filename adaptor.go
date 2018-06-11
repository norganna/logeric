package logeric

import (
	"bytes"
	"strings"
	"io"
)

// Adaptor is a standard logger io.Writer adaptor, which takes the std log and parses it into a logeric instance.
type Adaptor struct {
	logger FieldLogger
}
var _ io.Writer = (*Adaptor)(nil)

// Write is an implementation of io.Writer.
func (l *Adaptor) Write(p []byte) (n int, err error) {
	line := bytes.SplitN(p, []byte{' '}, 4)
	level := string(line[2])
	text := line[3]
	if level[0] != '[' {
		text = append(line[2], ' ')
		text = append(text, line[3]...)
		level = "[INFO]"
	}

	var prefix []byte
	if c := bytes.Index(text, []byte{':'}); c > 0 && c+1 == bytes.Index(text, []byte{' '}) {
		prefix = text[0:c]
		text = text[c+1:]
	}

	logger := l.logger
	if len(prefix) > 0 {
		logger = logger.WithField("prefix", string(prefix))
	}

	out := strings.Trim(string(text), " \t\r\n")
	switch level {
	case "[DEBUG]":
		logger.Debug(out)
	case "[INFO]", "[WARN]":
		logger.Info(out)
	default:
		logger.Warn(out)
	}
	return len(p), nil
}
