package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/fomichalopoulos/companiesMicroService/models"
	"gotest.tools/assert"
)

var jwtToken string

func TestEndpoints(t *testing.T) {

	t.Run("LoginNoName", func(t *testing.T) {

		client := &http.Client{}
		url := "http://localhost:8080/company/userLogin"
		contentType := "application/json; charset=UTF-8"

		user := models.User{
			Username: "",
			Password: "5657",
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("Error in marshaling user: %v ", err)
		}

		resp, err := client.Post(url, contentType, bytes.NewBuffer(userJSON))
		if err != nil {
			t.Fatalf("Error when getting JWT: %v", err)
		}

		responseStatus := resp.StatusCode
		statusExpected := http.StatusBadRequest
		assert.Equal(t, responseStatus, statusExpected, "Login code not as expected")

	})

	t.Run("Login", func(t *testing.T) {
		client := &http.Client{}
		url := "http://localhost:8080/company/userLogin"
		contentType := "application/json; charset=UTF-8"

		user := models.User{
			Username: "Fotis",
			Password: "5657",
		}
		userJSON, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("Error in marshaling user: %v ", err)
		}

		resp, err := client.Post(url, contentType, bytes.NewBuffer(userJSON))
		if err != nil {
			t.Fatalf("Error when getting JWT: %v", err)
		}

		responseStatus := resp.StatusCode
		statusExpected := http.StatusOK
		assert.Equal(t, responseStatus, statusExpected, "JWT GET not succesfull")

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Read body error: %v", err)
		}

		jwtToken = string(body)
		fmt.Println("token received is: ", string(body))

	})

	t.Run("CreateCompany", func(t *testing.T) {
		fmt.Println("jwtToken = ", jwtToken)

		url := "http://localhost:8080/company/create"

		amount := 45
		name := "XMAnother"
		registered := true
		typeCompany := "Corporations"

		// company
		aCompany := models.Company{
			Name:        &name,
			Description: "A company",
			Amount:      &amount,
			Registered:  &registered,
			Type:        &typeCompany,
		}
		aCompanyJSON, err := json.Marshal(aCompany)
		if err != nil {
			t.Fatalf("Error in marshalling company: %v", err)
		}

		// Create the POST request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(aCompanyJSON))
		if err != nil {
			t.Fatalf("Error in creating the POST request: %v", err)
		}

		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		// Use the default HTTP client to send the request
		client := &http.Client{
			Timeout: time.Second * 10, // Set a timeout for the request
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusOK
		assert.Equal(t, responseStatus, statusExpected, "Company POST not succesfull")

	})
	t.Run("CreateCompanyNoAmount", func(t *testing.T) {
		fmt.Println("jwtToken = ", jwtToken)

		url := "http://localhost:8080/company/create"

		name := "XMAnother"
		registered := true
		typeCompany := "Corporations"

		// company
		aCompany := models.Company{
			Name:        &name,
			Description: "A company",
			//Amount:      45,
			Registered: &registered,
			Type:       &typeCompany,
		}
		aCompanyJSON, err := json.Marshal(aCompany)
		if err != nil {
			t.Fatalf("Error in marshalling company: %v", err)
		}

		// Create the POST request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(aCompanyJSON))
		if err != nil {
			t.Fatalf("Error in creating the POST request: %v", err)
		}

		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		// Use the default HTTP client to send the request
		client := &http.Client{
			Timeout: time.Second * 10, // Set a timeout for the request
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusBadRequest
		assert.Equal(t, responseStatus, statusExpected, "Company POST not succesfull")

	})
	t.Run("CreateCompanyNameExists", func(t *testing.T) {
		fmt.Println("jwtToken = ", jwtToken)

		url := "http://localhost:8080/company/create"
		amount := 45
		name := "XM"
		registered := true
		typeCompany := "Corporations"

		// company
		aCompany := models.Company{
			Name:        &name,
			Description: "A company",
			Amount:      &amount,
			Registered:  &registered,
			Type:        &typeCompany,
		}
		aCompanyJSON, err := json.Marshal(aCompany)
		if err != nil {
			t.Fatalf("Error in marshalling company: %v", err)
		}

		// Create the POST request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(aCompanyJSON))
		if err != nil {
			t.Fatalf("Error in creating the POST request: %v", err)
		}

		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		// Use the default HTTP client to send the request
		client := &http.Client{
			Timeout: time.Second * 10, // Set a timeout for the request
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusBadRequest
		assert.Equal(t, responseStatus, statusExpected, "Company POST code not expected")

	})
	t.Run("CreateCompanyWrongNameSize", func(t *testing.T) {
		url := "http://localhost:8080/company/create"
		name := "XXM4gfdgdgddgdsdhjsdhsjhdjshdjshdjhsjdhsjhdjshdhsjhdjhsjhM"
		amount := 45
		registered := true
		typeCompany := "Corporations"

		// company
		aCompany := models.Company{
			Name:        &name,
			Description: "A company",
			Amount:      &amount,
			Registered:  &registered,
			Type:        &typeCompany,
		}
		aCompanyJSON, err := json.Marshal(aCompany)
		if err != nil {
			t.Fatalf("Error in marshalling company: %v", err)
		}

		// Create the POST request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(aCompanyJSON))
		if err != nil {
			t.Fatalf("Error in creating the POST request: %v", err)
		}

		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		// Use the default HTTP client to send the request
		client := &http.Client{
			Timeout: time.Second * 10, // Set a timeout for the request
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusBadRequest
		assert.Equal(t, responseStatus, statusExpected, "Company POST code not as expected")
	})
	t.Run("CreateCompanyNoName", func(t *testing.T) {
		url := "http://localhost:8080/company/create"
		amount := 45
		registered := true
		typeCompany := "Corporations"

		// company
		aCompany := models.Company{
			Description: "A company",
			Amount:      &amount,
			Registered:  &registered,
			Type:        &typeCompany,
		}
		aCompanyJSON, err := json.Marshal(aCompany)
		if err != nil {
			t.Fatalf("Error in marshalling company: %v", err)
		}

		// Create the POST request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(aCompanyJSON))
		if err != nil {
			t.Fatalf("Error in creating the POST request: %v", err)
		}

		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		// Use the default HTTP client to send the request
		client := &http.Client{
			Timeout: time.Second * 10, // Set a timeout for the request
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusBadRequest
		assert.Equal(t, responseStatus, statusExpected, "Company POST code not as expected")
	})
	t.Run("CreateCompanyWrongType", func(t *testing.T) {
		url := "http://localhost:8080/company/create"
		name := "XM"
		amount := 45
		registered := true
		typeCompany := "Corporations"

		// company
		aCompany := models.Company{
			Name:        &name,
			Description: "A company",
			Amount:      &amount,
			Registered:  &registered,
			Type:        &typeCompany,
		}
		aCompanyJSON, err := json.Marshal(aCompany)
		if err != nil {
			t.Fatalf("Error in marshalling company: %v", err)
		}

		// Create the POST request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(aCompanyJSON))
		if err != nil {
			t.Fatalf("Error in creating the POST request: %v", err)
		}

		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		// Use the default HTTP client to send the request
		client := &http.Client{
			Timeout: time.Second * 10, // Set a timeout for the request
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusBadRequest
		assert.Equal(t, responseStatus, statusExpected, "Company POST code not as expected")
	})

	t.Run("GetCompany", func(t *testing.T) {

		url := "http://localhost:8080/company?name=XM"
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			t.Fatalf("Error in company GET: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusOK

		assert.Equal(t, responseStatus, statusExpected, "company GET not as expected")

	})

	t.Run("GetCompanyDoesNotExists", func(t *testing.T) {

		url := "http://localhost:8080/company?name=XMsddd"
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			t.Fatalf("Error in company GET: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusNotFound

		assert.Equal(t, responseStatus, statusExpected, "company GET not as expected")

	})

	t.Run("GetCompanyNameExceedsChars", func(t *testing.T) {

		url := "http://localhost:8080/company?name=XMsdddjdskdjskdjksjdksjkdjskjdksjdkjsdkdj"
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			t.Fatalf("Error in company GET: %v", err)
		}
		responseStatus := resp.StatusCode
		statusExpected := http.StatusBadRequest

		assert.Equal(t, responseStatus, statusExpected, "company GET not as expected")

	})

	t.Run("DeleteCompany", func(t *testing.T) {
		url := "http://localhost:8080/company?name=XMAnother"
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			t.Fatalf("Error in creating the Delete request: %v", err)
		}
		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}

		responseStatus := resp.StatusCode
		statusExpected := http.StatusOK

		assert.Equal(t, responseStatus, statusExpected, "company DEL status not as expected")

	})

	t.Run("DeleteCompanyThatDoesNotExist", func(t *testing.T) {
		url := "http://localhost:8080/company?name=XMAnother"
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			t.Fatalf("Error in creating the Delete request: %v", err)
		}
		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}

		responseStatus := resp.StatusCode
		statusExpected := http.StatusNotFound

		assert.Equal(t, responseStatus, statusExpected, "company DEL status not as expected")

	})

	t.Run("DeleteCompanyNameExceedLength", func(t *testing.T) {
		url := "http://localhost:8080/company?name=XMAnotherjsdkdjskjdksjdkjskdjksjkdjskjdksjddsdk"
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			t.Fatalf("Error in creating the Delete request: %v", err)
		}
		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending POST request: %v", err)
		}

		responseStatus := resp.StatusCode
		statusExpected := http.StatusBadRequest

		assert.Equal(t, responseStatus, statusExpected, "company DEL status not as expected")

	})

	t.Run("PatchCompany", func(t *testing.T) {
		url := "http://localhost:8080/company?name=XM"

		name := "XMUpdated"
		amount := 34
		registered := true
		typeCompany := "Corporations"

		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		company := models.Company{
			Name:       &name,
			Amount:     &amount,
			Registered: &registered,
			Type:       &typeCompany,
		}
		companyJSON, err := json.Marshal(company)
		if err != nil {
			t.Fatalf("Error in marshalling company: %v", err)
		}

		req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(companyJSON))
		// Set the request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken) // Add JWT token to Authorization header

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error in sending PATCH request: %v", err)
		}

		responseStatus := resp.StatusCode
		statusExpected := http.StatusOK

		assert.Equal(t, responseStatus, statusExpected, "company DEL status not as expected")

	})

}
