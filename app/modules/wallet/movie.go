package wallet 

import(
	"time"
)

type movie struct {
	title string
	length time.Duration 
}

func NewMovie(title string, length time.Duration) *movie {
	return &movie{title: title, length: length}
}

func (m *movie) GetTitle() string {
	return m.title
}

func (m *movie) GetLenght() time.Duration {
	return m.length
}