package errorUtils

import (
	"fmt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/httperrors/contracts"
	defaultLogger "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger/defaultlogger"
	"runtime/debug"
	"strings"

	"emperror.dev/errors"
)

// CheckErrMessages check for specific messages contains in the error
func CheckErrMessages(err error, messages ...string) bool {
	for _, message := range messages {
		if strings.Contains(
			strings.TrimSpace(strings.ToLower(err.Error())),
			strings.TrimSpace(strings.ToLower(message)),
		) {
			return true
		}
	}
	return false
}

// ErrorsWithStack returns a string contains errors messages in the stack with its stack trace levels for given error
func ErrorsWithStack(err error) string {
	res := fmt.Sprintf("%+v\n", err)
	return res
}

// ErrorsWithoutStack just returns error messages without its callstack
func ErrorsWithoutStack(err error, format bool) string {
	res := fmt.Sprintf("%v\n", err)

	if format {
		var errStr string
		items := strings.Split(res, ":")
		for _, item := range items {
			errStr += fmt.Sprintf("%s\n", strings.TrimSpace(item))
		}
		return errStr
	}

	return res
}

// StackTrace returns all stack traces with a string contains just stack trace levels for the given error
func StackTrace(err error) string {
	var stackTrace contracts.StackTracer
	stackStr := ""
	for {
		s, ok := err.(contracts.StackTracer)
		stackTrace = s
		if ok {
			stackStr += fmt.Sprintf("%+v\n", stackTrace.StackTrace())

			if !ok {
				break
			}
		}
		err = errors.Unwrap(err)
		if err == nil {
			break
		}
	}

	return stackStr
}

// RootStackTrace returns root stack trace with a string contains just stack trace levels for the given error
func RootStackTrace(err error) string {
	var stackTrace contracts.StackTracer
	stackStr := ""
	for {
		s, ok := err.(contracts.StackTracer)
		stackTrace = s
		if ok {
			stackStr = fmt.Sprintf("%+v\n", stackTrace.StackTrace())

			if !ok {
				break
			}
		}
		err = errors.Unwrap(err)
		if err == nil {
			break
		}
	}

	return stackStr
}

func HandlePanic() {
	if r := recover(); r != nil {
		defaultLogger.GetLogger().
			Error("stacktrace from panic: \n" + string(debug.Stack()))
	}
}
