// tools will have additional repeated utilities
package tools

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func SetCookie(w http.ResponseWriter, name string, responsemap map[string]interface{}) {
	responsemap[name] = uuid.New()
	value, err := json.Marshal(responsemap)
	if err != nil {

	}
	cookie := http.Cookie{
		Name:  name,
		Value: string(value),
	}
	http.SetCookie(w, &cookie)
}

// GetCookieValue retrieves the value of the cookie with the given name from the request
func GetCookieValue(r *http.Request, name string) (map[string]interface{}, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}
	var responsemap map[string]interface{}
	err = json.Unmarshal([]byte(cookie.Value), &responsemap)
	if err != nil {

	}
	return responsemap, nil
}
