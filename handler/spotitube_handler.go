package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"spotiTube/constant"

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
			"response_type": {"code"},
			"code":          {code},
			"client_id":     {constant.SPOTIFY_CLIENT_ID},
			"client_secret": {constant.SPOTIFY_CLIENT_SECRET},
			"grant_type":    {constant.SPOTIFY_GRAND_TYPE},
			"redirect_uri":  {constant.HOME_PAGE_REDIRECT_URI},
		}
		encodedParams := url.Values(requestbody).Encode()
		fmt.Println(constant.SPOTIFY_AUTH_URL + "?" + encodedParams)

		response, err := http.Get(constant.SPOTIFY_AUTH_URL + "?" + encodedParams)
		if err != nil {
			// Handle error
			fmt.Println("Error sending request:", err)
			return
		}
		defer response.Body.Close()
		// Check the response status code
		if response.StatusCode != http.StatusOK {
			c.String(response.StatusCode, "POST request failed with status: %d", response.StatusCode)
			return
		}

		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			// If there's an error in reading the response body, respond with a server error status
			c.String(http.StatusInternalServerError, "Failed to read response body")
			return
		}
		fmt.Println(string(responseBody))
		c.String(200, string(responseBody))
		responsemap := make(map[string]interface{})
		json.Unmarshal(responseBody, &responsemap)
		// c.String(http.StatusOK, responsemap["access_token"])
		// c.String(http.StatusOK, responsemap["expires_in"])
		// c.String(http.StatusOK, responsemap["refresh_token"])
		// c.String(http.StatusOK, responsemap["scope"])
	}
}
