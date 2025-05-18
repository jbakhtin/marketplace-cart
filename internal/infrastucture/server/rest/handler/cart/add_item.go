package cart

import (
	"encoding/json"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"net/http"
)

type AddItemRequest struct {
	Item  domain.ItemSKU
	Count domain.ItemCount
}

type AddItemResponse struct{}

func (o *Handler) Cancel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "use_case/json")

	var request CancelRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	response := CancelResponse{}

	var buf []byte
	err = json.Unmarshal(buf, &response)
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
