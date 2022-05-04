package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// adaptErrors will adapt the errors to handle stacktrace information from
// the package github.om/pgk/errors
func adaptErrors(err error, fields []zapcore.Field) []zapcore.Field {

	var zapFields []zapcore.Field
	if err != nil {
		zapFields = make([]zapcore.Field, len(fields)+1)

		// discard all details error information (i.e. stacktrace by errors-pkg)
		zapFields[0] = zap.String("error", err.Error())

		// append the additional fields
		for index := range fields {
			zapFields[index+1] = fields[index]
		}

	} else {
		zapFields = fields
	}

	return zapFields
}
