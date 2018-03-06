package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Blog struct {
	Id          int64
	Title       string `orm:"size(128)"`
	Short_title string `orm:"size(128)"`
	Content     string `orm:"type(longtext)"`
	Updated_at  string `orm:"size(128)"`
	Created_at  string `orm:"size(128)"`
	Display     int
	Deleted     int
	Uid         int64
	Status      int
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(Blog))
	orm.RegisterDataBase("default", "mysql", "root:root@/beego?charset=utf8")
}

// AddBlog insert a new Blog into database and returns
// last inserted Id on success.
func AddBlog(m *Blog) (id int64, err error) {
	isHave := CheckBlog(m.Title)
	if isHave {
		o := orm.NewOrm()
		id, err = o.Insert(m)
		return id, err
	} else {
		id = 0
		return id, errors.New("博客已经存在")
	}
}

//jude Blog is exists
func CheckBlog(Blogname string) (bool) {
	Blog, _ := GetBlogByBlogName(Blogname)
	if Blog == nil {
		return true
	}
	return false
}

// GetBlogByBlogname retrieves Blog by Blogname. Returns error if
// Blogname doesn't exist
func GetBlogByBlogName(Title string) (v *Blog, err error) {
	o := orm.NewOrm()
	v = &Blog{Title: Title}
	if err = o.QueryTable(new(Blog)).Filter("title", Title).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetBlogById retrieves Blog by Id. Returns error if
// Id doesn't exist
func GetBlogById(id int64) (v *Blog, err error) {
	o := orm.NewOrm()
	v = &Blog{Id: id}
	if err = o.QueryTable(new(Blog)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBlog retrieves all Blog matches certain condition. Returns empty list if
// no records exist
func GetAllBlog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Blog))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Blog
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBlog updates Blog by Id and returns error if
// the record to be updated doesn't exist
func UpdateBlogById(m *Blog) (err error) {
	o := orm.NewOrm()
	v := Blog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBlog deletes Blog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBlog(id int64) (err error) {
	o := orm.NewOrm()
	v := Blog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Blog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
