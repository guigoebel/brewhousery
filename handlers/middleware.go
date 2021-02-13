package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		product := &data.Product{}

		err := data.FromJSON(product, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading product ...", http.StatusBadRequest)
			return
		}

		errs := p.v.Validate(product)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)
			http.Error(w, fmt.Sprintf("Error validating product: %s", errs), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, *product)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
