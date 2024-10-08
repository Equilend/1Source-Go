// Package api provides functions for HTTP verb access to 1Source REST API.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/EquiLend/1Source-Go/models"
)

// PostProposeLoan will perform an HTTP POST operation
// against the 1Source REST API to propose a loan
// https://www.kirandev.com/http-post-golang
func PostProposeLoan(apiEndPoint string, bearer string, body []byte) (string, error) {
	ctx := context.Background()
	transport := &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			r.Header.Set("Authorization", bearer)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			return nil, nil
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(15) * time.Second,
	}

	request, err := http.NewRequestWithContext(ctx, "POST", apiEndPoint, bytes.NewBuffer(body))
	request.Header.Set("Authorization", bearer)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Println("Error creating new HTTP Request: ", err)
		return "", err
	}

	log.Println("Calling API endpoint: ", apiEndPoint)
	resp, err := client.Do(request)

	// Close response body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal("Error in HTTP POST API call: ", err)
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading HTTP POST response: ", err)
		return "", err
	}

	if resp.StatusCode != http.StatusCreated {
		log.Println("Error proposing loan. HTTP Response Status:", resp.Status)
	} else {
		var cir models.LoanInitiationResponse

		err := json.Unmarshal(respBody, &cir)
		if err != nil {
			return "", err
		}

		return cir.Message, err
	}

	return "", err
}

// PostCancelLoan will perform an HTTP POST operation
// against the 1Source REST API to cancel a loan
// https://www.kirandev.com/http-post-golang
func PostCancelLoan(apiEndPoint string, bearer string) (string, error) {
	ctx := context.Background()
	transport := &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			r.Header.Set("Authorization", bearer)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			return nil, nil
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(15) * time.Second,
	}

	request, err := http.NewRequestWithContext(ctx, "POST", apiEndPoint, nil)
	request.Header.Set("Authorization", bearer)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Println("Error creating new HTTP Request: ", err)
		return "", err
	}

	log.Println("Calling API endpoint: ", apiEndPoint)
	resp, err := client.Do(request)

	// Close response body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal("Error in HTTP POST API call: ", err)
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading HTTP POST response: ", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error canceling proposed loan. HTTP Response Status:", resp.Status)
	} else {
		var ccr models.LoanCancelReponse

		err := json.Unmarshal(respBody, &ccr)
		if err != nil {
			return "", err
		}

		return ccr.Message, err
	}

	return "", err
}

// PostDeclineLoan will perform an HTTP POST operation
// against the 1Source REST API to decline a loan
// https://www.kirandev.com/http-post-golang
func PostDeclineLoan(apiEndPoint string, bearer string) (string, error) {
	ctx := context.Background()
	transport := &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			r.Header.Set("Authorization", bearer)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			return nil, nil
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(15) * time.Second,
	}

	request, err := http.NewRequestWithContext(ctx, "POST", apiEndPoint, nil)
	request.Header.Set("Authorization", bearer)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Println("Error creating new HTTP Request: ", err)
		return "", err
	}

	log.Println("Calling API endpoint: ", apiEndPoint)
	resp, err := client.Do(request)

	// Close response body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal("Error in HTTP POST API call: ", err)
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading HTTP POST response: ", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error declining proposed loan. HTTP Response Status:", resp.Status)
	} else {
		var cdr models.LoanDeclineReponse

		err := json.Unmarshal(respBody, &cdr)
		if err != nil {
			return "", err
		}

		return cdr.Message, err
	}

	return "", err
}
