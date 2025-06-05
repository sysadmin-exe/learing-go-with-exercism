package parsinglogfiles

import "regexp"
import "fmt"
func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(\~|\*|\=|\-)*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
    count := 0
    
    for _, l := range lines {
        if re.MatchString(l) {
            count ++
        }
    }
    
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(User\s+)+([A-Za-z]+\d+)`)
    fmt.Println(lines)
	tagged := []string{}
    
    for _, l := range lines {
        fmt.Println(re.MatchString(l))
    	fmt.Println(l)
        if re.MatchString(l) {
            tagged = append(tagged, "[USR] " + re.FindStringSubmatch(l)[2] + " " + l)
        } else {
        	tagged = append(tagged, l)
        }
    }
    return tagged
}
