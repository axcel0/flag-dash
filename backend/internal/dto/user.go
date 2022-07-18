package dto

type UserLoginRequest struct {
	Email	string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	NormalJWT string `json:"token"`
	RefreshJWT string `json:"refreshToken"`
}

type CreateUserRequest struct {
	Email 		string 		`json:"email"`
	Password 	string 		`json:"password"`
	FirstName 	string 		`json:"firstName"`
	LastName 	string 		`json:"lastName"`
	PhoneNumber string		`json:"phoneNumber"`
}

type CreateUserResponse struct {
	Status	string  `json:"status"`
	Msg		string 	`json:"msg"`
}

type UserLogin struct {
	Email		string 	`json:"email"`
	Password 	string  `json:"password"`
}

type EditUserRole struct {
	UserID uint32 `json:"userId"`
	RoleID uint32 `json:"roleId"`
}

type GetUserRequest struct {
	Email string `json:"email"`
}

type GetUserResponse struct {
	Status	string  `json:"status"`
	Msg		string 	`json:"msg"`
	User 	struct {
		Email string `json:"email"`
		Profile struct {
			FirstName string `json:"firstName"`
			LastName string `json:"lastName"`
		} `json:"userProfile"`
		Role struct {
			Name string `json:"Name"`
			Level uint32 `json:"Level"`
		} `json:"userRole"`
	}
}

type UserAuthorizedResponse struct {
	Status	string `json:"status"`
	Msg		string `json:"msg"`
}

type UserRefreshTokenRequest struct {
	RefreshToken	string	`json:"refresh_token"`
}

type UserRefreshTokenResponse struct {
	Status 	string `json:"status"`
	Token	string `json:"token"`
}

type EditUserRequest struct {
	UserID	uint32	`json:"id"`
	Email	string	`json:"email"`
	Password	string `json:"password"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	PhoneNumber	string	`json:"phone_number"`
	RoleLevel	uint32	`json:"role_level"`
}

type EditUserResponse struct {
	Status	string	`json:"status"`
	Msg		string	`json:"msg"`
	User 	struct {
		Email string `json:"email"`
		Profile struct {
			FirstName string `json:"firstName"`
			LastName string `json:"lastName"`
			PhoneNumber	string `json:"phone_number"`
		} `json:"userProfile"`
		Role struct {
			Name string `json:"Name"`
			Level uint32 `json:"Level"`
		} `json:"userRole"`
	}
}

type DeleteUserRequest struct {
	UserID	uint32	`json:"id"`
}

type DeleteUserResponse struct {
	Status	string	`json:"status"`
	Msg		string	`json:"msg"`
}