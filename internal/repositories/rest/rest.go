package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/unlar/alp-evaluator/internal/core/domain"

	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
	"github.com/unlar/alp-evaluator/internal/core/ports"
)

const BaseUrl = "http://localhost:8088"

type repository struct {
	predictionConfigRequestBuilder *rest.RequestBuilder
}

func NewProductionRepo() ports.Repository {
	predictionConfigPool := &rest.CustomPool{
		MaxIdleConnsPerHost: 100,
	}
	return &repository{
		predictionConfigRequestBuilder: &rest.RequestBuilder{
					Headers:        defaultHeaders(),
					Timeout:        2000 * time.Millisecond,
					BaseURL:        BaseUrl,
					ContentType:    rest.JSON,
					EnableCache:    false,
					DisableTimeout: false,
					CustomPool:     predictionConfigPool,
				},
	}
}

func defaultHeaders() http.Header {
	headers := make(http.Header)
	headers.Add("x-caller-scopes", "admin")
	headers.Add("x-admin-id", "fraudMP")
	return headers
}

func (r *repository) GetPrediction(img []uint8) (*domain.LicensePlate, error) {
	var asd map[string]interface{}
	response := r.predictionConfigRequestBuilder.Post("/execute/watson", asd)
	if response.Err != nil {
		return nil, response.Err
	}
	if response == nil || response.Response == nil {
		// TODO: change error
		return nil, response.Err
	}
	if response.StatusCode < 200 && response.StatusCode > 299 {
		// TODO: change error
		return nil, response.Err
	}
	body := response.Bytes()
	var licensePlateModel domain.LicensePlate
	if err := json.Unmarshal(body, &licensePlateModel); err != nil {
		return nil, err
	}
	return &licensePlateModel, nil
}
