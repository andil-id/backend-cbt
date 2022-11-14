package web

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type LoginAdminRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ChangePasswordAuthRequest struct {
	OldPassword string `json:"password_lama" binding:"required"`
	NewPassword string `json:"password_baru" binding:"required"`
}
type ForgetPasswordAuthRequest struct {
	Username    string `json:"username" binding:"required,email"`
	NewPassword string `json:"password_baru" binding:"required"`
}
