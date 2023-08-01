package utils

func TruncateText(s string, max int) string {
	if len(s) == 0 {
		return ""
	}
	return s[:max]
}
