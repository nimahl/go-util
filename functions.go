package util

import (
	"log"
	"strconv"
	"strings"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"encoding/json"
	"github.com/apex/go-apex"
)

// Business payloads for this function - needs ce
type APIGatewayReq struct {
	Resource 	string 		`json:"resource,omitempty"`
	Path 		string 		`json:"path,omitempty"`
	Method 		string 		`json:"httpMethod,omitempty"`
	Headers 	map[string]string `json:"headers,omitempty"`
	QueryParams 	map[string]string `json:"queryStringParameters,omitempty"`
	PathParams 	map[string]string `json:"pathParameters,omitempty"`
	StageVars 	map[string]string `json:"stageVariables,omitempty"`
	RequestContext	APIGatewayCtx 	  `json:"requestContext,omitempty"`
	Body		json.RawMessage   `json:"body,omitempty"`
	IsB64Encoded 	bool 		  `json:"isBase64Encoded,omitempty"`
}

type APIGatewayCtx struct {
	AccountID 	string `json:"accountId,omitempty"`
	ResourceID 	string `json:"resourceId,omitempty"`
	Stage 		string `json:"stage,omitempty"`
	RequestId 	string `json:"requestId,omitempty"`
	Identity 	APIGatewayIdentity `json:"identity,omitempty"`
	ResourcePath 	string `json:"resourcePath,omitempty"`
	HttpMethod 	string `json:"httpMethod,omitempty"`
	APIID 		string `json:"apiId,omitempty"`
}

type APIGatewayIdentity struct {
	CognitoIdentityPoolID 	string `json:"cognitoIdentityPoolId,omitempty"`
	AccountID 		string `json:"accountId,omitempty"`
	CognitoIdentityID 	string `json:"cognitoIdentityId,omitempty"`
	Caller 			string `json:"caller,omitempty"`
	APIKey 			string `json:"apiKey,omitempty"`
	SourceIp 		string `json:"sourceIp,omitempty"`
	AccessKey 		string `json:"accessKey,omitempty"`
	CognitoAuthenticationType 	string `json:"cognitoAuthenticationType,omitempty"`
	CognitoAuthenticationProvider 	string `json:"cognitoAuthenticationProvider,omitempty"`
	UserARN 		string `json:"userArn,omitempty"`
	UserAgent 		string `json:"userAgent,omitempty"`
	User 			string `json:"user,omitempty"`
}

