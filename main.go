package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// {	sample response
// 	userId: 1,
// 	id: 1,
// 	title: "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
// 	body: "quia et suscipit
// 	suscipit recusandae consequuntur expedita et cum
// 	reprehenderit molestiae ut ut quas totam
// 	nostrum rerum est autem sunt rem eveniet architecto"
// 	},

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// sample model for beego orm table
// Model Struct
// type Details struct {
// 	Id   int    `orm:"auto"`
// 	Name string `orm:"size(100)"`
// 	Age  int    `orm:"size(100)"`
// }

type StorePost struct {
	UserId int    `orm:"size(100)"`
	Id     int    `orm:"size(100)"`
	Title  string `orm:"size(200)"`
	Body   string `orm:"size(10000)"`
}

func init() {
	// register model
	orm.RegisterModel(new(StorePost))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8", 30)

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var posts []Post
	if err := json.Unmarshal(content, &posts); err != nil {
		fmt.Println("Error Unmarshalling data")
	}

	for _, p := range posts {
		fmt.Println("The user_id is", p.UserId)
		fmt.Println("The id is", p.Id)
		fmt.Println("The title is", p.Title)
		fmt.Println("The body is", p.Body)
	}

	fmt.Println("Unmarshalling Done and stored in a slice")

	o := orm.NewOrm()

	for _, p := range posts {
		_, err := o.Raw("insert into store_post(user_id, id, title, body) values(?,?,?,?);", p.UserId, p.Id, p.Title, p.Body).Exec()
		if err != nil {
			fmt.Println("Cannot insert in the table")
			fmt.Println(err)
		}
	}
	fmt.Println("Done")
}
