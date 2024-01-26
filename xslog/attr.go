package xslog

import "log/slog"

var KeyErr = "err"

// Err returns an Attr for an error value. By default, it uses the key "err",
// however one can easily monkey-patch the behavior by changing the value of xslog.KeyErr variable.
func Err(e error) slog.Attr {
	return slog.Attr{
		Key:   KeyErr,
		Value: slog.StringValue(e.Error()),
	}
}
