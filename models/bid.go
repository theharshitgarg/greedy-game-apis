package models

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"project_greedy_game/utils"
)

type BidRequest struct {
	SID       uuid.UUID `json:"sid"`
	ID        int64     `json:"id"`
	IsBooking bool      `json:"is_booking"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int64     `json:"uid"`
}

type Avdertisement struct {
	SID       uuid.UUID `json:"sid"`
	BidSID    uuid.UUID `json:"bid_id"`
	ID        int64     `json:"id"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (bdR *BidRequest) Create(b_id int64, is_book bool) {
	bdR.SID = uuid.NewV4()
	bdR.ID = b_id
	bdR.IsBooking = is_book

	return
}

func (avt *Avdertisement) CreateAdvt(bidId uuid.UUID) {
	avt.BidSID = bidId
	avt.SID = uuid.NewV4()
	avt.ID = 1
	avt.Price = int64(utils.GenerateRandomRange(100, 200))
	return
}
