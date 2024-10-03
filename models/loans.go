// Package models contains the models for the application
package models

type (
	Loan struct {
		LoanId             string `json:"loanId"`
		LastEventId        uint32 `json:"lastEventId"`
		LoanStatus         string `json:"loanStatus"`
		SettlementStatus   string `json:"settlementStatus"`
		LastUpdatePartyId  string `json:"lastUpdatePartyId"`
		LastUpdateDateTime string `json:"lastUpdateDateTime"`
		Trade              trade
		Settlement         settlement
	}

	trade struct {
		ExecutionVenue     executionvenue
		Instrument         instrument
		Rate               rate    `json:"rate"`
		Quantity           uint32  `json:"quantity"`
		BillingCurrency    string  `json:"billingCurrency"`
		DividendRatePct    float32 `json:"dividendRatePct"`
		TradeDate          string  `json:"tradeDate"`
		SettlementType     string  `json:"settlementType"`
		Collateral         collateral
		TransactingParties []transactingparties
	}

	executionvenue struct {
		VenueType    string `json:"type"`
		Platform     platform
		VenueParties []venueparties
	}

	venueparties struct {
		PartyRole string `json:"partyRole"`
	}

	platform struct {
		GleifLei   string `json:"gliefLei"`
		LegalName  string `json:"legalName"`
		VenueName  string `json:"venueName"`
		VenueRefId string `json:"venueRefId"`
	}

	instrument struct {
		Ticker      string `json:"ticker"`
		Cusip       string `json:"cusip"`
		Isin        string `json:"isin"`
		Sedol       string `json:"sedol"`
		Figi        string `json:"figi"`
		Description string `json:"description"`
	}

	rate struct {
		Rebate rebate
	}

	rebate struct {
		Fixed fixed
	}

	fixed struct {
		BaseRate      float32 `json:"baseRate"`
		EffectiveDate string  `json:"effectiveDate"`
		EffectiveRate float32 `json:"effectiveRate"`
	}

	collateral struct {
		LoanValue       float64 `json:"loanValue"`
		CollateralValue float64 `json:"collateralValue"`
		Currency        string  `json:"currency"`
		RoundingRule    uint32  `json:"roundingRule"`
		RoundingMode    string  `json:"roundingMode"`
		Margin          uint32  `json:"margin"`
	}

	transactingparties struct {
		PartyRole string `json:"partyRole"`
		Party     party
	}

	party struct {
		PartyId         string `json:"partyId"`
		PartyName       string `json:"partyName"`
		GleifLei        string `json:"gleifLei"`
		InternalPartyId string `json:"internalPartyId"`
	}

	settlement struct {
		PartyRole   string `json:"partyRole"`
		Instruction instruction
	}

	instruction struct {
		SettlementBic     string `json:"SettlementBic"`
		LocalAgentBic     string `json:"localAgentBic"`
		LocalAgentName    string `json:"localAgentName"`
		LocalAgentAcct    string `json:"localAgentAcct"`
		LocalMarketFields localmarketfields
	}

	localmarketfields struct {
		LocalFieldName  string `json:"localFieldName"`
		LocalFieldValue string `json:"localFieldValue"`
	}
)
