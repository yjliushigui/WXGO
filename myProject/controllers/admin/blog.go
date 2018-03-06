package admin

import (
	"github.com/astaxie/beego"
	"myProject/models"
	"fmt"
	"strconv"
	"time"
)

// BlogController operations for Table
type BlogController struct {
	beego.Controller
}

// URLMapping ...
func (c *BlogController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *BlogController) Get() {
	var ma map[string]string
	field := []string{"Id", "Title", "Short_title", "Created_at", "Updated_at", "Display", "Status", "Deleted"}
	sortby := []string{"id"}
	order := []string{"desc"}
	c.Data["Active"] = "blog"
	m := c.GetSession("User")
	Blogs, err := models.GetAllBlog(ma, field, sortby, order, 0, 10)
	if m != nil {
		if len(Blogs) > 0 {
			if err == nil {
				c.Data["Blogs"] = Blogs
			}
		} else {
			Blogs = nil
		}
		c.Data["User"] = m

		c.TplName = "admin/blog/blog_list.html"
	} else {
		c.Ctx.Redirect(301, "/login")
	}
}

func (c *BlogController) Operate() {
	c.Data["Active"] = "blog"
	m := c.GetSession("User")
	c.Data["User"] = m
	id, _ := c.GetInt64("id")
	if id > 0 {
		blog, err := models.GetBlogById(id)
		if err == nil {
			c.Data["Blog"] = blog
		}
	} else {
		c.Data["Blog"] = nil
	}
	c.TplName = "admin/blog/blog_edit.html"

}

// Post ...
// @Title Create
// @Description create Table
// @Param	body		body 	models.Table	true		"body for Table content"
// @Success 201 {object} models.Table
// @Failure 403 body is empty
// @router / [post]

func (c *BlogController) Post() {
	//var m  *models.User
	blog := new(models.Blog)
	blog.Title = c.GetString("title")
	blog.Short_title = c.GetString("short_title")
	blog.Content = c.GetString("content")
	blog.Status, _ = c.GetInt("status")
	blog.Display, _ = c.GetInt("display")
	blog.Id, _ = c.GetInt64("id")
	m := c.GetSession("User")
	if m != nil {
		fmt.Println(m)
		user := m.(*models.User)
		if user.Id > 0 {
			blog.Uid = user.Id
			fmt.Println(user.Id)
			if blog.Id > 0 {
				//编辑博客
				blog.Updated_at = strconv.FormatInt(time.Now().Unix(), 10)
				models.UpdateBlogById(blog)
				c.Ctx.Redirect(301, "/error/1/编辑成功！")
			} else {
				//新增博客
				blog.Updated_at = strconv.FormatInt(time.Now().Unix(), 10)
				blog.Created_at = strconv.FormatInt(time.Now().Unix(), 10)
				id, err := models.AddBlog(blog)
				if id > 0 {
					c.Ctx.Redirect(301, "/error/1/新增成功！")
				} else {
					c.Ctx.Redirect(301, "/error/0/"+err.Error())
				}
			}
		}
	} else {
		c.Ctx.Redirect(301, "/login")
	}

}

// GetOne ...
// @Title GetOne
// @Description get Table by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Table
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BlogController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Table
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Table
// @Failure 403
// @router / [get]
func (c *BlogController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Table
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Table	true		"body for Table content"
// @Success 200 {object} models.Table
// @Failure 403 :id is not int
// @router /:id [put]
func (c *BlogController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Table
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *BlogController) Delete() {

}
