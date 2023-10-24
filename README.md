To initialize the client:
```go
api := client.API{}
api.Init("api_key", "api_secret")
```
Create a parcel:
```go
params := &sendcloud.ParcelParams{
	Name:             "Sendcloud-GO",
	CompanyName:      "Afosto SaaS BV",
	Street:           "Grondzijl",
	HouseNumber:      "16",
	City:             "Groningen",
	PostalCode:       "9731DG",
	PhoneNumber:      "0507119519",
	EmailAddress:     "peter@afosto.io",
	CountryCode:      "NL",
	IsLabelRequested: true,
	Method:           8,
	ExternalID:       uuid.New().String(),
}
parcel, err := api.Parcel.New(params)
```