# log

[![Documentation](https://godoc.org/github.com/mhemmings/log?status.svg)](http://godoc.org/github.com/mhemmings/log)

A super-lightweight wrapper around [apex/log](https://github.com/apex/log) to make it context aware. Inspired by [juju/zaputil/zapctx](https://github.com/juju/zaputil).

## Usage

All the methods in (`log.Interface`)[https://godoc.org/github.com/apex/log#Interface] are defined, with a `context.Context`
as the first parameter.

```go
func WithFields(ctx context.Context, fields log.Fielder) context.Context
func WithField(ctx context.Context, key string, value interface{}) context.Context
func WithError(ctx context.Context, err error) context.Context
func Entry(ctx context.Context) log.Interface
func Debug(ctx context.Context, msg string)
func Info(ctx context.Context, msg string)
func Warn(ctx context.Context, msg string)
func Error(ctx context.Context, msg string)
func Fatal(ctx context.Context, msg string)
func Debugf(ctx context.Context, msg string, v ...interface{})
func Infof(ctx context.Context, msg string, v ...interface{})
func Warnf(ctx context.Context, msg string, v ...interface{})
func Errorf(ctx context.Context, msg string, v ...interface{})
func Fatalf(ctx context.Context, msg string, v ...interface{})
func Trace(ctx context.Context, msg string) *log.Entry
```

A common usecase is adding default fields to HTTP logs:

```go
package main

import (
  "net/http"

  apex "github.com/apex/log"
  "github.com/mhemmings/log"
)

func middleware(next http.HandlerFunc) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    f := apex.Fields{
      "method": r.Method,
      "url":    r.URL.Path,
    }
    ctx := log.WithFields(r.Context(), f)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

func handler(w http.ResponseWriter, r *http.Request) {
  log.Info(r.Context(), "handling request") // logs message with fields "method" and "url"
  w.Write([]byte("Hello world"))
}

func main() {
  http.Handle("/", middleware(handler))
  http.ListenAndServe(":8080", nil)
}
```
