package log

import (
	"context"
	"errors"
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
	"github.com/stretchr/testify/assert"
)

func TestWithLogger(t *testing.T) {
	assert := assert.New(t)
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}
	ctx := context.Background()
	ctx = WithLogger(ctx, l)

	assert.Equal(log.NewEntry(l), Entry(ctx))
}

func TestWithEntry(t *testing.T) {
	assert := assert.New(t)
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}
	ctx := context.Background()
	ctx = WithEntry(ctx, log.NewEntry(l))

	assert.Equal(log.NewEntry(l), Entry(ctx))
}

func TestWithFields(t *testing.T) {
	assert := assert.New(t)
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}
	ctx := WithLogger(context.Background(), l)
	f := log.Fields{"foo": true}
	ctx = WithFields(ctx, f)
	Info(ctx, "done")
	assert.Equal(f, h.Entries[0].Fields)
}

func TestWithField(t *testing.T) {
	assert := assert.New(t)
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}
	ctx := WithLogger(context.Background(), l)
	f := log.Fields{"foo": true}
	ctx = WithField(ctx, "foo", true)
	Info(ctx, "done")
	assert.Equal(f, h.Entries[0].Fields)
}

func TestWithError(t *testing.T) {
	assert := assert.New(t)
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}
	ctx := WithLogger(context.Background(), l)
	err := errors.New("an error")
	ctx = WithError(ctx, err)
	Info(ctx, "done")
	assert.Equal("an error", h.Entries[0].Fields["error"])
}

func TestLevels(t *testing.T) {
	assert := assert.New(t)
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.DebugLevel,
	}
	ctx := WithLogger(context.Background(), l)
	Debug(ctx, "Debug")
	Info(ctx, "Info")
	Warn(ctx, "Warn")
	Error(ctx, "Error")
	Debugf(ctx, "Debugf")
	Infof(ctx, "Infof")
	Warnf(ctx, "Warnf")
	Errorf(ctx, "Errorf")
	assert.Equal(log.DebugLevel, h.Entries[0].Level)
	assert.Equal(log.InfoLevel, h.Entries[1].Level)
	assert.Equal(log.WarnLevel, h.Entries[2].Level)
	assert.Equal(log.ErrorLevel, h.Entries[3].Level)
	assert.Equal(log.DebugLevel, h.Entries[4].Level)
	assert.Equal(log.InfoLevel, h.Entries[5].Level)
	assert.Equal(log.WarnLevel, h.Entries[6].Level)
	assert.Equal(log.ErrorLevel, h.Entries[7].Level)
}
