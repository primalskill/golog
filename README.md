# Log handler for the Go [log/slog](https://pkg.go.dev/log/slog) package

## Overview

There are three handlers implemented in this package with a heavy focus on the development handler.

1. Development - logging in development mode.
2. Discard - all gets discarded, useful for integration testing where log lines can become noise.
3. Production - a wrapper around slog's JSON handler with sensible defaults.


## Usage - Development

This handler should only be used in development. It formats the log lines where the attribute keys are color coded.

It is encouraged to always set the default logger with `slog.SetDefault(...)` in order to not "depend" directly on this logger in your code base.

### Example

```go
devLogger := golog.NewDevelopment()
slog.SetDefault(devLogger)

slog.Info("this log will be pretty printed")
```

### Output

Example code used to generate the output: https://go.dev/play/p/Y0d-i5_SutP

![Pasted_Image_21_04_2024__14_39](https://github.com/primalskill/golog/assets/489775/92d76a6f-78b1-42a8-95b7-1cc17c92c3d5)

### JSON Output

If the data passed in as a `string` or `[]byte` is a valid JSON object, it will be automatically formatted, colorized, and its type marked as JSON.

![image](https://github.com/primalskill/golog/assets/489775/231e5257-bf65-4791-b1d4-0c9ca19060cb)

-----------

## Usage - Discard

All log lines will be discarded. Should be used for integration tests.

### Example

```go
discardLogger := golog.NewDiscard()
slog.SetDefault(discardLogger)

slog.Info("this log will be discarded")
```

------------

## Usage - Production

The production handler is just a wrapper for [slog's JSON handler](https://pkg.go.dev/log/slog#NewJSONHandler) with a few sensible defaults:

- Logs are output to `stderr` to align it to the [POSIX standard](https://linux.die.net/man/3/stderr)
- The default log level is `slog.LevelInfo`
- slog's TimeKey is replaced in JSON to a property named `t`
- slog's source is enabled and formatted to `<file:line>`, e.g. `my_file.go:17`

### Example

```go
prodLogger := golog.NewProduction()
slog.SetDefault(prodLogger)

slog.Info("this log will be output to stderr and formatted as JSON")
```

