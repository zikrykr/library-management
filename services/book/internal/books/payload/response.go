package payload

type (
	GetBookResponse struct {
		Title         string       `json:"title"`
		Description   string       `json:"description"`
		ISBN          string       `json:"isbn"`
		Author        BookAuthor   `json:"author"`
		Category      BookCategory `json:"category"`
		PublishedYear int          `json:"published_year"`
	}

	BookAuthor struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}

	BookCategory struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)
