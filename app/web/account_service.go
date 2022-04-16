package web

import (
	"github.com/blue-health/go-backend-exercise/app/account"
	"github.com/go-chi/chi"
)

type AccountService struct {
	accountService account.Service
}

func NewAccountService(accountService account.Service) *AccountService {
	return &AccountService{accountService: accountService}
}

func (s *AccountService) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/my_route", nil)

	return r
}
