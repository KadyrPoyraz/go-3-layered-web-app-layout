package helloworld

import (
	"context"
	"fmt"

	"github.com/KadyrPoyraz/httplayout/internal/repository/helloworld"
)

type service struct {
    helloRepo helloworld.Repo
}

func New(
    helloRepo helloworld.Repo,
) Service {
	return &service{
        helloRepo: helloRepo,
    }
}

func (s *service) SayHello(ctx context.Context) ([]string, error) {
    text, err := s.helloRepo.GetHelloWorld(ctx)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    return text, nil
}

