package main

import (
	"github.com/mikeqian/beego"
	"github.com/mikeqian/beego/example/mmapi/controllers"
)

//		mm

//	URL					 HTTP Verb				Functionality
//	/mm/create		 POST					Create mm
//	/mm/update	   	 POST					Update mm
//	/mm/list		 GET					Get mm list

func main() {
	beego.AutoRouter(&controllers.MMController{})

	beego.Run()
}
