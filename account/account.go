package account

import (
	"github.com/kataras/iris"
	"github.com/kmhtoo/busappoint/app"
	"time"
)

type Account struct {
	Id             int64  `db:"id,primarykey, autoincrement"`
	Email          string `db:"email,size:50"`
	Name           string `db:"name,size:100"`
	Username       string `db:"username,size:20"`
	Password       string `db:"password"`
	HashedPassword []byte `db:"hashed_password"`
	Activated      bool   `db:"activated"`
	ActivatedKey   string `db:"activated_key"`

	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}

type AccountService struct {
	*iris.Context
	*app.BusApp
}

// GET /account
func (c AccountService) Get() {
	c.Write("Get from /accounts")
	// c.JSON(iris.StatusOK,myDb.AllAccounts())
}

// GET /account/:param1 which its value passed to the id argument
func (c AccountService) GetBy(id string) {
	c.Write("Get from /accounts/%s", id)
	// c.JSON(iris.StatusOK, myDb.GetAccountById(id))
}

// PUT /accounts
func (c AccountService) Put() {
	name := c.FormValue("name")
	// myDb.InsertAccount(...)
	println(string(name))
	println("Put from /users")

}

// POST /accounts/:param1
func (c AccountService) PostBy(id string) {
	name := c.FormValue("name") // you can still use the whole Context's features!
	// myDb.UpdateUser(...)
	println(string(name))
	print("Post from /users/" + id)
}

// DELETE /accounts/:param1
func (c AccountService) DeleteBy(id string) {
	// myDb.DeleteAccount(id)
	println("Delete from /" + id)
}

func (c AccountService) Forget(ctx *iris.Context) {
	println("Forget password")
}
