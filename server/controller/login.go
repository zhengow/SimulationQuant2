package controller

import(
	"github.com/kataras/iris"
	"brown/model"
	"brown/service"
	"log"
)

type LoginController struct {
	Ctx     iris.Context
	Service service.LoginService
}

func NewLoginController() *LoginController {
	return &LoginController{Service: service.NewLoginServices()}
}

func (l *LoginController) Post() model.Result {
	var user model.User
	err := l.Ctx.ReadJSON(&user)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	// result := model.Result{}
	result := l.Service.Login(user)
	return result
}