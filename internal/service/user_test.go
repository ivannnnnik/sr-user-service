package service

import (
	"context"
	"testing"

	"github.com/ivannnnnik/sr-user-service/internal/model"
)

type stubUserRepo struct {
    user *model.User
    err  error
}


func (s *stubUserRepo) Create(_ context.Context, u *model.User) error {
    return s.err
}

func (s *stubUserRepo) GetByID(_ context.Context, _ string) (*model.User, error) {
    return s.user, s.err
}

func TestUserService_Register_Success(t *testing.T) {
	t.Parallel()

    repo := &stubUserRepo{}
    svc := NewUserService(repo)

    user, err := svc.Register(context.Background(), "test@example.com", "testuser", "password123")
    
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if user.Email != "test@example.com" {
        t.Errorf("got email %q, want %q", user.Email, "test@example.com")
    }
}