package parser

import (
	"regexp"
	"wordcounter/utils"
)

func ExtractWords(content utils.Content) []string {
	text := content.Title + " " + content.Heading + " " + content.Description
	re := regexp.MustCompile(`[a-zA-Z]+`)
	return re.FindAllString(text, -1)
}
