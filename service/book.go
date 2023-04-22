package service

import (
	"Challange-7/model"
)

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBooks() (res []model.Book, err error)
	GetBookById(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	// call repository
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil
}
func (s *Service) GetBooks() (res []model.Book, err error) {
	// return s.repo.GetBooks()
	res, err = s.repo.GetBooks()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetBookById(id int64) (res model.Book, err error) {
	res, err = s.repo.GetBookById(id)
	if err != nil {
		return res, err
	}
	// res.Author = "bayu"

	return res, nil
	// return s.repo.GetBookById(id)
}

func (s *Service) UpdateBook(in model.Book) (res model.Book, err error) {
	// return s.repo.UpdateBook(in)
	res, err = s.repo.UpdateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) DeleteBook(id int64) (err error) {
	// return s.repo.DeleteBook(id)
	err = s.repo.DeleteBook(id)
	if err != nil {
		return err
	}

	return nil
}
