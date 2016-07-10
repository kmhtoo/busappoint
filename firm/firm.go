package firm

import (
	"github.com/kataras/iris"
	"github.com/kmhtoo/busappoint/app"
)

type Firm struct {
}

type FirmService struct {
	*iris.Context
	*app.BusApp
}

// GET /account
func (c FirmService) Get() {
	c.Write("Get from /firms")
	// c.JSON(iris.StatusOK,myDb.Allfirms())
}

// GET /account/:param1 which its value passed to the id argument
func (c FirmService) GetBy(id string) {
	c.Write("Get from /firms/%s", id)
	// c.JSON(iris.StatusOK, myDb.GetAccountById(id))
}

// PUT /firms
func (c FirmService) Put() {
	name := c.FormValue("name")
	// myDb.InsertAccount(...)
	println(string(name))
	println("Put from /users")

}

// POST /firms/:param1
func (c FirmService) PostBy(id string) {
	name := c.FormValue("name") // you can still use the whole Context's features!
	// myDb.UpdateUser(...)
	println(string(name))
	print("Post from /users/" + id)
}

// DELETE /firms/:param1
func (c FirmService) DeleteBy(id string) {
	// myDb.DeleteAccount(id)
	println("Delete from /" + id)
}

func (c FirmService) Forget(ctx *iris.Context) {
	println("Forget password")
}
