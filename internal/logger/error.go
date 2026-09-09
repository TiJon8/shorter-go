package logger

import "log/slog"


func ErrorAttr(key string, value error) slog.Attr {
	return slog.Attr{Key: key, Value: slog.StringValue(value.Error())}
}