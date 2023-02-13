package src

import "regexp"

type conversion struct{
	re *regexp.Regexp
	replace interface{}
}

// ParseTextToMarkdown takes the text string and returns markdown content
func ParseTextToMarkdown(str string)string{
	listOfConversion := []conversion{
		// for bold markup conversion
		{
			re: regexp.MustCompile(`\*(\S[^*]*)\*`),
			replace: "**$1**",
		},
		// for italics markup conversion
		{
			re: regexp.MustCompile(`\b\_(\S[^_]*)\_`),
			replace: "*$1*",
		},
	}

	for _ , conversionContent := range listOfConversion{
		switch v := conversionContent.replace.(type){
		case string:
			str = conversionContent.re.ReplaceAllString(str, v)
		}
	}

	return str
}
