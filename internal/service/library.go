package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "homework-week4/api/libraray_system/v1"
	"homework-week4/internal/biz"
)

// LibraryService is a greeter service.
type LibraryService struct {
	v1.UnimplementedLibraryServer

	uc  *biz.LibraryUseCase
	log *log.Helper
}

// NewLibraryService new a greeter service.
func NewLibraryService(uc *biz.LibraryUseCase, logger log.Logger) *LibraryService {
	return &LibraryService{uc: uc, log: log.NewHelper("service/library", logger)}
}

// QueryBook implements helloworld.Library
func (s *LibraryService) QueryBook(ctx context.Context, in *v1.QueryBookRequest) (*v1.Book, error) {
	s.log.Infof("Query Book Received: %v", in.GetName())
	if in.GetName() == "error" {
		return nil, errors.NotFound("library", v1.ErrorReason_USER_NOT_FOUND.String(), in.GetName())
	}
	b, err := s.uc.QueryBook(ctx, in.GetName())
	if err != nil {
		return nil, err
	}
	res := &v1.Book{
		Name:   b.Name,
		Isbn:   b.Isbn,
		Author: b.Author,
		BookId: b.Id,
	}
	return res, nil
}

func (s *LibraryService) BorrowBook(ctx context.Context, in *v1.BorrowBookRequest) (*v1.BorrowBookResponse, error) {
	s.log.Infof("Borrow Book Received: %v", in.GetBookId())
	if in.GetBookId() == 0 {
		return nil, errors.BadRequest("library", "", "")
	}
	req := &biz.BorrowReq{
		BookId:     in.GetBookId(),
		BorrowDate: in.BorrowDate,
		MemberId:   in.Borrower.MemberId,
	}
	returnDate, err := s.uc.BorrowBook(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.BorrowBookResponse{
		Status:     "success",
		ReturnDate: returnDate,
	}, nil
}

func (s *LibraryService) ReturnBook(ctx context.Context, in *v1.ReturnBookRequest) (*v1.ReturnBookResponse, error) {
	s.log.Infof("Return Book Received: %v", in.GetIsbn())
	if in.GetIsbn() == "" {
		return nil, errors.BadRequest("library", "", "")
	}
	_, err := s.uc.ReturnBook(ctx, in.GetIsbn(), in.GetBorrower().MemberId)
	if err != nil {
		return nil, err
	}
	return &v1.ReturnBookResponse{Status: "success"}, nil
}
