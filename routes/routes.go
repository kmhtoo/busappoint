package app

import (
	"github.com/kataras/iris"
	"github.com/kmhtoo/busappoint/firm"

	acc "github.com/kmhtoo/busappoint/account"
)

func serve(ctx *iris.Context) {
	ctx.MustRender("index.html", map[string]interface{}{"username": "iris", "is_admin": true})
}

type Router struct {
}

func (r Router) Init() {
	server := iris.New()
	//server.Config.Render.Template.Engine = iris.PongoEngine
	server.Config.Render.Template.Directory = "src"
	server.StaticWeb("/dist/client/", "./", 0)
	server.StaticWeb("/static/", "./", 0)
	home := server.Party("")
	{
		home.UseFunc(func(c *iris.Context) {
			// authentication logic here
			println("from ", c.PathString())
			authorised := true
			if authorised {
				c.Next()
			} else {
				c.Text(401, c.PathString()+" is not authorized for you")
			}
		})
		home.Get("/", serve)
		home.Get("/home", serve)
		home.Get("/about", serve)
		home.Get("/**", serve)
	}
	accountRoute := server.Party("/accounts")
	{
		accountRoute.UseFunc(func(c *iris.Context) {
			// authentication logic here
			println("from ", c.PathString())
			authorised := true
			if authorised {
				c.Next()
			} else {
				c.Text(401, c.PathString()+" is not authorized for you")
			}
		})
		accountRoute.API("/", acc.AccountService{})
	}
	firmRoute := server.Party("/firms")
	{
		firmRoute.UseFunc(func(c *iris.Context) {
			// authentication logic here
			println("from ", c.PathString())
			authorised := true
			if authorised {
				c.Next()
			} else {
				c.Text(401, c.PathString()+" is not authorized for you")
			}
		})
		firmRoute.API("/", firm.FirmService{})
	}
	server.Handle("GET", "/forget", acc.AccountForget{})
	server.Get("/data.json", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{"data": "This fake data came from the db on the server and I modified it."})
	})

	server.Listen(":8080")

}
