package api

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var relativeDayRe = regexp.MustCompile(`^-(\d+)d$`)

func ParseDate(input string) (string, error) {
	switch input {
	case "today":
		return time.Now().Format("2006-01-02"), nil
	case "yesterday":
		return time.Now().AddDate(0, 0, -1).Format("2006-01-02"), nil
	}

	if m := relativeDayRe.FindStringSubmatch(input); m != nil {
		days, _ := strconv.Atoi(m[1])
		return time.Now().AddDate(0, 0, -days).Format("2006-01-02"), nil
	}

	_, err := time.Parse("2006-01-02", input)
	if err != nil {
		return "", fmt.Errorf("invalid date %q: use YYYY-MM-DD, 'today', 'yesterday', or '-Nd' (e.g. '-7d')", input)
	}
	return input, nil
}
