package controller

type AppController struct {
	User interface{ UserController }
}

func NewAppController(user UserController) *AppController {
	return &AppController{
		User: user,
	}
}
