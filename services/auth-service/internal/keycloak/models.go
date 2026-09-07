package keycloak


type TokenResponse struct {

	AccessToken string `json:"access_token"`

	ExpiresIn int `json:"expires_in"`

	TokenType string `json:"token_type"`
}



type CreateUserRequest struct {

	Username string `json:"username"`

	Email string `json:"email"`

	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	Enabled bool `json:"enabled"`

	Credentials []Credential `json:"credentials"`
}



type Credential struct {

	Type string `json:"type"`

	Value string `json:"value"`

	Temporary bool `json:"temporary"`
}