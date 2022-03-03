package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	UUID        uuid.UUID `json:"order_uid"`
	TrackNumber string    `json:"track_number"`
	Entry       string    `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  uuid.UUID `json:"transaction"`
		RequestId    string    `json:"request_id"`
		Currency     string    `json:"currency"`
		Provider     string    `json:"provider"`
		Amount       uint      `json:"amount"`
		PaymentDate  uint      `json:"payment_dt"`
		Bank         string    `json:"bank"`
		DeliveryCost uint      `json:"delivery_cost"`
		GoodsTotal   uint      `json:"goods_total"`
		CustomFee    uint      `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      uint   `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       uint   `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        uint   `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  uint   `json:"total_price"`
		NmId        uint   `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      uint   `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              uint      `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}
