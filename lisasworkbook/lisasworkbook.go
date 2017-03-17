package main

import "fmt"

type Page struct {
	number   int
	problems map[int]struct{}
}

func (p Page) hasSpecialProblem() bool {
	hasSpecial := false

	for problemNumber, _ := range p.problems {
		if problemNumber == p.number {
			hasSpecial = true
			break
		}
	}

	return hasSpecial
}

func (p Page) String() string {
	return fmt.Sprintf("{Page number: %v, problems: %v}", p.number, p.problems)
}

type Chapter struct {
	number int
	pages  []Page
}

func (c Chapter) String() string {
	return fmt.Sprintf("{Chapter number: %v, pages: %v}", c.number, c.pages)
}

func main() {
	var numChapters int
	var maxProblemsPerPage int

	fmt.Scanf("%v %v\n", &numChapters, &maxProblemsPerPage)

	var numProblemsInChapter int
	currentPageNumber := 1

	chapters := []Chapter{}

	for chapter := 1; chapter <= numChapters; chapter++ {
		fmt.Scanf("%v", &numProblemsInChapter)

		c := convertToChapter(chapter, numProblemsInChapter, currentPageNumber, maxProblemsPerPage)

		chapters = append(chapters, c)

		currentPageNumber += len(c.pages)
	}

	var numSpecialProblems int

	for _, chapter := range chapters {
		for _, page := range chapter.pages {
			if page.hasSpecialProblem() {
				numSpecialProblems++
			}
		}
	}

	fmt.Println(numSpecialProblems)
}

func convertToChapter(chapterNumber, numProblemsInChapter, currentPageNumber, maxProblemsPerPage int) Chapter {
	chapter := Chapter{number: chapterNumber, pages: []Page{}}

	problemNumber := 1

	page := Page{number: currentPageNumber, problems: map[int]struct{}{}}

	pageFull := false

	for problemNumber <= numProblemsInChapter {
		page.problems[problemNumber] = struct{}{}
		pageFull = problemNumber%maxProblemsPerPage == 0

		if pageFull {
			chapter.pages = append(chapter.pages, page)

			currentPageNumber++
			page = Page{number: currentPageNumber, problems: map[int]struct{}{}}
		}

		problemNumber++
	}

	if !pageFull {
		chapter.pages = append(chapter.pages, page)
	}

	return chapter
}
