package controller

type AppController struct {
	User interface{ UserController }
}

func ProvideAppController(user UserController) *AppController {
	return &AppController{
		User: user,
	}
}
