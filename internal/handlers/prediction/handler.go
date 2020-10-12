package prediction

import (
	"net/http"
	"github.com/unlar/alp-evaluator/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"bytes"
	"mime/multipart"
	"io"
	"os"
	"io/ioutil"
)

type Handler struct {
	service ports.PredictionService
}

func NewHandler(service ports.PredictionService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Execute(c *gin.Context) {

	mpFile, header, err := c.Request.FormFile("file")

	file, err := os.Create(header.Filename)
	if err != nil {
		logger.Errorf("Hubo un error creando el archivo a guardar", err)
	}
	defer file.Close()

	_, err = io.Copy(file, mpFile)
	if err != nil {
		logger.Errorf("Hubo un error copiando el archivo a guardar", err)
	}

	c.String(http.StatusOK, postFileAndReturnResponse(header.Filename))

	err = os.Remove(header.Filename)
	if err != nil {
		logger.Errorf("Hubo un error borrando el archivo", err)
	}
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
