package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"regexp"
	"strconv"
)

var fileName string
var show string
var resolution string

func main() {
	flag.StringVarP(&fileName, "file", "f", "", "File name")
	flag.StringVarP(&show, "show", "s", "", "Show name")
	flag.StringVarP(&resolution, "resolution", "r", "1080p", "Resolution; default is 1080p")
	flag.Parse()

	show = "Bizarre Foods with Andrew Zimmern"
	//fileName := "107 - Trinidad and Tobago.avi"
	re := regexp.MustCompile(`(\d)(\d\d) - (.*)\.(.*)`)
	m := re.FindStringSubmatch(fileName)
	if m != nil && len(m) == 5 {
		season := m[1]
		i, _ := strconv.Atoi(m[2])
		episode := i + 1
		title := m[3]
		fType := m[4]
		fmt.Printf("%s - S0%sE%02d - %s.%s\n", show, season, episode, title, fType)
	}

	//re := regexp.MustCompile(`(S\d\dE\d\d)\.?(\S*)\.1080p`)
	//m := re.FindStringSubmatch(fileName)
	//if m != nil && len(m) == 3 && m[2] != "" {
	//	episode := m[1]
	//	title := m[2]
	//	title = strings.Replace(title, ".", " ", -1)
	//	fmt.Printf("%s - %s - %s.mkv\n", show, episode, title)
	//} else {
	//	episode := m[1]
	//	fmt.Printf("%s - %s.mkv\n", show, episode)
	//}
}
