/*
 * Cisco PSIRT openVuln API
 *
 * The Cisco Product Security Incident Response Team (PSIRT) openVuln API is a RESTful API that allows customers to obtain Cisco Security Vulnerability information in different machine-consumable formats. APIs are important for customers because they allow their technical staff and programmers to build tools that help them do their job more effectively (in this case, to keep up with security vulnerability information). For more information about the Cisco PSIRT openVuln API visit https://developer.cisco.com/site/PSIRT/discover/overview  For detail steps on how to use the API go to: https://developer.cisco.com/site/PSIRT/get-started/getting-started.gsp  This is a beta release of a swagger YAML for the Cisco PSIRT openVuln API  To access the API sign in with your Cisco CCO account at http://apiconsole.cisco.com and register an application to recieve a client_id and a client_secret  You can then get your token using curl or any other method you prefer.  'curl -s -k -H \"Content-Type: application/x-www-form-urlencoded\" -X POST -d \"client_id=<your_client_id>\" -d \"client_secret=<your_client_secret>\" -d \"grant_type=client_credentials\" https://cloudsso.cisco.com/as/token.oauth2'  You will receive an access token as demonstrated in the following example:  '{\"access_token\":\"I7omWtBDAieSiUX3shOxNJfuy4J6\",\"token_type\":\"Bearer\",\"expires_in\":3599}'  In Swagger, click on Change Authentication  enter the text \"I7omWtBDAieSiUX3shOxNJfuy4J6\" (which is the token you received)  then click on \"Try this operation\"
 *
 * API version: 0.0.4
 * Contact: os@cisco.com
 */

package swagger

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2    	= contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth 	= contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken 	= contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
 	ContextAPIKey 		= contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName      string            `json:"userName,omitempty"`
	Password      string            `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key 	string
	Prefix	string
}

type Configuration struct {
	BasePath      string            	`json:"basePath,omitempty"`
	Host          string            	`json:"host,omitempty"`
	Scheme        string            	`json:"scheme,omitempty"`
	DefaultHeader map[string]string 	`json:"defaultHeader,omitempty"`
	UserAgent     string            	`json:"userAgent,omitempty"`
	HTTPClient 	  *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://api.cisco.com",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
