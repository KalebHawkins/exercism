package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	if !re.Match([]byte(text)) {
		return false
	}

	return true
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-=~\*]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*(password).*"`)
	var count int

	for _, l := range lines {
		if re.Match([]byte(l)) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile((`\sUser\s+(\S+)`))
	var res []string

	for _, l := range lines {
		m := re.FindStringSubmatch(l)
		if len(m) > 1 {
			l = fmt.Sprintf("[USR] %s %s", m[1], l)
		}
		res = append(res, l)
	}

	return res
}
