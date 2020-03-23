package log

import "time"


type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	Fatal() LoggerEvent
	Error() LoggerEvent
	Warn() LoggerEvent
	Info() LoggerEvent
	Debug() LoggerEvent
}

type LoggerEvent interface {
	Trace(ID string) LoggerEvent
	Org(clientID, userID string) LoggerEvent
	Req(ID, IP, host, scheme, method, URL, body string, headers map[string]string) LoggerEvent
	Res(status int, elapsedTime time.Duration, body string, bodyByteLength int, headers map[string]string) LoggerEvent
	Err(err error) LoggerEvent
	ErrWithStack(err error, stacktrace string) LoggerEvent
	Send(message string)
	Sendf(message string, args ...interface{})
}
