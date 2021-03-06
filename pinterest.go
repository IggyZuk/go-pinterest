package pinterest

import (
	"net/http"
	"time"

	"github.com/BrandonRomano/wrecker"
	"github.com/iggyzuk/go-pinterest/controllers"
)

// Client is an API client that connects you with the
// Pinterest API.  All API requests will be called through
// an instance of this struct.
//
// Do not instantiate a Client manually, but call the NewClient
// function, which will return you a properly prepared instance.
//
// For more information about the Pinterest API,
// check out https://developers.pinterest.com/
type Client struct {
	OAuth         *controllers.OAuthController
	Users         *controllers.UsersController
	Boards        *controllers.BoardsController
	Pins          *controllers.PinsController
	Me            *controllers.MeController
	wreckerClient *wrecker.Wrecker
}

// NewClient generates a new instance of a Client, which will
// allow you to interact with the Pinterest API.
func NewClient() *Client {
	// Build Wrecker client
	wc := &wrecker.Wrecker{
		BaseURL: "https://api.pinterest.com/v1",
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		DefaultContentType: "application/json",
		RequestInterceptor: nil,
	}

	// Build Pinterest client
	return &Client{
		wreckerClient: wc,
		OAuth:         controllers.NewOAuthController(wc),
		Users:         controllers.NewUsersController(wc),
		Boards:        controllers.NewBoardsController(wc),
		Pins:          controllers.NewPinsController(wc),
		Me:            controllers.NewMeController(wc),
	}
}

// RegisterAccessToken registers an AccessToken on an existing Client.
// All following requests made with the Client will be authorized with
// the specified AccessToken.
func (pc *Client) RegisterAccessToken(accessToken string) *Client {
	pc.wreckerClient.RequestInterceptor = func(req *wrecker.Request) error {
		req.URLParam("access_token", accessToken)
		return nil
	}
	return pc
}

// SetHttpClient sets the underlying http.Client that runs all API requests
func (pc *Client) SetHttpClient(client *http.Client) *Client {
	pc.wreckerClient.HttpClient = client
	return pc
}
