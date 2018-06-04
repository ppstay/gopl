func basename2(s string) string {
	slash := string.lastIndex(s, "/")
	s = s[slash+1:]
	if dot := string.lastIndex(s, "."); dot >=0 {
		s = s[:dot]
	}
	return s
}