package handle

import "go-chat-server/internal/service"

type UserHandler struct {
	emailService service.EmailService
	userService  service.UserService
}
