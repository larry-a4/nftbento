package product

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/larry-a4/nftbento/internal/entities"
)

type Product struct {
	entities.Base
	TotalPrice            int64     `json:"total_price"`
	CollectionCreatedDate time.Time `json:"collection_created_date"`
	ListingTime           time.Time `json:"listing_time"`
	TxTimestamp           time.Time `json:"tx_timestamp"`
	AssetName             string    `json:"asset_name"`
	CollectionImageUrl    string    `json:"collection_image_url"`
	CollectionName        string    `json:"collection_name"`
	CollectionSlug        string    `json:"collection_slug"`
	ContractAddress       string    `json:"contract_address"`
	EventType             string    `json:"event_type"`
	SellerAddress         string    `json:"seller_address"`
	TransactionHash       string    `json:"transaction_hash"`
	WinnerAddress         string    `json:"winner_address"`
}

func InterfaceToModel(data interface{}) (instance *Product, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return instance, err
	}
	return instance, json.Unmarshal(bytes, &instance)
}

func (p *Product) GetFilterId() map[string]interface{} {
	return map[string]interface{}{"_id": p.ID.String()}
}

func (p *Product) TableName() string {
	return "activities"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":                     p.ID.String(),
		"createdAt":               p.CreatedAt.Format(entities.GetTimeFormat()),
		"updatedAt":               p.UpdatedAt.Format(entities.GetTimeFormat()),
		"total_price":             p.TotalPrice,
		"collection_created_date": p.CollectionCreatedDate.Format(entities.GetTimeFormat()),
		"listing_time":            p.ListingTime.Format(entities.GetTimeFormat()),
		"tx_timestamp":            p.TxTimestamp.Format(entities.GetTimeFormat()),
		"asset_name":              p.AssetName,
		"collection_image_url":    p.CollectionImageUrl,
		"collection_name":         p.CollectionName,
		"collection_slug":         p.CollectionSlug,
		"contract_address":        p.ContractAddress,
		"event_type":              p.EventType,
		"seller_address":          p.SellerAddress,
		"transaction_hash":        p.TransactionHash,
		"winner_address":          p.WinnerAddress,
	}
}

func ParseDynamoAttributeToStruct(response map[string]*dynamodb.AttributeValue) (p Product, err error) {
	if response == nil || (response != nil && len(response) == 0) {
		return p, errors.New("Item not found")
	}
	for key, value := range response {
		if key == "_id" {
			p.ID, err = uuid.Parse(*value.S)
			if p.ID == uuid.Nil {
				err = errors.New("Item not found")
			}
		}
		if key == "createdAt" {
			p.CreatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "updatedAt" {
			p.UpdatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "total_price" {
			p.TotalPrice, err = strconv.ParseInt(*value.N, 10, 64)
		}
		if key == "collection_created_date" {
			p.CollectionCreatedDate, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "listing_time" {
			p.ListingTime, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "tx_timestamp" {
			p.TxTimestamp, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "asset_name" {
			p.AssetName = *value.S
		}
		if key == "collection_image_url" {
			p.CollectionImageUrl = *value.S
		}
		if key == "collection_name" {
			p.CollectionName = *value.S
		}
		if key == "collection_slug" {
			p.CollectionSlug = *value.S
		}
		if key == "contract_address" {
			p.ContractAddress = *value.S
		}
		if key == "event_type" {
			p.EventType = *value.S
		}
		if key == "seller_address" {
			p.SellerAddress = *value.S
		}
		if key == "transaction_hash" {
			p.TransactionHash = *value.S
		}
		if key == "winner_address" {
			p.WinnerAddress = *value.S
		}
		if err != nil {
			return p, err
		}
	}

	return p, nil
}
