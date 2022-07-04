package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/chillyNick/librarySearch/internal/mocks"
	"github.com/chillyNick/librarySearch/internal/models"
	"github.com/chillyNick/librarySearch/pkg/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthors(t *testing.T) {
	type fields struct {
		repo *mocks.MockRepo
	}
	type args struct {
		ctx  context.Context
		book *api.Book
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    *api.Authors
		wantErr assert.ErrorAssertionFunc
	}{
		{"repo.GetAuthors returns error, error expected",
			func(f *fields) {
				f.repo.EXPECT().GetAuthors(gomock.Any(), gomock.Any()).Return(
					nil, fmt.Errorf("dummyErr"),
				).Times(1)
			},
			args{context.Background(), new(api.Book)},
			nil,
			assert.Error,
		},
		{"repo.GetAuthors returns authors, testing func convert models into api.Authors",
			func(f *fields) {
				f.repo.EXPECT().GetAuthors(context.Background(), "BookTitle").Return(
					[]models.Author{{1, "First"}, {2, "Second"}}, nil,
				).Times(1)
			},
			args{context.Background(), &api.Book{Title: "BookTitle"}},
			&api.Authors{Authors: []*api.Author{{Name: "First"}, {Name: "Second"}}},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			f := fields{
				repo: mocks.NewMockRepo(ctrl),
			}
			tt.prepare(&f)

			s := &grpcServer{
				repo: f.repo,
			}

			got, err := s.GetAuthors(tt.args.ctx, tt.args.book)
			if !tt.wantErr(t, err, fmt.Sprintf("GetAuthors(%v, %v)", tt.args.ctx, tt.args.book)) {
				return
			}

			assert.Equalf(t, tt.want, got, "GetAuthors(%v, %v)", tt.args.ctx, tt.args.book)
		})
	}
}

func TestGetBooks(t *testing.T) {
	type fields struct {
		repo *mocks.MockRepo
	}
	type args struct {
		ctx    context.Context
		author *api.Author
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    *api.Books
		wantErr assert.ErrorAssertionFunc
	}{
		{"repo.GetBooks returns error, error expected",
			func(f *fields) {
				f.repo.EXPECT().GetBooks(gomock.Any(), gomock.Any()).Return(
					nil, fmt.Errorf("dummyErr"),
				).Times(1)
			},
			args{context.Background(), new(api.Author)},
			nil,
			assert.Error,
		},
		{"repo.GetBooks returns books, testing func convert models into api.Books",
			func(f *fields) {
				f.repo.EXPECT().GetBooks(context.Background(), "AuthorName").Return(
					[]models.Book{{1, "First"}, {2, "Second"}}, nil,
				).Times(1)
			},
			args{context.Background(), &api.Author{Name: "AuthorName"}},
			&api.Books{Books: []*api.Book{{Title: "First"}, {Title: "Second"}}},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			f := fields{
				repo: mocks.NewMockRepo(ctrl),
			}
			tt.prepare(&f)

			s := &grpcServer{
				repo: f.repo,
			}

			got, err := s.GetBooks(tt.args.ctx, tt.args.author)
			if !tt.wantErr(t, err, fmt.Sprintf("GetBooks(%v, %v)", tt.args.ctx, tt.args.author)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetBooks(%v, %v)", tt.args.ctx, tt.args.author)
		})
	}
}
