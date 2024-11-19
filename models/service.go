package models

type Service interface {
	// Start the service
	GetBooks() ([]Book, error)
	GetBookByID(id int) (Book, error)
	CreateBook(book BookRequest) (Book, error)
	UpdateBook(id int, book Book) (Book, error)
	DeleteBook(id int) error
}

type service struct {
	// Service implementation
	repository Repository
}

func NewService(repository Repository) Service {
	// Initialize the service
	return &service{repository}
}

// GetBooks implements Service.
func (s *service) GetBooks() ([]Book, error) {
	return s.repository.GetBooks()
}

// GetBookByID implements Service.
func (s *service) GetBookByID(id int) (Book, error) {
	return s.repository.GetBookByID(id)
}

// UpdateBook implements Service.
func (s *service) UpdateBook(id int, book Book) (Book, error) {
	return s.repository.UpdateBook(id, book)
}

// CreateBook implements Service.
func (s *service) CreateBook(BookRequest BookRequest) (Book, error) {
	book := Book{
		Title:  BookRequest.Title,
		Author: BookRequest.Author,
		Price:  BookRequest.Price,
	}
	return s.repository.CreateBook(book)
} //karena disini beda dengan di repo untuk requestnya , disini minta bookreq sedangkan di repo minta book, jadi harus diubah dulu ke books

// DeleteBook implements Service.
func (s *service) DeleteBook(id int) error {
	panic("unimplemented")
}
