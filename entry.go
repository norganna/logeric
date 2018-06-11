package logeric

import (
	"fmt"

	"github.com/norganna/erroneous"
)

// Entry represents a log entry.
type Entry struct {
	log    *Log
	stack  int
	fields Fields
	order  []string
}

var _ FieldLogger = (*Entry)(nil)

// WithError adds an error to the entry.
func (e *Entry) WithError(err error) *Entry {
	return e.WithFields(Fields{"error": err})
}

// WithField adds a single field to the entry.
func (e *Entry) WithField(key string, value interface{}) *Entry {
	return e.WithFields(Fields{key: value})
}

// WithFields adds a bunch of fields to the entry.
func (e *Entry) WithFields(fields Fields) *Entry {
	f, o := combineFields(e.fields, fields, e.order)

	return &Entry{
		log:    e.log,
		fields: f,
		order:  o,
	}
}

// Debugf is an implementation of FieldLogger.
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.log.send(2+e.stack, Debug, fmt.Sprintf(format, args...), e)
}

// Infof is an implementation of FieldLogger.
func (e *Entry) Infof(format string, args ...interface{}) {
	e.log.send(2+e.stack, Info, fmt.Sprintf(format, args...), e)
}

// Printf is an implementation of FieldLogger.
func (e *Entry) Printf(format string, args ...interface{}) {
	e.log.send(2+e.stack, Info, fmt.Sprintf(format, args...), e)
}

// Warnf is an implementation of FieldLogger.
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.log.send(2+e.stack, Warn, fmt.Sprintf(format, args...), e)
}

// Warningf is an implementation of FieldLogger.
func (e *Entry) Warningf(format string, args ...interface{}) {
	e.log.send(2+e.stack, Warn, fmt.Sprintf(format, args...), e)
}

// Errorf is an implementation of FieldLogger.
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.log.send(2+e.stack, Error, fmt.Sprintf(format, args...), e)
}

// Fatalf is an implementation of FieldLogger.
func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.log.send(2+e.stack, Fatal, fmt.Sprintf(format, args...), e)
}

// Panicf is an implementation of FieldLogger.
func (e *Entry) Panicf(format string, args ...interface{}) {
	e.log.send(2+e.stack, Panic, fmt.Sprintf(format, args...), e)
}

// Debug is an implementation of FieldLogger.
func (e *Entry) Debug(args ...interface{}) {
	e.log.send(2+e.stack, Debug, fmt.Sprint(args...), e)
}

// Info is an implementation of FieldLogger.
func (e *Entry) Info(args ...interface{}) {
	e.log.send(2+e.stack, Info, fmt.Sprint(args...), e)
}

// Print is an implementation of FieldLogger.
func (e *Entry) Print(args ...interface{}) {
	e.log.send(2+e.stack, Info, fmt.Sprint(args...), e)
}

// Warn is an implementation of FieldLogger.
func (e *Entry) Warn(args ...interface{}) {
	e.log.send(2+e.stack, Warn, fmt.Sprint(args...), e)
}

// Warning is an implementation of FieldLogger.
func (e *Entry) Warning(args ...interface{}) {
	e.log.send(2+e.stack, Warn, fmt.Sprint(args...), e)
}

// Error is an implementation of FieldLogger.
func (e *Entry) Error(args ...interface{}) {
	e.log.send(2+e.stack, Error, fmt.Sprint(args...), e)
}

// Fatal is an implementation of FieldLogger.
func (e *Entry) Fatal(args ...interface{}) {
	e.log.send(2+e.stack, Fatal, fmt.Sprint(args...), e)
}

// Panic is an implementation of FieldLogger.
func (e *Entry) Panic(args ...interface{}) {
	e.log.send(2+e.stack, Panic, fmt.Sprint(args...), e)
}

// Debugln is an implementation of FieldLogger.
func (e *Entry) Debugln(args ...interface{}) {
	e.log.send(2+e.stack, Debug, fmt.Sprint(args...), e)
}

// Infoln is an implementation of FieldLogger.
func (e *Entry) Infoln(args ...interface{}) {
	e.log.send(2+e.stack, Info, fmt.Sprint(args...), e)
}

// Println is an implementation of FieldLogger.
func (e *Entry) Println(args ...interface{}) {
	e.log.send(2+e.stack, Info, fmt.Sprint(args...), e)
}

// Warnln is an implementation of FieldLogger.
func (e *Entry) Warnln(args ...interface{}) {
	e.log.send(2+e.stack, Warn, fmt.Sprint(args...), e)
}

// Warningln is an implementation of FieldLogger.
func (e *Entry) Warningln(args ...interface{}) {
	e.log.send(2+e.stack, Warn, fmt.Sprint(args...), e)
}

// Errorln is an implementation of FieldLogger.
func (e *Entry) Errorln(args ...interface{}) {
	e.log.send(2+e.stack, Error, fmt.Sprint(args...), e)
}

// Fatalln is an implementation of FieldLogger.
func (e *Entry) Fatalln(args ...interface{}) {
	e.log.send(2+e.stack, Fatal, fmt.Sprint(args...), e)
}

// Panicln is an implementation of FieldLogger.
func (e *Entry) Panicln(args ...interface{}) {
	e.log.send(2+e.stack, Panic, fmt.Sprint(args...), e)
}

// Err returns a concatenated error from given arguments.
func (e *Entry) Err(args ...interface{}) error {
	return erroneous.New(
		erroneous.Depth(2+e.stack),
		erroneous.Msg(fmt.Sprint(args...), erroneous.ErrFields(e.fields)),
	)
}

// Errf returns a formatted error from given arguments.
func (e *Entry) Errf(format string, args ...interface{}) error {
	return erroneous.New(
		erroneous.Depth(2+e.stack),
		erroneous.Msg(fmt.Sprintf(format, args...), erroneous.ErrFields(e.fields)),
	)
}
