package admin

import (
	"github.com/astaxie/beego"
)

// ComponentsController operations for Components
type ComponentsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ComponentsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *ComponentsController) Calendar() {
	c.TplName = "admin/calendar.html"
}

func (c *ComponentsController) Gallery() {
	c.TplName = "admin/gallery.html"
}

func (c *ComponentsController) TodoList() {
	c.TplName = "admin/todo_list.html"
}

// Post ...
// @Title Create
// @Description create Components
// @Param	body		body 	models.Components	true		"body for Components content"
// @Success 201 {object} models.Components
// @Failure 403 body is empty
// @router / [post]
func (c *ComponentsController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Components by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Components
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ComponentsController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Components
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Components
// @Failure 403
// @router / [get]
func (c *ComponentsController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Components
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Components	true		"body for Components content"
// @Success 200 {object} models.Components
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ComponentsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Components
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ComponentsController) Delete() {

}
