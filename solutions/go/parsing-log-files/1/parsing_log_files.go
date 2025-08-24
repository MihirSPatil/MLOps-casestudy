package parsinglogfiles
import ("regexp"
       )
func IsValidLine(text string) bool {
    re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    if re.FindString(text) != ""{
        return true
    }else{return false}
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<(~|\*|=|-|)+\>`)
    loglines := re.Split(text, -1)
    return loglines
}

func CountQuotedPasswords(lines []string) int {
    re := regexp.MustCompile(`(?i)".*password.*"`)
    count := 0
    for _, line := range lines {
        if re.MatchString(line) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line\d+)`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re, _ := regexp.Compile(`User\s+(\S+)`)
    var result []string
    for _,line := range lines{
        matches := re.FindStringSubmatch(line)
        if len(matches) > 1 {
			result = append(result, "[USR] "+matches[1]+" "+line)
		} else {
			result = append(result, line)
		}
    }
    return result
}
