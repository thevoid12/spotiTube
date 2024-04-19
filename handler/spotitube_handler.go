package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"spotiTube/constant"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IndexPageHandler(c *gin.Context) {
	tmpl, err := template.ParseFiles("ui/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Execute the template and write the output to the response
	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func LoginHandler(c *gin.Context) {
	//redirecting to spotify official login page
	redirectmap := map[string][]string{
		"response_type": {"code"},
		"client_id":     {constant.SPOTIFY_CLIENT_ID},
		"scope":         {constant.SPOTIFY_SCOPE},
		"redirect_uri":  {constant.HOME_PAGE_REDIRECT_URI},
		"state":         {uuid.NewString()}, //we can cross verify if we got the same state as the response.
		//this helps to prevent  Cross-Site Request Forgery. todo in future
	}

	encodedParams := url.Values(redirectmap).Encode()
	fmt.Println(constant.SPOTIFY_AUTH_URL + "?" + encodedParams)
	//redirecting for spotify authentication
	c.Redirect(http.StatusMovedPermanently, constant.SPOTIFY_AUTH_URL+"?"+encodedParams)
}

func HomePageHandler(c *gin.Context) {
	//if not a succesful authendication attempt
	code := c.Query("code")
	err := c.Query("error")
	//state := c.Query("state")
	if err != "" {
		//log the error and return
		c.String(http.StatusConflict, "Authentication failed", err)
	} else {
		//now use this code to get authid,token,expir

		requestbody := map[string][]string{
			"code":          {code},
			"client_id":     {constant.SPOTIFY_CLIENT_ID},
			"client_secret": {constant.SPOTIFY_CLIENT_SECRET},
			"grant_type":    {constant.SPOTIFY_GRAND_TYPE},
			"redirect_uri":  {constant.HOME_PAGE_REDIRECT_URI},
		}
		encodedParams := url.Values(requestbody).Encode()
		req, err := http.NewRequest("POST", constant.SPOTIFY_TOKEN_URL, strings.NewReader(encodedParams))
		if err != nil {
			// Handle error
			c.String(404, "POST request failed with status: %d", err)
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Create a HTTP client
		client := &http.Client{}

		// Send the request
		response, err := client.Do(req)
		if err != nil {
			// Handle error
			c.String(404, "POST request failed with status: %d", err)
			return
		}

		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			// If there's an error in reading the response body, respond with a server error status
			c.String(http.StatusInternalServerError, "Failed to read response body")
			return
		}
		fmt.Println(string(responseBody))
		//c.String(200, string(responseBody))
		var responsemap map[string]interface{}
		err = json.Unmarshal(responseBody, &responsemap)
		if err != nil {
			fmt.Println(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("%v", responsemap["access_token"]))
		c.String(http.StatusOK, fmt.Sprintf("%v", responsemap["expires_in"]))
		c.String(http.StatusOK, fmt.Sprintf("%v", responsemap["refresh_token"]))
		c.String(http.StatusOK, fmt.Sprintf("%v", responsemap["scope"]))
	}
}
