// Package models contains the type structures related to 1source-go
package models

type LoanInitiationResponse struct {
	Timestamp string `json:"timestamp"`
	Status    uint32 `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

type LoanCancelReponse struct {
	Timestamp string `json:"timestamp"`
	Status    uint32 `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

type LoanDeclineReponse struct {
	Timestamp string `json:"timestamp"`
	Status    uint32 `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}
