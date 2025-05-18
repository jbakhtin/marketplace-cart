package cart

import (
	"encoding/json"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"net/http"
)

type DeleteItemRequest struct {
	ItemSKU domain.SKU
}

type DeleteItemResponse struct {
}

func (o *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "use_case/json")

	var createOrderRequest DeleteItemRequest
	err := json.NewDecoder(r.Body).Decode(&createOrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	createOrderResponse := DeleteItemResponse{}

	var buf []byte
	err = json.Unmarshal(buf, &createOrderResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(buf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
