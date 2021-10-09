package controllers

import (
	"github.com/revel/revel"
	"app/app/models"
	"fmt"

	"crypto/md5"
	"encoding/hex"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	name := c.Session["userName"]
	if name == nil {
		return c.Redirect(App.Login)
	}

	result := []models.Baseballs{}
	DB.Find(&result)

	return c.Render(result)
}

func (c App) Show(id int) revel.Result {

	result := []models.Baseballs{}
	DB.Where("id = ?", id).First(&result)
	team := result[0]
	fmt.Println(team.Name)

	return c.Render(team)
}

func (c App) Create() revel.Result {

	return c.Render()
}

func (c App) Store() revel.Result {

	name := c.Params.Get("name")
	manager := c.Params.Get("manager")
	home := c.Params.Get("home")

	c.Validation.Required(name).Message("name is required")
	c.Validation.Required(manager).Message("manager is required")
	c.Validation.Required(home).Message("home is required")

	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Create)
	}

	DB.Create(&models.Baseballs{
		Name: name,
		Manager: manager,
		Home: home,
	})

	return c.Render(name)
}

func (c App) Edit(id int) revel.Result {

	result := []models.Baseballs{}
	DB.Where("id = ?", id).First(&result)
	team := result[0]

	return c.Render(team)
}

func (c App) Update() revel.Result {

	id := c.Params.Get("id")
	name := c.Params.Get("name")
	manager := c.Params.Get("manager")
	home := c.Params.Get("home")

	DB.Model(models.Baseballs{}).Where("id = ?", id).Update(&models.Baseballs{
		Name: name,
		Manager: manager,
		Home: home,
	})

	return c.Render(name)
}

func (c App) Delete(id int) revel.Result {

	DB.Where("id = ?", id).Delete(&models.Baseballs{})

	return c.Redirect(App.Index)
}

// Login
func (c App) Login() revel.Result {

	return c.Render()
}

func (c App) Auth(name string, password string) revel.Result {

	c.Validation.Required(name).Message("name is required")
	c.Validation.Required(password).Message("password is required")

	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}

	user := []models.Users{}
	DB.Where("name = ?", name).First(&user)

	if len(user) != 0 {
		hashPassword := getMD5Hash(password)
		if user[0].Password == hashPassword {
			c.Session["userName"] = name
			c.Session.SetNoExpiration()
			c.Flash.Success("Welcome, " + name)
			return c.Redirect(App.Index)
		}
	}

	c.Flash.Error("Login failed")
	return c.Redirect(App.Login)
}

func getMD5Hash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (c App) Logout() revel.Result {

	delete(c.Session, "userName")

	return c.Redirect(App.Login)
}

// func (c App) Logined() revel.Result{
// 	name := c.Session["userName"]
// 	fmt.Println(name)

// 	if name == nil {
// 		return c.Redirect(App.Login)
// 	}
// 	return nil
// }
