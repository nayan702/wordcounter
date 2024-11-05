package test

import (
	"testing"
	"wordcounter/handler"
	"github.com/stretchr/testify/assert"
)

func TestFetchURL(t *testing.T) {
	url := "https://www.engadget.com/2019/08/25/sony-and-yamaha-sc-1-sociable-cart/"

	content, err := handler.FetchURL(url)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if content.Title == "" {
		t.Error("expected a title, got empty")
	}
	assert.Equal(t, "Sony and Yamaha are making a self-driving cart for theme parks", content.Title)
	assert.Equal(t, "It's smarter, more comfortable and longer-lasting than its ancestors.", content.Heading)
	assert.Equal(t, "Remember how we said Sony's self-driving SC-1 concept would make for a great party bus? Apparently, Sony had the same idea. The company is partnering with Yamaha on the SC-1 Sociable Cart, an expansion of the concept designed for entertainment purposes like theme parks, golf courses and \"commercial facilities.\" The new version seats five people instead of three (and in greater comfort), lasts longer through replaceable batteries and uses additional image sensors to improve its situational awareness.\nAs before, Sony feels the sensors eliminate the need for windows. A 49-inch 4K monitor on the inside provides a mixed reality view of the world, while four 55-inch 4K displays bombard passers-by with ads and other material. It will even use AI to optimize promos for outside people based on factors like age and gender -- not quite Minority Report levels of eerily accurate ad targeting, but getting there.\nThe two companies expect to use the Sociable Cart for services in Japan sometime in fiscal 2019 (that is, before the end of March 2020). It won't, however, be available for sale. Not that you'd really want one given its glacial 11.8MPH top speed. This is strictly for fun on closed circuits, not your next pub crawl.\n", content.Description)
}
