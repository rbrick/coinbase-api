package coinbase

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	//PathUser is for grabbing the current authenticated user
	PathUser = "user"
	//PathAuthInfo is for grabbing the current authentication information
	PathAuthInfo = "user/auth"
	//PathUsers is for grabbing a public user from their ID.
	PathUsers = "users/%s"
)

type User struct {
	Resource
	//Name is the user's public name. Optional.
	Name string `json:"name,omitempty"`
	//Username is the user's username
	Username string `json:"username,omitempty"`
	//User's location
	Location string `json:"profile_location,omitempty"`
	//User's bio
	Bio string `json:"profile_bio,omitempty"`
	//ProfileURL is the user's profile URL
	ProfileURL string `json:"profile_url,omitempty"`
	//AvatarURL is the path to the user's avatar
	AvatarURL string `json:"avatar_url,omitempty"`

	// This section is for Authenticated users
	TimeZone       string    `json:"time_zone,omitempty"`
	NativeCurrency string    `json:"native_currency,omitempty"`
	Country        *Country  `json:"country,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type Country struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) CurrentUser() (user *User, err error) {
	req, err := c.makeRequest(http.MethodGet, PathUser, url.Values{})

	if err != nil {
		return
	}

	user = &User{}
	err = c.execute(req, user)
	return
}

func (c *Client) GetUser(id string) (user *User, err error) {
	path := fmt.Sprintf(PathUsers, id)
	req, err := c.makeRequest(http.MethodGet, path, url.Values{})

	if err != nil {
		return
	}

	user = &User{}
	err = c.execute(req, user)
	return
}
