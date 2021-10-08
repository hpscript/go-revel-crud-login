package controllers

import (
	"github.com/revel/revel"
	"app/app/models"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

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
