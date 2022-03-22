package web

import (
	"net/http"

	"github.com/blue-health/go-backend-exercise/app/loyalty"
	"github.com/go-chi/chi"
)

// Campaign web service handles issuing of the coupons.
type CampaignService struct {
	campaigns []loyalty.Campaign
}

// Router return the HTTP router for the
// REST API exposed by the campaign web service.
func (s *CampaignService) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/coupons/{code}", func(r chi.Router) {
		r.Post("/apply", s.applyCoupon)
	})

	return r
}

// applyCoupon attempts to apply the coupon identified by the coupon code found in the URL.
// If the coupon was applied successfully, we return HTTP 200 and the new basket.
// Otherise we return an application code 400-499, like HTTP 404 if coupon was not found.
// You can assume the shopping basket is in the JSON request body.
func (s *CampaignService) applyCoupon(w http.ResponseWriter, r *http.Request) {}
