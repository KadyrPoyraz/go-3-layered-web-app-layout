package helloworld

import "context"

type Repo interface {
    GetHelloWorld(ctx context.Context) ([]string, error)
}
