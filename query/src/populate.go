package src

import "fmt"

func populateTutorial() []Tutorial {
	var tutorials []Tutorial
	for i := 0; i < 5; i++ {
		author := &Author{Name: fmt.Sprintf("Author %+v", i), Tutorials: []int{1}}
		tutorial := Tutorial{
			ID:     int64(i),
			Title:  fmt.Sprintf("Go GraphQL Tutorial %+v", i),
			Author: *author,
			Comments: []Comment{
				Comment{Body: "First Comment"},
			},
		}
		tutorials = append(tutorials, tutorial)
	}

	return tutorials
}
