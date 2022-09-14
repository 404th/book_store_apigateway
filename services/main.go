package services

import (
	"fmt"

	"github.com/404th/book_store_apigateway/config"
	"github.com/404th/book_store_apigateway/genproto/book_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	BookService() book_service.BookServiceClient
}

type grpcClients struct {
	bookService book_service.BookServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	// ...1
	connBookService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BookServiceHost, conf.BookServicePort), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		bookService: book_service.NewBookServiceClient(connBookService),
	}, nil
}

// ...1
func (g *grpcClients) BookService() book_service.BookServiceClient {
	return g.bookService
}
