package rest

import (
	"encoding/json"
	"net/http"

	"github.com/unlar/alp-evaluator/internal/core/domain"

	"github.com/unlar/alp-evaluator/internal/core/ports"
	"mime/multipart"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"io"
	"io/ioutil"
	"bytes"
	"os"
)

const ScorpionModelUrl = "http://localhost:8089/execute/scorpion"
const WatsonModelUrl = "http://localhost:8088/execute/watson"

type repository struct {
}

func NewProductionRepo() ports.RestRepository {
	return &repository{}
}

func (r *repository) GetPrediction(file multipart.File, header *multipart.FileHeader) (*domain.LicensePlate, error) {
	fileDataBuffer := bytes.Buffer{}
	multipartWritter := multipart.NewWriter(&fileDataBuffer)

	file, err := os.Open(header.Filename)
	if err != nil {
		logger.Errorf("Error", err)
	}

	formFile, err := multipartWritter.CreateFormFile("patente", header.Filename)
	if err != nil {
		logger.Errorf("Error", err)
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	multipartWritter.Close()

	firstModel, _ := postToScorpion(fileDataBuffer, multipartWritter)
	secondModel, _ := postToWatson(fileDataBuffer, multipartWritter)

	logger.Infof("firstModel: %v - secondModel: %v", firstModel.Score, secondModel.Score)

	return firstModel, err
}

func postToScorpion(fileDataBuffer bytes.Buffer, multipartWritter *multipart.Writer) (*domain.LicensePlate, error) {
	req, err := http.NewRequest("POST", ScorpionModelUrl, &fileDataBuffer)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	req.Header.Set("Content-Type", multipartWritter.FormDataContentType())

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	return mapToLicensePlate(data)
}

func postToWatson(fileDataBuffer bytes.Buffer, multipartWritter *multipart.Writer) (*domain.LicensePlate, error) {
	req, err := http.NewRequest("POST", WatsonModelUrl, &fileDataBuffer)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	req.Header.Set("Content-Type", multipartWritter.FormDataContentType())

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	return mapToLicensePlate(data)
}

func mapToLicensePlate(data []byte) (*domain.LicensePlate, error) {
	var licensePlateModel domain.LicensePlate
	if err := json.Unmarshal(data, &licensePlateModel); err != nil {
		return nil, err
	}

	return &licensePlateModel, nil
}