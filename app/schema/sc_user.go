package schema

// User schema
type User struct {
	ID       string      `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Extra    interface{} `json:"extra,omitempty"`
}

// RegisterBodyParams schema
type RegisterBodyParams struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	RoleID   string `json:"role_id"`
}

// LoginBodyParams schema
type LoginBodyParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

// RefreshBodyParams schema
type RefreshBodyParams struct {
	RefreshToken string `json:"refresh_token,omitempty" validate:"required"`
}

// UserQueryParam schema
type UserQueryParam struct {
	Username string `json:"username,omitempty" form:"username,omitempty"`
	Email    string `json:"email,omitempty" form:"email,omitempty"`
	Offset   int    `json:"-" form:"offset,omitempty"`
	Limit    int    `json:"-" form:"limit,omitempty"`
}

// UserTokenInfo schema
type UserTokenInfo struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	TokenType    string   `json:"token_type"`
	Roles        []string `json:"roles"`
}

// UserUpdateBodyParam schema
type UserUpdateBodyParam struct {
	Password     string `json:"password,omitempty"`
	RoleID       string `json:"role_id,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
