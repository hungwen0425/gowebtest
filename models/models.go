package models

import (
	"os"
	"path"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/unknown/com"
)

const (
	_DB_NAME        = "data/beego.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id          int64
	Uid         int64
	Title       string
	Content     string `orm:"size(3000)"`
	Attachment  string
	Created     time.Time `orm:"index"`
	Updated     time.Time `orm:"index"`
	Views       int64
	Author      string
	ReplyTime   time.Time `orm:"index"`
	ReplyCount  int64
	ReplyUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	//註冊 model
	orm.RegisterModel(new(Category), new(Topic))
	//註冊driver
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	//註冊默認DB
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := o.Insert(topic)
	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	category := &Category{Title: name}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(category)
	if err == nil {
		return err
	}

	_, err = o.Insert(category)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	o := orm.NewOrm()
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)

	return err
}

func DeleteTopic(tid string) error {
	o := orm.NewOrm()
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)

	return nil
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")

	_, err := qs.All(&cates)

	return cates, err
}

func GetTopic(tid string) (*Topic, error) {
	o := orm.NewOrm()
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	return topic, err
}

func ModifyTopic(tid, title, content string) error {
	o := orm.NewOrm()
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return nil
}
