package repository

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/monitoror/monitoror/config"
	ping "github.com/monitoror/monitoror/monitorable/http"
	"github.com/monitoror/monitoror/monitorable/http/models"
)

type (
	httpRepository struct {
		httpClient *http.Client
	}
)

func NewHttpRepository(config *config.Http) ping.Repository {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.SSLVerify},
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(config.Timeout) * time.Millisecond}

	return &httpRepository{client}
}

func (r *httpRepository) Get(url string) (response *models.Response, err error) {
	resp, err := r.httpClient.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response = &models.Response{
		StatusCode: resp.StatusCode,
		Body:       bytes,
	}

	return
}
