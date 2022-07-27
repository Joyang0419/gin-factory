package userService

type InterfaceUserService interface {
	CreateUser(vendor string, account string,
		accountType string, hashedPwd string, name string) bool
	GetAllUsers() []map[string]interface{}
	DeleteUser(id int) bool
	GetUser(id int) map[string]interface{}
	UpdateUser(id int, name string) bool
}
