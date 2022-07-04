package server

import (
	"context"

	pb "github.com/chillyNick/librarySearch/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GetBooks returns books by author name from db
func (s *grpcServer) GetBooks(ctx context.Context, author *pb.Author) (*pb.Books, error) {
	books, err := s.repo.GetBooks(ctx, author.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Some internal error")
	}

	pbBooks := new(pb.Books)
	for _, b := range books {
		pbBooks.Books = append(pbBooks.Books, &pb.Book{Title: b.Title})
	}

	return pbBooks, nil
}

//GetAuthors returns authors entities by book title from db
func (s *grpcServer) GetAuthors(ctx context.Context, book *pb.Book) (*pb.Authors, error) {
	authors, err := s.repo.GetAuthors(ctx, book.GetTitle())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Some internal error")
	}

	pbAuthors := new(pb.Authors)
	for _, a := range authors {
		pbAuthors.Authors = append(pbAuthors.Authors, &pb.Author{Name: a.Name})
	}

	return pbAuthors, nil
}
