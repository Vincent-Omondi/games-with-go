package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func story(page *storyPage) {
	if page != nil {
		fmt.Println(page.text)
		story(page.nextPage)
	}
}

//this function allow users to add a page to the end of another page
func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

//This fub=nctionallows the user to add a page between pages
func (page *storyPage) addAfter(text string) {
	page.nextPage = &storyPage{text, page.nextPage}
}

func main() {

	page1 := storyPage{"Every morning when I wake up, first this is to thankful for being present at the moment.", nil}
	page2 := storyPage{"I do morning morning medition to connect to my inner self", nil}
	page3 := storyPage{"Then I do morning pages, to unwind my mind and to be cogitive of my thoughts", nil}
	page4 := storyPage{"In the end of the day I do movement to exercise my body and to freshen up.", nil}
	page4.addToEnd("Good day")
	page4.addAfter("I pray.")

	page1.nextPage = &page2
	page2.nextPage = &page3
	page3.nextPage = &page4

	story(&page1)
}
