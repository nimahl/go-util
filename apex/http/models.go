package http

type APIGatewayReq struct {
	Resource       string            `json:"resource,omitempty"`
	Path           string            `json:"path,omitempty"`
	Method         string            `json:"httpMethod,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	QueryParams    map[string]string `json:"queryStringParameters,omitempty"`
	PathParams     map[string]string `json:"pathParameters,omitempty"`
	StageVars      map[string]string `json:"stageVariables,omitempty"`
	RequestContext APIGatewayCtx     `json:"requestContext,omitempty"`
	Body           interface{}       `json:"body,omitempty"`
	IsB64Encoded   bool              `json:"isBase64Encoded,omitempty"`
}

type APIGatewayCtx struct {
	AccountID    string             `json:"accountId,omitempty"`
	ResourceID   string             `json:"resourceId,omitempty"`
	Stage        string             `json:"stage,omitempty"`
	RequestId    string             `json:"requestId,omitempty"`
	Identity     APIGatewayIdentity `json:"identity,omitempty"`
	ResourcePath string             `json:"resourcePath,omitempty"`
	HttpMethod   string             `json:"httpMethod,omitempty"`
	APIID        string             `json:"apiId,omitempty"`
}

type APIGatewayIdentity struct {
	CognitoIdentityPoolID         string `json:"cognitoIdentityPoolId,omitempty"`
	AccountID                     string `json:"accountId,omitempty"`
	CognitoIdentityID             string `json:"cognitoIdentityId,omitempty"`
	Caller                        string `json:"caller,omitempty"`
	APIKey                        string `json:"apiKey,omitempty"`
	SourceIp                      string `json:"sourceIp,omitempty"`
	AccessKey                     string `json:"accessKey,omitempty"`
	CognitoAuthenticationType     string `json:"cognitoAuthenticationType,omitempty"`
	CognitoAuthenticationProvider string `json:"cognitoAuthenticationProvider,omitempty"`
	UserARN                       string `json:"userArn,omitempty"`
	UserAgent                     string `json:"userAgent,omitempty"`
	User                          string `json:"user,omitempty"`
}

type APIGatewayResp struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Body       interface{}       `json:"body,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
}
