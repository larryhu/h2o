package h2o

import (
	"fmt"
	"regexp"
	"strconv"
)

// seconds = time.Since(detal).Seconds()
// seconds to day, hour, minute, secs
func TimeFormat(seconds int) (day, hour, minute, secs int) {
	day = seconds / (60 * 60 * 24)
	seconds = seconds % (60 * 60 * 24)

	hour = seconds / (60 * 60)
	seconds = seconds % (60 * 60)

	minute = seconds / 60
	secs = seconds % 60
	return
}

// seconds = time.Since(detal).Seconds()
// seconds => xdxhxmxs
func TimeFormatStyle(seconds int) string {
	day, hour, minute, secs := TimeFormat(seconds)
	if day+hour+minute == 0 {
		return fmt.Sprintf("%ds", seconds)
	}
	if day+hour == 0 {
		return fmt.Sprintf("%dm%ds", minute, secs)
	}
	if day == 0 {
		return fmt.Sprintf("%dh%dm%ds", hour, minute, secs)
	}
	return fmt.Sprintf("%dd%dh%dm%ds", day, hour, minute, secs)
}

// human like: 1m 10s 2h 2d
// return seconds, error
func TimeParse(human string) (seconds int, err error) {
	p := regexp.MustCompile(`([0-9]+[s|m|h|d])`)

	if !p.MatchString(human) {
		return 0, fmt.Errorf("ParseTime error : %s", human)
	}
	args := p.FindAllString(human, -1)
	for _, s := range args {
		_num, err := strconv.ParseInt(s[:len(s)-1], 0, 0)
		if err != nil {
			return 0, err
		}
		num := int(_num)
		switch s[len(s)-1] {
		case 'd':
			seconds += num * 3600 * 24
		case 'h':
			seconds += num * 3600
		case 'm':
			seconds += num * 60
		case 's':
			seconds += num
		}
	}
	return seconds, nil
}
