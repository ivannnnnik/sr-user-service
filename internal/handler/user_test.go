package handler_test

import (
	"context"
	"testing"
	"time"

	userv1 "github.com/ivannnnnik/sr-proto/gen/go/user/v1"
	"github.com/ivannnnnik/sr-user-service/internal/handler"
	"github.com/ivannnnnik/sr-user-service/internal/model"
)

type stubUserService struct {
    user *model.User
    err  error
}

func (s *stubUserService) Register(_ context.Context, _, _, _ string) (*model.User, error) {
    return s.user, s.err
}

func (s *stubUserService) GetProfile(_ context.Context, _ string) (*model.User, error) {
    return s.user, s.err
}

func TestUserHandler_Register_Success(t *testing.T) {
    t.Parallel()

    fixedUser := &model.User{
        ID:        "u1",
        Email:     "alice@example.com",
        Username:  "alice",
        CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
    }

    h := handler.NewUserHandler(&stubUserService{user: fixedUser})

    resp, err := h.Register(context.Background(), &userv1.RegisterRequest{
        Email:    "alice@example.com",
        Username: "alice",
        Password: "secret123",
    })
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if resp.User.Id != "u1" {
        t.Errorf("got id %q, want %q", resp.User.Id, "u1")
    }
    if resp.User.Email != "alice@example.com" {
        t.Errorf("got email %q, want %q", resp.User.Email, "alice@example.com")
    }
}