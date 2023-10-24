package parcel

import (
	"context"
	"strconv"

	"github.com/itsrever/sendcloud-go"
)

type Client struct {
	apiKey    string
	apiSecret string
}

func New(apiKey string, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

// Create a new parcel
func (c *Client) New(ctx context.Context, params *sendcloud.ParcelParams) (*sendcloud.Parcel, error) {
	parcel := sendcloud.ParcelResponseContainer{}
	err := sendcloud.Request(ctx, "POST", "/api/v2/parcels", params, c.apiKey, c.apiSecret, &parcel)

	if err != nil {
		return nil, err
	}
	r := parcel.GetResponse().(*sendcloud.Parcel)
	return r, nil
}

// Return a single parcel
func (c *Client) Get(ctx context.Context, parcelID int64) (*sendcloud.Parcel, error) {
	parcel := sendcloud.ParcelResponseContainer{}
	err := sendcloud.Request(ctx, "GET", "/api/v2/parcels/"+strconv.Itoa(int(parcelID)), nil, c.apiKey, c.apiSecret, &parcel)

	if err != nil {
		return nil, err
	}
	r := parcel.GetResponse().(*sendcloud.Parcel)
	return r, nil
}

// Get a label as bytes based on the url that references the PDF
func (c *Client) GetLabel(ctx context.Context, labelURL string) ([]byte, error) {
	data := &sendcloud.LabelData{}
	err := sendcloud.Request(ctx, "GET", labelURL, nil, c.apiKey, c.apiSecret, data)
	if err != nil {
		return nil, err
	}
	return *data, nil
}
