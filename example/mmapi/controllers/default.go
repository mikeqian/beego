package controllers

import (
	"github.com/mikeqian/beego"
	_ "github.com/mikeqian/beego/example/mmapi/models"
)

type MMController struct {
	beego.Controller
}

func (this *MMController) List() {

}
