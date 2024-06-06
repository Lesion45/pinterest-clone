package request

// Request fields
// * login
// * password
// * sessionInfo - NOT RELEASED - IN PROCESS
type RegisterReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Request fields
// * username
// * password
// * sessionInfo - NOT RELEASED - IN PROCESS
type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Request fields
// * username
// * image-url
// * sessionInfo - NOT RELEASED - IN PROCESS
type AddPinReq struct {
	Username string `json:"username" validate:"required"` // may be changed to username
	ImageURL string `json:"image-url" validate:"required"`
	// Description string `json:"description"`
}

// Request fields
// * pin-id
// * sessionInfo - NOT RELEASED - IN PROCESS
type DeletePinReq struct {
	// Username string `json:"username" validate:"required"`
	PinID string `json:"pin-id" validate:"required"`
}
