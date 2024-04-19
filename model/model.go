// this go file has all the structures which we use in this project
package model

type Authorize struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
	RedirectURI  string `json:"redirect_uri"`
	State        string `json:"state"`
}

type Authorizecode struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
}
