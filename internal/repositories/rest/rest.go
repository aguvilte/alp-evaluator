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

const BaseUrl = "http://localhost:8088"

type repository struct {
}

func NewProductionRepo() ports.RestRepository {
	return &repository{}
}

func defaultHeaders() http.Header {
	headers := make(http.Header)
	headers.Add("x-caller-scopes", "admin")
	headers.Add("x-admin-id", "fraudMP")
	return headers
}

func (r *repository) GetPrediction(file multipart.File, header *multipart.FileHeader) (*domain.LicensePlate, error) {
	//var asd map[string]interface{}
	//response := r.predictionConfigRequestBuilder.Post("/execute/watson", asd)
	//if response.Err != nil {
	//	return nil, response.Err
	//}
	//if response == nil || response.Response == nil {
	//	// TODO: change error
	//	return nil, response.Err
	//}
	//if response.StatusCode < 200 && response.StatusCode > 299 {
	//	// TODO: change error
	//	return nil, response.Err
	//}
	//body := response.Bytes()
	//var licensePlateModel domain.LicensePlate
	//if err := json.Unmarshal(body, &licensePlateModel); err != nil {
	//	return nil, err
	//}
	//return &licensePlateModel, nil

	logger.Infof("agulog")

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

	req, err := http.NewRequest("POST", "http://localhost:8089/execute/scorpion", &fileDataBuffer)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	req.Header.Set("Content-Type", multipartWritter.FormDataContentType())

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorf("error: ", err)
	}

	defer response.Body.Close()

	logger.Infof("agulog2")

	var licensePlateModel domain.LicensePlate
	data, err := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(data, &licensePlateModel); err != nil {
		logger.Errorf("agulog3", err)
		return nil, err
	}


	logger.Infof("patente: %v", licensePlateModel.Plate)
	logger.Infof("score: %v", licensePlateModel.Score)

	return &licensePlateModel, nil
}

func postFileAndReturnResponse(filename string) string {
	// create a buffer we can write the file to
	fileDataBuffer := bytes.Buffer{}
	multipartWritter := multipart.NewWriter(&fileDataBuffer)
	// open the local file we want to upload
	file, err := os.Open(filename)
	if err != nil {
		logger.Errorf("error: ", err)
	}
	// create an http formfile. This wraps our local file in a format that can be sent to the server
	formFile, err := multipartWritter.CreateFormFile("patente", file.Name())
	if err != nil {
		logger.Errorf("error: ", err)
	}
	// copy the file we want to upload into the form file wrapper
	_, err = io.Copy(formFile, file)
	if err != nil {
		logger.Errorf("error: ", err)
	}
	// close the file writter. This lets it know we're done copying in data
	multipartWritter.Close()
	// create the POST request to send the file data to the server
	req, err := http.NewRequest("POST", "http://localhost:8089/execute/scorpion", &fileDataBuffer)
	if err != nil {
		logger.Errorf("error: ", err)
	}
	// we set the header so the server knows about the files content
	req.Header.Set("Content-Type", multipartWritter.FormDataContentType())
	// send the POST request and recieve the response data
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorf("error: ", err)
	}
	// get data from the response body
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("error: ", err)
	}
	// return the response data
	return string(data)
}
