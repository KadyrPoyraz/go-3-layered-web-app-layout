package helloworld

import "context"

type Service interface {
    SayHello(ctx context.Context) ([]string, error)
}
