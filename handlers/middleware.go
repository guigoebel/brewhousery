package handlers

import (
	"context"
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Println("[INFO] Middleware validation handler function")
		w.Header().Add("Content-Type", "application/json")

		product := &data.Product{}
		p.l.Println("[DEBUG] Retrieved product data from URL")

		err := data.FromJSON(product, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			p.l.Println("[ERROR] deserializing product", err)
			return
		}

		errs := p.v.Validate(product)
		if len(errs) != 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			p.l.Println("[ERROR] validating product", errs)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, *product)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
