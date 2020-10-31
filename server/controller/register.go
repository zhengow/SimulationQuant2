package controller

import(
	"github.com/kataras/iris"
	"brown/model"
	"brown/service"
	"log"
)

type RegisterController struct {
	Ctx     iris.Context
	Service service.RegisterService
}

func NewRegisterController() *RegisterController {
	return &RegisterController{Service: service.NewRegisterServices()}
}

func (r *RegisterController) Post() model.Result {
	var user model.User
	err := r.Ctx.ReadJSON(&user)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	// result := model.Result{}
	result := r.Service.Register(user)
	return result
}