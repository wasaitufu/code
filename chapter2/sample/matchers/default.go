package matchers

import "code/chapter2/sample/search"

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	search.Register("default", matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	return nil, nil
}
