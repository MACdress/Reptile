package utils

import "regexp"

func ExtractString (contents []byte,re  *regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	result := ""
	if match!=nil {
		result = string(match[1])
	}
	return result
}
