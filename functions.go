package util

import (
	"encoding/json"
	"fmt"

	apexhttp "github.com/River-Island/go-util/apex/http"
	mapApi "github.com/River-Island/go-util/geo/google"
	"github.com/apex/go-apex"
)

var apiKey = "AIzaSyC_dFResqlzZdLUUUvvj1nbQEbiPIzC_eo"

func CalculateGeo() apex.HandlerFunc {
	return func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var agr apexhttp.APIGatewayReq
		if err := json.Unmarshal(event, &agr); err != nil {
			return nil, err
		}

		var mBody map[string]interface{}
		switch agr.Body.(type) {
		case string:
			json.Unmarshal([]byte(agr.Body.(string)), &mBody)
		case map[string]interface{}:
			mBody = agr.Body.(map[string]interface{})
		}

		var geoReq *mapApi.GeoReq
		var err error
		switch agr.Method {
		case "POST":
			geoReq, err = postRequest(mBody)
		case "GET":
			geoReq, err = getRequest(agr.PathParams)
		default:
			// Invalid method
			return nil, fmt.Errorf("Invalid method")
		}

		client := mapApi.NewGoogleGeo(apiKey)
		resp, err := client.ResolveLocation(geoReq)
		if err != nil {
			return nil, err
		}

		return respond(resp, err)
	}
}

func getRequest(params map[string]string) (*mapApi.GeoReq, error) {
	var postcode, language, region string
	postcode, ok := params["postcode"]
	if !ok {
		return nil, fmt.Errorf("Postcode not provided")
	}
	if postcode != "" {
		language, ok = params["language"]
		if !ok {
			language = "en"
		}
		region, ok = params["region"]
		if !ok {
			region = "uk"
		}

	}
	return &mapApi.GeoReq{
		Postcode: postcode,
		Lang:     language,
		Region:   region,
	}, nil
}

func postRequest(body map[string]interface{}) (*mapApi.GeoReq, error) {
	var p mapApi.GeoReq
	b, _ := json.Marshal(body)
	if err := json.Unmarshal(b, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func respond(resp interface{}, err error) (interface{}, error) {
	if err != nil {
		return apexhttp.APIGatewayResp{
			StatusCode: 500,
		}, err
	}

	var headers = map[string]string{"Content-Type": "application/json"}

	if err != nil {
		return apexhttp.APIGatewayResp{
			StatusCode: 500,
		}, err
	}
	agr := apexhttp.APIGatewayResp{
		StatusCode: 200,
		Body:       resp,
		Headers:    headers,
	}
	return agr, nil
}
