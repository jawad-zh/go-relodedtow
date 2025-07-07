package goreloaded

func IsControlStr(s string) bool {
	return s == "\n" || s == "\t" || s == "\r"
}
