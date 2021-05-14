package data

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"homework-week4/internal/biz"
	"time"
)

type libraryRepo struct {
	data *Data
	log  *log.Helper
}

// NewLibraryRepo .
func NewLibraryRepo(data *Data, logger log.Logger) biz.LibraryRepo {
	return &libraryRepo{
		data: data,
		log:  log.NewHelper("data/library", logger),
	}
}

func (l *libraryRepo) QueryBook(ctx context.Context, s string) (*biz.Book, error) {
	b := &Book{}
	query := "SELECT id, name, isbn, author FROM book WHERE name=?"
	err := l.data.db.QueryRowContext(ctx, query, s).Scan(&b.Id, &b.Name, &b.Isbn, &b.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("library", "book name not found", "")
		}
		return nil, err
	}
	return &biz.Book{
		Id:     b.Id,
		Name:   b.Name,
		Isbn:   b.Isbn,
		Author: b.Author,
	}, nil
}

func (l *libraryRepo) BorrowBook(ctx context.Context, req *biz.BorrowReq) (int64, error) {
	s := `INSERT INTO ops (member_id, book_id, borrow_date) VALUES (?, ?, ?)`
	st , err := l.data.db.Prepare(s)
	if err != nil {
		return 0, err
	}
	_, err = st.Exec(req.MemberId, req.BookId, req.BorrowDate)
	if err != nil {
		return 0, err
	}
	expireDate := time.Unix(req.BorrowDate,0).Add(time.Hour * 24 * 7).Unix()
	return expireDate, nil
}

func (l *libraryRepo) ReturnBook(ctx context.Context, isbn string, memberID int64) (bool, error) {
	var id int64
	query := "SELECT id FROM book WHERE isbn=?"
	err := l.data.db.QueryRowContext(ctx, query, isbn).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.NotFound("library", "book isbn not found", "")
		}
		return false, err
	}

	s := `UPDATE ops SET return_date=? WHERE book_id=? AND member_id=? `
	st , err := l.data.db.Prepare(s)
	if err != nil {
		return false, err
	}
	_, err = st.Exec(time.Now().Unix(),id, memberID)
	if err != nil {
		return false, err
	}
	return true, nil
}
