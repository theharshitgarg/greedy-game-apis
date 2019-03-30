package db

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"project_greedy_game/models"
)

var (
	Avdertisements []models.Avdertisement
	BidRequests    []models.BidRequest
)

func GerInitialAdvertisements() []models.Avdertisement {

	t, _ := uuid.FromString("a04631b3-440e-422b-9886-95ffb9950ea6")
	ads := Avdertisements
	advt := models.Avdertisement{uuid.NewV4(), t, 1, 100, time.Now(), time.Now()}
	ads = append(ads, advt)

	advt = models.Avdertisement{uuid.NewV4(), uuid.NewV4(), 2, 200, time.Now(), time.Now()}
	ads = append(ads, advt)

	advt = models.Avdertisement{uuid.NewV4(), uuid.NewV4(), 3, 300, time.Now(), time.Now()}
	ads = append(ads, advt)

	if len(Avdertisements) == 0 {
		Avdertisements = ads
	}

	return ads
}

func GerInitialBidRequests() []models.BidRequest {

	t, _ := uuid.FromString("a04631b3-440e-422b-9886-95ffb9950ea6")
	bids := BidRequests
	bid := models.BidRequest{t, 1, false, time.Now(), time.Now(), 1}
	bids = append(bids, bid)

	bid = models.BidRequest{uuid.NewV4(), 2, true, time.Now(), time.Now(), 1}
	bids = append(bids, bid)

	bid = models.BidRequest{uuid.NewV4(), 3, false, time.Now(), time.Now(), 1}
	bids = append(bids, bid)

	if len(BidRequests) == 0 {
		BidRequests = bids
	}

	return bids
}
