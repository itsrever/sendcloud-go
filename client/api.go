package client

import (
	"github.com/itsrever/sendcloud-go/parcel"
)

type API struct {
	Parcel *parcel.Client
}

// Initialize the client
func (a *API) Init(apiKey string, apiSecret string) {
	a.Parcel = parcel.New(apiKey, apiSecret)
}
