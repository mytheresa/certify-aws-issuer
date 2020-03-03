package logger

import (
	"github.com/sirupsen/logrus"
)

type logrusMapper struct{
	logrus *logrus.Entry
}

// Trace proxy to logrus.Trace
func (cl logrusMapper) Trace(msg string, fields ...map[string]interface{}) {
	cl.logrus.Trace(msg, fields)
}

// Debug proxy to logrus.Debug
func (cl logrusMapper) Debug(msg string, fields ...map[string]interface{}) {
	cl.logrus.Debug(msg, fields)
}

// Info proxy to logrus.Info
func (cl logrusMapper) Info(msg string, fields ...map[string]interface{}) {
	cl.logrus.Info(msg, fields)
}

// Warn proxy to logrus.Warn
func (cl logrusMapper) Warn(msg string, fields ...map[string]interface{}) {
	cl.logrus.Warn(msg, fields)
}

// Warn proxy to logrus.Warn
func (cl logrusMapper) Error(msg string, fields ...map[string]interface{}) {
	cl.logrus.Error(msg, fields)
}

// NewCertifyLogrusLogger create new instance of certiry logrus logger mapper
func NewCertifyLogrusMapper(logrus *logrus.Entry) *logrusMapper {
	return &logrusMapper{logrus}
}
