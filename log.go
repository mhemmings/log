package log

import (
	"context"

	"github.com/apex/log"
)

// entryKey holds the context key used for loggers.
type entryKey struct{}

// WithLogger returns a new context derived from ctx that is associated with the given logger.
func WithLogger(ctx context.Context, logger *log.Logger) context.Context {
	return WithEntry(ctx, log.NewEntry(logger))
}

// WithEntry returns a new context derived from ctx that is associated with the given entry.
func WithEntry(ctx context.Context, entry *log.Entry) context.Context {
	return context.WithValue(ctx, entryKey{}, entry)
}

// WithFields returns a new context derived from ctx that has an entry that always logs the given fields.
func WithFields(ctx context.Context, fields log.Fielder) context.Context {
	return WithEntry(ctx, Entry(ctx).WithFields(fields))
}

// WithField returns a new context derived from ctx that has an entry that always logs the given key and value.
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithEntry(ctx, Entry(ctx).WithField(key, value))
}

// WithError returns a new context derived from ctx that has an entry that always logs the given error.
func WithError(ctx context.Context, err error) context.Context {
	return WithEntry(ctx, Entry(ctx).WithError(err))
}

// Entry returns the entry associated with the given context. If there is no Entry, it will return the default log.Log.
func Entry(ctx context.Context) log.Interface {
	if ctx == nil {
		panic("nil context passed to Entry")
	}
	if entry, _ := ctx.Value(entryKey{}).(*log.Entry); entry != nil {
		return entry
	}
	return log.Log
}

// Debug level message.
func Debug(ctx context.Context, msg string) {
	Entry(ctx).Debug(msg)
}

// Info level message.
func Info(ctx context.Context, msg string) {
	Entry(ctx).Info(msg)
}

// Warn level message.
func Warn(ctx context.Context, msg string) {
	Entry(ctx).Warn(msg)
}

// Error level message.
func Error(ctx context.Context, msg string) {
	Entry(ctx).Error(msg)
}

// Fatal level message, followed by an exit.
func Fatal(ctx context.Context, msg string) {
	Entry(ctx).Fatal(msg)
}

// Debugf level formatted message.
func Debugf(ctx context.Context, msg string, v ...interface{}) {
	Entry(ctx).Debugf(msg, v...)
}

// Infof level formatted message.
func Infof(ctx context.Context, msg string, v ...interface{}) {
	Entry(ctx).Infof(msg, v...)
}

// Warnf level formatted message.
func Warnf(ctx context.Context, msg string, v ...interface{}) {
	Entry(ctx).Warnf(msg, v...)
}

// Errorf level formatted message.
func Errorf(ctx context.Context, msg string, v ...interface{}) {
	Entry(ctx).Errorf(msg, v...)
}

// Fatalf level formatted message, followed by an exit.
func Fatalf(ctx context.Context, msg string, v ...interface{}) {
	Entry(ctx).Fatalf(msg, v...)
}

// Trace returns a new entry with a Stop method to fire off
// a corresponding completion log, useful with defer.
func Trace(ctx context.Context, msg string) *log.Entry {
	return Entry(ctx).Trace(msg)
}
