package booking

import (
	"encoding/json"
	"net/http"

	"github.com/sarastee/application-design-test/internal/model"
	"github.com/sarastee/application-design-test/internal/utils/logger"
)

func (i *Implementation) CreateOrder(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.LogErrorf(err.Error())
		}
	}()

	var newOrder model.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "incorrect request format", http.StatusBadRequest)
		logger.LogErrorf("error while decoding message: %v", err)
		return
	}

	if err := i.bookingService.CreateOrder(newOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogErrorf("error while creating order: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newOrder); err != nil {
		logger.LogErrorf("error while encoding message: %v", err)
	}

	logger.LogInfo("booking created successfully: %v", newOrder)

}
