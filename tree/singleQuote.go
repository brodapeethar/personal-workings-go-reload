package tree

import "regexp"


func SingleQuote(s string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	result := re.ReplaceAllString(s, `'$1'`)
	return result
}




func Punctuate(s string) string {
	puncWord := regexp.MustCompile(`\s+([.,!?;:]+)\s*`)
	s = puncWord.ReplaceAllString(s, `$1 `)
	return s
}