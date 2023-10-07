package logger

type ErrLogger interface {
	Fatal(err error)
}

type InfoLogger interface {
	Printf(format string, v any)
}

