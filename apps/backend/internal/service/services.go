package service

import (
	"fmt"

	"github.com/SR-SHREYAS/Go-Custom-Boilerplate/internal/lib/aws"
	"github.com/SR-SHREYAS/Go-Custom-Boilerplate/internal/lib/job"
	"github.com/SR-SHREYAS/Go-Custom-Boilerplate/internal/repository"
	"github.com/SR-SHREYAS/Go-Custom-Boilerplate/internal/server"
)

type Services struct {
	Auth     *AuthService
	Job      *job.JobService
	Todo     *TodoService
	Category *CategoryService
	Comment  *CommentService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	awsClient, err := aws.NewAWS(s)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS client: %w", err)
	}

	return &Services{
		Job:      s.Job,
		Auth:     authService,
		Category: NewCategoryService(s, repos.Category),
		Comment:  NewCommentService(s, repos.Comment, repos.Todo),
		Todo:     NewTodoService(s, repos.Todo, repos.Category, awsClient),
	}, nil
}
