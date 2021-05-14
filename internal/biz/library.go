package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Book struct {
	Id int64
	Name string
	Isbn string
	Author string
}

type BorrowReq struct {
	BookId int64
	BorrowDate int64
	MemberId int64
}

type LibraryRepo interface {
	QueryBook(context.Context, string) (*Book, error)
	BorrowBook(context.Context, *BorrowReq) (int64, error)
	ReturnBook(context.Context, string, int64) (bool, error)
}

type LibraryUseCase struct {
	repo LibraryRepo
	log  *log.Helper
}

func NewLibraryUseCase(repo LibraryRepo, logger log.Logger) *LibraryUseCase {
	return &LibraryUseCase{repo: repo, log: log.NewHelper("usecase/library", logger)}
}

func (uc *LibraryUseCase) QueryBook(ctx context.Context, bookName string) (*Book, error) {
	return uc.repo.QueryBook(ctx, bookName)
}

func (uc *LibraryUseCase) BorrowBook(ctx context.Context, req *BorrowReq) (int64, error) {
	return uc.repo.BorrowBook(ctx, req)
}

func (uc *LibraryUseCase) ReturnBook(ctx context.Context, isbn string, memberID int64) (bool, error) {
	return uc.repo.ReturnBook(ctx, isbn, memberID)
}