package controllers

import (
	"github.com/astaxie/beego"
	"myProject/models"
	"github.com/astaxie/beego/validation"
	"log"
	"crypto/md5"
	"encoding/hex"
	"time"
	"strconv"
	"myProject/cmm"
)

// RegisterController operations for Register
type RegisterController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegisterController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Register
// @Param	body		body 	models.Register	true		"body for Register content"
// @Success 201 {object} models.Register
// @Failure 403 body is empty
// @router / [post]
func (c *RegisterController) Post() {
	user := new(models.User)
	user.Email = c.GetString("email")
	user.Nickname = c.GetString("nickname")
	user.Username = c.GetString("username")
	code := md5.New()
	code.Write([]byte(c.GetString("password")))
	pwd := code.Sum(nil)
	user.Password = hex.EncodeToString(pwd)
	user.Age, _ = c.GetInt("age")
	user.Updated =  strconv.FormatInt(time.Now().Unix(),10)
	user.Created =  strconv.FormatInt(time.Now().Unix(),10)
	user.Token = cmm.GetRandomString(4)
	//验证
	valid := validation.Validation{}
	valid.Required(user.Username, "username")
	valid.Required(user.Email, "email")
	valid.Required(user.Password, "password")
	valid.Required(user.Nickname, "nickname")
	valid.MaxSize(user.Username, 16, "nameMax")
	valid.MinSize(user.Username, 6, "nameMin")
	valid.Email(user.Email, "email")
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			c.Ctx.WriteString(err.Message)
			return
		}
	}
	// or use like this
	if v := valid.Max(user.Age, 140, "age"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
		return
	}

	if v := valid.Min(user.Age, 18, "age"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
		c.Ctx.WriteString(v.Error.Message)
		return
	}

	id, err := models.AddUser(user)

	if id > 0 {
		c.Ctx.Redirect(301, "/login")
	} else {
		c.Ctx.Redirect(301, "/error/0/"+err.Error()+"/admin/blog_operate")
	}
}

// GetOne ...
// @Title GetOne
// @Description get Register by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Register
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RegisterController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Register
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Register
// @Failure 403
// @router / [get]
func (c *RegisterController) GetAll() {
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

// Put ...
// @Title Put
// @Description update the Register
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Register	true		"body for Register content"
// @Success 200 {object} models.Register
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RegisterController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Register
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RegisterController) Delete() {

}
