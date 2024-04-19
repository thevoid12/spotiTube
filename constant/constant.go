// this file will have all the constants used for this project
package constant

const (
	SPOTIFY_CLIENT_ID      = "09b8932df912423fa73d01dc37003f2e"
	SPOTIFY_CLIENT_SECRET  = "90824f5ce69e4270af6abf3c6560bd26"
	SPOTIFY_SCOPE          = "playlist-read-private user-library-read" //what are the access which we are giving to this app
	SPOTIFY_GRAND_TYPE     = "authorization_code"
	INDEX_PAGE_URI         = "http://localhost:8080/index"
	HOME_PAGE_REDIRECT_URI = "http://localhost:8080/home"
	SPOTIFY_AUTH_URL       = "https://accounts.spotify.com/authorize"
	SPOTIFY_TOKEN_URL      = "https://accounts.spotify.com/api/token"
	SPOTIFY_API_BASE_URL   = "https://accounts.spotify.com/v1/"
)
