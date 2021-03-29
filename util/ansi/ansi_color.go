package ansi

const (
	BoldBlack   = "\033[30;1m"
	BoldRed     = "\033[31;1m"
	BoldGreen   = "\033[32;1m"
	BoldYellow  = "\033[33;1m"
	BoldBlue    = "\033[34;1m"
	BoldMagenta = "\033[35;1m"
	BoldCyan    = "\033[36;1m"
	BoldWhite   = "\033[37;1m"
	YELLOW      = "\033[0;33m"
	CLEAR       = "\033[0K"
)

const (
	ERROR = BoldRed
	WARN  = BoldRed
	INFO  = BoldGreen
	DEBUG = BoldYellow
	RESET = CLEAR
)