type APIGatewayResp struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Body       CalculateGeoResp  `json:"body,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
}

// Business payloads for this function
type CalculateGeoReq struct {
	Postcode	string	`json:"postcode,omitempty"`
	FmtAdd		string 	`json:"fmtAdd,omitempty"`
	Lang		string 	`json:"language,omitempty"`
	Region		string 	`json:"region,omitempty"`
	Components	string 	`json:"components,omitempty"`
	Bounds		string 	`json:"bounds,omitempty"`
	LatLng		string 	`json:"latlng,omitempty"`
	ResultType	string 	`json:"resultType,omitempty"`
	LocationType	string 	`json:"locationType,omitempty"`
}

type CalculateGeoResp struct {
	Lat 		float64 `json:"lat,omitempty"`
	Long 		float64 `json:"long,omitempty"`
	FmtAdd 		string 	`json:"fmtAdd,omitempty"`
}

func CalculateGeo() apex.HandlerFunc {
	return func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var apiKey       = "AIzaSyC_dFResqlzZdLUUUvvj1nbQEbiPIzC_eo"
		var address	 = ""
		var components   = ""
		var bounds       = ""
		var latlng       = ""
		var resultType   = ""
		var locationType = ""
		var language	 = ""
		var region	 = ""
		var client *maps.Client
		var err error

		var agr APIGatewayReq
		if err := json.Unmarshal(event, &agr); err != nil {
			return nil, err
		}
		log.Print(string(event))
		//unquoteBody, _ := strconv.Unquote(agr.Body)
		//log.Print(unquoteBody)

		var payload CalculateGeoReq
		if err := json.Unmarshal(agr.Body, &payload); err != nil {
			return nil, err
		}

		var postcode = agr.PathParams["postcode"]
		if &postcode != nil {
			log.Print("postcode from path = "+ postcode)
		} else {
			postcode = payload.Postcode
			log.Print("postcode from payload = "+ postcode)
		}

		address = postcode
		log.Print("address for google API = "+ address)

		language = payload.Lang
		log.Print("language = "+ language)

		region = payload.Region
		log.Print("region = "+ region)

		client, err = maps.NewClient(maps.WithAPIKey(apiKey))
		check(err)

		r := &maps.GeocodingRequest{
			Address:  address,
			Language: language,
			Region:   region,
		}

		parseComponents(components, r)
		parseBounds(bounds, r)
		parseLatLng(latlng, r)
		parseResultType(resultType, r)
		parseLocationType(locationType, r)
		log.Print(r)
		resp, err := client.Geocode(context.Background(), r)
		lat := resp[0].Geometry.Viewport.NorthEast.Lat
		long := resp[0].Geometry.Viewport.NorthEast.Lng
		fmtAdd := resp[0].FormattedAddress
		log.Print(lat)
		log.Print(long)
		log.Print(fmtAdd)
		check(err)
		cgr := CalculateGeoResp{
			Lat: 	lat,
			Long:	long,
			FmtAdd:	fmtAdd,
		}
		if err != nil {
			log.Print(err)
			return nil, err
		}

		return respond(cgr, err)
	}
}

func respond(resp CalculateGeoResp, err error) (interface{}, error){
	//b, err := json.Marshal(resp)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	var headers = map[string]string{"Content-Type":"application/json"}

	agr := APIGatewayResp{
		StatusCode: 	200,
		Body:       	resp,
		Headers:	headers,
	}
	log.Print(agr)
	return agr, nil
}

func check(err error) {
	if err != nil {
		log.Print("fatal error: %s", err)
	}
}

func parseComponents(components string, r *maps.GeocodingRequest) {
	if components != "" {
		c := strings.Split(components, "|")
		for _, cf := range c {
			i := strings.Split(cf, ":")
			switch i[0] {
			case "route":
				r.Components[maps.ComponentRoute] = i[1]
			case "locality":
				r.Components[maps.ComponentLocality] = i[1]
			case "administrative_area":
				r.Components[maps.ComponentAdministrativeArea] = i[1]
			case "postal_code":
				r.Components[maps.ComponentPostalCode] = i[1]
			case "country":
				r.Components[maps.ComponentCountry] = i[1]
			}
		}
	}
}

func parseBounds(bounds string, r *maps.GeocodingRequest) {
	if bounds != "" {
		b := strings.Split(bounds, "|")
		sw := strings.Split(b[0], ",")
		ne := strings.Split(b[1], ",")

		swLat, err := strconv.ParseFloat(sw[0], 64)
		if err != nil {
			log.Fatalf("Couldn't parse bounds: %#v", err)
		}
		swLng, err := strconv.ParseFloat(sw[1], 64)
		if err != nil {
			log.Fatalf("Couldn't parse bounds: %#v", err)
		}
		neLat, err := strconv.ParseFloat(ne[0], 64)
		if err != nil {
			log.Fatalf("Couldn't parse bounds: %#v", err)
		}
		neLng, err := strconv.ParseFloat(ne[1], 64)
		if err != nil {
			log.Fatalf("Couldn't parse bounds: %#v", err)
		}

		r.Bounds = &maps.LatLngBounds{
			NorthEast: maps.LatLng{Lat: neLat, Lng: neLng},
			SouthWest: maps.LatLng{Lat: swLat, Lng: swLng},
		}
	}
}

func parseLatLng(latlng string, r *maps.GeocodingRequest) {
	if latlng != "" {
		l := strings.Split(latlng, ",")
		lat, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			log.Fatalf("Couldn't parse latlng: %#v", err)
		}
		lng, err := strconv.ParseFloat(l[1], 64)
		if err != nil {
			log.Fatalf("Couldn't parse latlng: %#v", err)
		}
		r.LatLng = &maps.LatLng{
			Lat: lat,
			Lng: lng,
		}
	}
}

func parseResultType(resultType string, r *maps.GeocodingRequest) {
	if resultType != "" {
		r.ResultType = strings.Split(resultType, "|")
	}
}

func parseLocationType(locationType string, r *maps.GeocodingRequest) {
	if locationType != "" {
		for _, l := range strings.Split(locationType, "|") {
			switch l {
			case "ROOFTOP":
				r.LocationType = append(r.LocationType, maps.GeocodeAccuracyRooftop)
			case "RANGE_INTERPOLATED":
				r.LocationType = append(r.LocationType, maps.GeocodeAccuracyRangeInterpolated)
			case "GEOMETRIC_CENTER":
				r.LocationType = append(r.LocationType, maps.GeocodeAccuracyGeometricCenter)
			case "APPROXIMATE":
				r.LocationType = append(r.LocationType, maps.GeocodeAccuracyApproximate)
			}
		}

	}
}

