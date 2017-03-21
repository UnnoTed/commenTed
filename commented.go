package commenTed

import (
	"log"
	"regexp"
	"strings"
)

const (
	comment         = `(\/\/[\s]+`
	startRegexp     = comment + `c:remove)`
	endRegexp       = comment + `?)c:end` // ([^\n]+)
	lineRegexp      = comment + `)`
	tooRegexp       = comment + `c:too)`
	replaceUpRegexp = comment + `c:replace:up[\s]+)`
)

var startExp = regexp.MustCompile(startRegexp)
var exp = regexp.MustCompile(startRegexp + `([\s\S]+?)` + endRegexp)
var lineExp = regexp.MustCompile(lineRegexp)
var tooExp = regexp.MustCompile(tooRegexp)
var endExp = regexp.MustCompile(endRegexp)
var replaceUpExp = regexp.MustCompile(replaceUpRegexp)

// Parse ya txt
func Parse(data []byte, debug bool) []byte {
	src := string(data)

	src = exp.ReplaceAllStringFunc(src, func(s string) string {
		lines := strings.Split(s, "\n")
		for i := range lines {

			if startExp.MatchString(lines[i]) || tooExp.MatchString(lines[i]) || endExp.MatchString(lines[i]) {
				if debug {
					log.Println("REMOVED", i, lines[i])
				}

				lines[i] = ""
			} else {
				if debug {
					log.Println("UNCOMMENTED", i, lines[i])
				}

				lines[i] = lineExp.ReplaceAllString(lines[i], "")
			}
		}

		j := strings.Join(lines, "\n")
		if debug {
			log.Println("REPLACED", j)
		}

		return j
	})

	return []byte(src)
}

// ParseReplace ya txt
func ParseReplace(data []byte, start string, end string, debug bool) []byte {
	lines := strings.Split(string(data), "\n")

	var lastLine string
	for i := range lines {
		if replaceUpExp.MatchString(lines[i]) {
			replaces := replaceUpExp.ReplaceAllString(lines[i], "")
			list := strings.Split(replaces, " - ")

			for _, r := range list {
				r = strings.TrimSpace(r)
				r = strings.TrimPrefix(r, start)
				r = strings.TrimSuffix(r, end)
				data := strings.Split(r, "|")
				log.Println("data", r)

				lastLine = strings.Replace(lastLine, data[0], data[1], -1)
				lines[i-1] = lastLine
			}

			lines[i] = ""
		}

		lastLine = lines[i]
	}

	return []byte(strings.Join(lines, "\n"))
}
