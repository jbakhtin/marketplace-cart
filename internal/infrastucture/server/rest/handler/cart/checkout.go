package cart

import (
	"encoding/json"
	"net/http"
)

type CheckoutRequest struct{}

type CheckoutResponse struct{}

func (o *Handler) checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "use_case/json")

	var request CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: add logic
	// ...

	response := CheckoutResponse{}

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
