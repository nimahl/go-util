package util

import (
	"encoding/json"
	"testing"

	"github.com/River-Island/go-util/apex/http"
	"github.com/apex/go-apex"
)

var event = json.RawMessage(`
{
  "resource":"/v1/geo/calculate",
  "path":"/v1/geo/calculate",
  "httpMethod":"POST",
  "headers": {
    "Accept":"application/json",
    "Cache-Control":"No Cache"
  },
  "queryStringParameters": {
    "yes":"2","hello":"1"
  },
  "pathParameters": {
    "postcode":"N18JR"
  },
  "stageVariables":null,
  "requestContext": {
    "accountId":"556748783639",
    "resourceId":"txz6fn",
    "stage":"test-invoke-stage",
    "requestId":"test-invoke-request",
    "identity":{
      "cognitoIdentityPoolId":null,
      "accountId":"556748783639",
      "cognitoIdentityId":null,
      "caller":"AIDAJ7TA4SLQYPM6FNOVK",
      "apiKey":"test-invoke-api-key",
      "sourceIp":"test-invoke-source-ip",
      "accessKey":"ASIAIXWHK2ADKS4POUMA",
      "cognitoAuthenticationType":null,
      "cognitoAuthenticationProvider":null,
      "userArn":"arn:aws:iam::556748783639:user/kdmaile",
      "userAgent":"Apache-HttpClient/4.5.x (Java/1.8.0_102)",
      "user":"AIDAJ7TA4SLQYPM6FNOVK"
    },
    "resourcePath":"/v1/geo/calculate",
    "httpMethod":"POST",
    "apiId":"8f41tsdkvh"
  },
  "body" : {
    "postcode" : "N18JR",
    "language" : "en",
    "region" : "uk"
  },
  "isBase64Encoded":false
}
`)

func TestCalculateGeo(t *testing.T) {
	resp, err := CalculateGeo().Handle(event, &apex.Context{})
	if err != nil {
		t.Fatal(err)
	}
	r := resp.(http.APIGatewayResp)
	t.Log(r.Body)
}
