package logeric

import (
		"errors"
	"fmt"
	"log"
	"os"

	"github.com/norganna/erroneous"
	"github.com/sirupsen/logrus"
)

// Log is a proxy wrapper around another logger.
type Log struct {
	std   StdLogger
	out   OutputLogger
	field logrus.FieldLogger

	ordered bool
}

var _ FieldLogger = (*Log)(nil)
var _ OutputLogger = (*Log)(nil)
var _ StdLogger = (*Log)(nil)

// New creates a new logeric Log.
func New(logger interface{}) (*Log, error) {
	l := &Log{}
	if logger == nil {
		l.out = log.New(os.Stderr, "", log.LstdFlags)
	} else if g, ok := logger.(*Log); ok {
		l = g
	} else if g, ok := logger.(logrus.FieldLogger); ok {
		l.field = g
	} else if g, ok := logger.(OutputLogger); ok {
		l.out = g
	} else if g, ok := logger.(StdLogger); ok {
		l.std = g
	} else {
		return nil, errors.New("unknown logging engine")
	}

	return l, nil
}

// Ordered if set sets a _order field in the output fields which can be used by log formatters to insert-order fields.
func (l *Log) Ordered(set bool) {
	l.ordered = set
}

// WithField is a stub for entry.WithField.
func (l *Log) WithField(key string, value interface{}) *Entry {
	e := &Entry{log: l, stack: 1}
	return e.WithField(key, value)
}

// WithFields is a stub for entry.WithFields.
func (l *Log) WithFields(fields Fields) *Entry {
	e := &Entry{log: l, stack: 1}
	return e.WithFields(fields)
}

// WithError is a stub for entry.WithError.
func (l *Log) WithError(err error) *Entry {
	e := &Entry{log: l, stack: 1}
	return e.WithError(err)
}

// Debugf is a stub for entry.Debugf.
func (l *Log) Debugf(format string, args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Debugf(format, args...)
}

// Infof is a stub for entry.Infof.
func (l *Log) Infof(format string, args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Infof(format, args...)
}

// Printf is a stub for entry.Printf.
func (l *Log) Printf(format string, args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Printf(format, args...)
}

// Warnf is a stub for entry.Warnf.
func (l *Log) Warnf(format string, args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Warnf(format, args...)
}

// Warningf is a stub for entry.Warningf.
func (l *Log) Warningf(format string, args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Warningf(format, args...)
}

// Errorf is a stub for entry.Errorf.
func (l *Log) Errorf(format string, args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Errorf(format, args...)
}

// Fatalf is a stub for entry.Fatalf.
func (l *Log) Fatalf(format string, args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Fatalf(format, args...)
}

// Panicf is a stub for entry.Panicf.
func (l *Log) Panicf(format string, args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Panicf(format, args...)
}

// Debug is a stub for entry.Debug.
func (l *Log) Debug(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Debug(args...)
}

// Info is a stub for entry.Info.
func (l *Log) Info(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Info(args...)
}

// Print is a stub for entry.Print.
func (l *Log) Print(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Print(args...)
}

// Warn is a stub for entry.Warn.
func (l *Log) Warn(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Warn(args...)
}

// Warning is a stub for entry.Warning.
func (l *Log) Warning(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Warning(args...)
}

// Error is a stub for entry.Error.
func (l *Log) Error(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Error(args...)
}

// Fatal is a stub for entry.Fatal.
func (l *Log) Fatal(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Fatal(args...)
}

// Panic is a stub for entry.Panic.
func (l *Log) Panic(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Panic(args...)
}

// Debugln is a stub for entry.Debugln.
func (l *Log) Debugln(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Debugln(args...)
}

// Infoln is a stub for entry.Infoln.
func (l *Log) Infoln(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Infoln(args...)
}

// Println is a stub for entry.Println.
func (l *Log) Println(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Println(args...)
}

// Warnln is a stub for entry.Warnln.
func (l *Log) Warnln(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Warnln(args...)
}

// Warningln is a stub for entry.Warningln.
func (l *Log) Warningln(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Warningln(args...)
}

// Errorln is a stub for entry.Errorln.
func (l *Log) Errorln(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Errorln(args...)
}

// Fatalln is a stub for entry.Fatalln.
func (l *Log) Fatalln(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Fatalln(args...)
}

// Panicln is a stub for entry.Panicln.
func (l *Log) Panicln(args ...interface{}) {
	e := &Entry{log: l, stack: 1}
	e.Panicln(args...)
}

// Err returns a concatenated error from given arguments.
func (l *Log) Err(args ...interface{}) error {
	e := &Entry{log: l, stack: 1}
	return e.Err(args...)
}

// Errf returns a formatted error from given arguments.
func (l *Log) Errf(format string, args ...interface{}) error {
	e := &Entry{log: l, stack: 1}
	return e.Errf(format, args...)
}

// Output conforms to the OutputLogger interface.
func (l *Log) Output(n int, msg string) error {
	return l.send(n+1, Info, msg, nil)
}

func (l *Log) send(n int, level Level, msg string, e *Entry) (err error) {
	if l.field != nil {
		logger := l.field
		if e != nil && e.fields != nil {
			fields := e.fields
			if v, ok := fields["error"]; ok {
				if err, ok := v.(*erroneous.Erroneous); ok {
					f, o := combineFields(e.fields, Fields(err.Fields()), e.order)
					f.Add("error", err.Message(), &o)
					if file, line := err.Source(); file != "" {
						f.Add("source", fmt.Sprintf("%s:%d", file, line), nil)
					}
				}
			}
			if l.ordered {
				fields["_order"] = e.order
			}

			logger = logger.WithFields(logrus.Fields(fields))
		}

		switch level {
		case Panic:
			logger.Panic(msg)
		case Fatal:
			logger.Fatal(msg)
		case Error:
			logger.Error(msg)
		case Warn:
			logger.Warn(msg)
		case Info:
			logger.Info(msg)
		case Debug:
			logger.Debug(msg)
		}
	}

	if l.out != nil {
		if e != nil && e.fields != nil {
			data, dErr := e.fields.Dump()
			if dErr == nil {
				msg = msg + "  " + string(data)
			} else {
				msg = msg + " <<" + dErr.Error() + ">>"
			}
		}

		oErr := l.out.Output(n+1, msg)
		if oErr != nil {
			err = oErr
		}
	}

	if l.std != nil {
		if e != nil && e.fields != nil {
			data, dErr := e.fields.Dump()
			if dErr == nil {
				msg = msg + "  " + string(data)
			} else {
				msg = msg + " <<" + dErr.Error() + ">>"
			}
		}

		l.std.Print(msg)
	}

	switch level {
	case Panic:
		panic(msg)
	case Fatal:
		os.Exit(1)
	}

	return
}
