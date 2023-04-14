package moody

import logging "github.com/ipfs/go-log/v2"

type Object map[string]interface{}

type Formater func(Object) string

type Logger interface {
	Error(Object) error
	Errorf(format string, args ...interface{}) error
	Debug(Object) error
	Warn(Object) error
	Info(Object) error
	Render(Object) string
	Std() logging.StandardLogger

	Sub(string) Logger
}
