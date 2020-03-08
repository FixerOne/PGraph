package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, email string, firtsName string, lastName string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "firtsName", firtsName, "lastName", lastName, "e0", e0)
	}()
	return l.next.Create(ctx, email, firtsName, lastName)
}
