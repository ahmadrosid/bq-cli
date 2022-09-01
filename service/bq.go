package service

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type bqService struct {
	bq *bigquery.Client
}

type BiqueryService interface {
	GetDataFromBQ(query string, ctx context.Context) ([]string, error)
}

func NewBiqueryService(bq *bigquery.Client) BiqueryService {
	return &bqService{
		bq: bq,
	}
}

func (r *bqService) GetDataFromBQ(query string, ctx context.Context) ([]string, error) {
	q := r.bq.Query(query)
	q.DisableQueryCache = true
	row, err := q.Read(ctx)

	if err != nil {
		return nil, err
	}
	var totalData []string

	for {
		var plat map[string]bigquery.Value
		err = row.Next(&plat)
		if err != nil {
			if err == iterator.Done {
				break
			} else {
				return nil, err
			}
		}
		val, err := json.Marshal(plat)
		if err != nil {
			return nil, err
		}

		totalData = append(totalData, string(val))
	}
	return totalData, nil
}
