package route

import (
	"html/template"
	"login/dao"
	"login/model"
	"net/http"
)

var userDao = new(dao.UserDao)

func Route() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/doLogin.do", doLoginHandler)
	http.HandleFunc("/register.do", registerHandler)
	http.HandleFunc("/doRegister.do", doRegisterHandler)
	http.HandleFunc("/userList.do", userListHandler)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//ParseFiles函数创建一个模板并解析filenames指定的文件里的模板定义。
	//返回的模板的名字是第一个文件的文件名（不含扩展名），内容为解析后的第一个文件的内容。
	//至少要提供一个文件。如果发生错误，会停止解析并返回nil。
	tmpl, err := template.ParseFiles("view/index.html")
	checkErr(err)

	err = tmpl.Execute(w, nil)
	checkErr(err)
}

func doLoginHandler(w http.ResponseWriter, r *http.Request) {
	//FormValue返回key为键查询r.Form字段得到结果[]string切片的第一个值。
	name := r.FormValue("name")
	password := r.FormValue("password")

	sql := "select * from t_user where name='" + name + "' and password='" + password + "'"

	objs := userDao.QueryList(sql)
	tmpl, err := template.ParseFiles("view/logined.html")
	checkErr(err)

	if len(objs) == 0 {
		err = tmpl.Execute(w, map[string]string{"err_msg": "username or password is invalid!"})
		checkErr(err)
		return
	}

	var user *model.User = objs[0]
	err = tmpl.Execute(w, map[string]interface{}{"user": user})
	checkErr(err)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/register.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

func doRegisterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")

	sql := "insert into t_user (name, password) values ('" + name + "','" + password + "')"
	userDao.Save(sql)

	tmpl, err := template.ParseFiles("view/registered.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

func userListHandler(w http.ResponseWriter, r *http.Request) {
	sql := "select * from t_user"

	objs := userDao.QueryList(sql)

	m := make(map[string]interface{})
	m["user"] = objs

	tmpl, err := template.ParseFiles("view/list.html")
	checkErr(err)
	err = tmpl.Execute(w, m)
	checkErr(err)
}
