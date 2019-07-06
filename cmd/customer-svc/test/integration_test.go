package test

import (
	"bytes"
	"customer-app/cmd/customer-svc/handler"
	"customer-app/cmd/customer-svc/pkg"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAPI(t *testing.T) {
	app := pkg.InmemApp()
	ts := httptest.NewServer(handler.Handle(app))
	defer ts.Close()

	requestBody, _ := ioutil.ReadFile("create-request.json")
	buffer := bytes.NewBuffer(requestBody)
	res, err := http.Post(ts.URL+"/customers/", "application/json", buffer)
	if err != nil {
		log.Fatal(err)
	}

	if res.Status != "200 OK" {
		t.Errorf("Status was %s expecting 200", res.Status)
		return
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}
