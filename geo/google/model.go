package google

// Business payloads for this function
type GeoReq struct {
	Postcode     string `json:"postcode,omitempty"`
	FmtAdd       string `json:"fmtAdd,omitempty"`
	Lang         string `json:"language,omitempty"`
	Region       string `json:"region,omitempty"`
	Components   string `json:"components,omitempty"`
	Bounds       string `json:"bounds,omitempty"`
	LatLng       string `json:"latlng,omitempty"`
	ResultType   string `json:"resultType,omitempty"`
	LocationType string `json:"locationType,omitempty"`
}

type GeoResp struct {
	Lat    float64 `json:"lat,omitempty"`
	Long   float64 `json:"long,omitempty"`
	FmtAdd string  `json:"fmtAdd,omitempty"`
	Type   string  `json:"locationType,omitempty"`
}
