package cart

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest/handler/response"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"net/http"
)

type DeleteItemRequest struct {
	ItemSKU domain.SKU
}

func (o *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var request DeleteItemRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		response.WriteStandardResponse(w, r, http.StatusBadRequest, nil, err)
		return
	}

	// TODO: add logic
	// ...

	response.WriteStandardResponse(w, r, http.StatusOK, nil, nil)
}
