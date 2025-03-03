// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	about "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/about"
	auth "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/auth"
	cart "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/cart"
	category "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/category"
	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/checkout"
	home "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/home"
	order "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/order"
	product "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	about.Register(r)

	order.Register(r)

	checkout.Register(r)

	auth.Register(r)

	cart.Register(r)

	category.Register(r)

	product.Register(r)

	home.Register(r)
}
