package service

import (
	"Challange-7/model"
	"Challange-7/repository/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_GetBookByID(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTabel []testCase

	testTabel = append(testTabel, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(model.Book{
				ID:       1,
				NameBook: "Into Deep...",
				Author:   "ayu",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:       1,
			NameBook: "Into Deep...",
			Author:   "ayu",
		},
	})

	testTabel = append(testTabel, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	testTabel = append(testTabel, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTabel {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBookById(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			NameBook: "Into Deep...",
			Author:   "ayu",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{
				ID:       1,
				NameBook: "Into Deep...",
				Author:   "ayu",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:       1,
			NameBook: "Into Deep...",
			Author:   "ayu",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Book{
			NameBook: "Into Deep...",
			Author:   "ayu",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid book name length",
		wantError: true,
		input: model.Book{

			NameBook: "Int", // case negative
			Author:   "adi",
		},
		expectedError: errors.New("invalid book name length"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("invalid book name length")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_UpdateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTabel []testCase

	testTabel = append(testTabel, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			ID:       1,
			NameBook: "Into Deep...",
			Author:   "ayu",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{
				ID:       1,
				NameBook: "Into Deep...",
				Author:   "ayu",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:       1,
			NameBook: "Into Deep...",
			Author:   "ayu",
		},
	})

	testTabel = append(testTabel, testCase{
		name:      "record not found",
		wantError: true,
		input: model.Book{
			ID:       1,
			NameBook: "Into Deep...",
			Author:   "ayu",
		},
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	testTabel = append(testTabel, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Book{
			ID:       1,
			NameBook: "Into Deep...",
			Author:   "ayu",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTabel {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.UpdateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_DeleteBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		id             int64
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTabel []testCase

	testTabel = append(testTabel, testCase{
		name:      "success",
		wantError: false,
		id:        1,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(nil).Times(1)
		},
	})

	testTabel = append(testTabel, testCase{
		name:          "record not found",
		wantError:     true,
		id:            1,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTabel {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			err := service.DeleteBook(testCase.id)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}

func Test_BookService_GetBooks(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult []model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTabel []testCase

	testTabel = append(testTabel, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBooks().Return([]model.Book{}, nil).Times(1)
		},
		expectedResult: []model.Book{},
	})

	testTabel = append(testTabel, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBooks().Return([]model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTabel {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBooks()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}
