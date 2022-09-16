package gocdg

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type BillService service

type BillData struct {
	Bills []Bill `json:"bills"`
}

type Bill struct {
	Congress     int `json:"congress"`
	LatestAction struct {
		ActionDate string `json:"actionDate"`
		Text       string `json:"text"`
	} `json:"latestAction"`
	Number                  string    `json:"number"`
	OriginChamber           string    `json:"originChamber"`
	OriginChamberCode       string    `json:"originChamberCode"`
	Title                   string    `json:"title"`
	Type                    string    `json:"type"`
	UpdateDate              string    `json:"updateDate"`
	UpdateDateIncludingText time.Time `json:"updateDateIncludingText"`
	URL                     string    `json:"url"`
}

func (s *BillService) ListAll(ctx context.Context) (*BillData, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/bill", s.BaseURL), nil)
	if err != nil {
		return &BillData{}, err
	}

	req = req.WithContext(ctx)
	billData := new(BillData)

	err = s.newRequest(req, billData)
	if err != nil {
		return &BillData{}, err
	}

	return billData, nil
}
