package commenTed

import (
	"log"
	"regexp"
	"strings"
)

const (
	comment     = `(\/\/[\s]+`
	startRegexp = comment + `c:remove)`
	endRegexp   = comment + `?)c:end` // ([^\n]+)
	lineRegexp  = comment + `)`
	tooRegexp   = comment + `c:too)`
)

var startExp = regexp.MustCompile(startRegexp)
var exp = regexp.MustCompile(startRegexp + `([\s\S]+?)` + endRegexp)
var lineExp = regexp.MustCompile(lineRegexp)
var tooExp = regexp.MustCompile(tooRegexp)
var endExp = regexp.MustCompile(endRegexp)

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
