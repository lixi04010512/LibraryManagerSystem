package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"

)

type Manager_lixi struct {
	ManagerId             int    `sql:"manager_id" json:"manager_id" form:"manager_id"`
	ManagerName           string `sql:"manager_name" json:"manager_name" form:"manager_name"`
	ManagerPassword       int    `sql:"manager_password" json:"manager_password" form:"manager_password"`
}
type User_lixi struct {
	UserId             int    `db:"user_id" json:"user_id" form:"user_id"`
	UserName           string `db:"user_name" json:"user_name" form:"user_name"`
	UserPassword       int    `db:"user_password" json:"user_password" form:"user_password"`
}
type Book_lixi struct {
	BookId     int `db:"book_id" json:"book_id" form:"book_id"`
	Bookname string  `db:"book_name" json:"book_name" form:"book_name"`
	Bookauthor string `db:"book_author" json:"book_author" form:"book_author"`
	Bookprice int   `db:"book_price" json:"book_price" form:"book_price"`
	Bookposition string `db:"book_position" json:"book_position" form:"book_position"`
	Booktotal int   `db:"book_total" json:"book_total" form:"book_total"`
}
type Borrow_lixi struct {
	BorrowId int    `db:"borrow_id" json:"borrow_id" form:"borrow_id"`
    Bookname string `db:"book_name" json:"book_name" form:"book_name"`
	Username string `db:"user_name" json:"user_name" form:"user_name"`
}

var Db *sql.DB

//连接数据库
func Init() {
	database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/library")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database

	var sqlStr=`CREATE TABLE IF NOT EXIST manager(
        manager_id INT PRIMARY KEY AUTOINCREMENT,
        manager_name VARCHAR(255) NOT NULL,
        manager_password INT NOT NULL
      )
    `
	var sqlStr1=`CREATE TABLE IF NOT EXIST user(
        user_id INT PRIMARY KEY AUTOINCREMENT,
        user_name VARCHAR(255) NOT NULL,
        user_password INT NOT NULL
      )
    `
	var sqlStr2=`CREATE TABLE IF NOT EXIST book(
        book_id INT PRIMARY KEY AUTOINCREMENT,
        book_name VARCHAR(2O) NOT NULL,
        book_author VARCHAR(10) NOT NULL,
        book_price INT NOT NULL,
		book_total INT NOT NULL,
		book_position VARCHAR(20) NOT NULL
      )
    `
	var sqlStr3=`CREATE TABLE IF NOT EXIST borrow(
        borrow_id INT PRIMARY KEY AUTOINCREMENT,
        book_name VARCHAR(2O) NOT NULL,
        user_name VARCHAR(20) NOT NULL,
      )
    `
	_, err = Db.Exec(sqlStr)
	if err != nil {
		return
	}
	_, err = Db.Exec(sqlStr1)
	if err != nil {
		return
	}
	_, err = Db.Exec(sqlStr2)
	if err != nil {
		return
	}
	_, err = Db.Exec(sqlStr3)
	if err != nil {
		return
	}
	//defer Db.Close()
}

//增加管理员
func (m *Manager_lixi) Insertmanager() (err error) {
	r, err := Db.Exec("insert into manager(manager_name,manager_password)values(?, ?)", m.ManagerName,m.ManagerPassword)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert manager succ:", id)
	return
}

//删除管理员
func (m *Manager_lixi) Deletemanager() (err error) {
	 r,err := Db.Exec("delete from manager where manager_id=?", m.ManagerId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println(r)
	return
}

//修改管理员
func (m *Manager_lixi)Updatemanager()(err error)  {
	r, err := Db.Exec("update manager set manager_name=?,manager_password=? where manager_id=?", m.ManagerName,m.ManagerPassword,m.ManagerId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println(r)
	return
}

//查询所有管理员
func Selectmanagers() (todos []Manager_lixi,err error) {
	var sqlStr =`select * from manager`
	r, err := Db.Query(sqlStr)
	if err != nil {
		return
	}
	for r.Next() {
		todo := Manager_lixi{}
		err = r.Scan(&todo.ManagerId, &todo.ManagerName, &todo.ManagerPassword)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select managers succ:",todos)
	return
}

//查询指定管理员
func (m *Manager_lixi) Selectmanager() (todos []Manager_lixi,err error) {
	var sqlStr =`select * from manager where manager_name=?`
	r, err := Db.Query(sqlStr,m.ManagerName)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := Manager_lixi{}
		err = r.Scan(&todo.ManagerId, &todo.ManagerName, &todo.ManagerPassword)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select manager succ:",todos)
	return
}

//管理员登录验证是否存在
func (m *Manager_lixi) Selectloginmanager() (todos []Manager_lixi,err error) {
	var sqlStr =`select * from manager where manager_name=? and manager_password=?`
	r, err := Db.Query(sqlStr,m.ManagerName,m.ManagerPassword)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := Manager_lixi{}
		err = r.Scan(&todo.ManagerId, &todo.ManagerName, &todo.ManagerPassword)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select login manager succ:",todos)
	return
}

//查询修改管理员的信息
var arrManager[] Manager_lixi
func (m *Manager_lixi) Selectupdatemanager() (todos []Manager_lixi,err error) {
	var sqlStr =`select * from manager where manager_id=?`
	r, err := Db.Query(sqlStr,m.ManagerId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := Manager_lixi{}
		err = r.Scan(&todo.ManagerId, &todo.ManagerName, &todo.ManagerPassword)
		if err != nil {
			return
		}
		todos = append(todos, todo)
		arrManager=todos
	}
	fmt.Println("select manager succ:",todos)
	return
}

//返回管理员信息到修改页面
func Backupdatemanager() (todos []Manager_lixi,err error) {
	return arrManager,nil
}

//增加用户
func (u *User_lixi)Insertuser() (err error) {
	r, err := Db.Exec("insert into user(user_name,user_password)values(?, ?)",u.UserName,u.UserPassword)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert user succ:", r)
	return
}

//删除用户
func (u *User_lixi)Deleteuser() (err error) {
	r, err := Db.Exec("delete from user where user_id=?",  u.UserId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("delete user succ:", id)
	return
}

//修改用户
func (u *User_lixi)Updateuser()(err error)  {
	r, err := Db.Exec("update user set user_name=?,user_password=? where user_id=?",u.UserName,u.UserPassword,u.UserId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("update user succ:", id)
	return
}

//查询所有用户
func Selectusers() (todos []User_lixi,err error) {
	var sqlStr =`select * from user`

	r, err := Db.Query(sqlStr)
	if err != nil {
		return
	}
	for r.Next() {
		todo := User_lixi{}
		err = r.Scan(&todo.UserId, &todo.UserName, &todo.UserPassword)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select users succ:",todos)
	return
}

//查询指定用户
func (u *User_lixi) Selectuser() (todos []User_lixi,err error) {
	var sqlStr =`select * from user where user_name=?`
	r, err := Db.Query(sqlStr,u.UserName)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := User_lixi{}
		err = r.Scan(&todo.UserId, &todo.UserName, &todo.UserPassword)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select user succ:",todos)
	return
}

//用户登录验证是否存在
func (u *User_lixi) Selectloginuser() (todos []User_lixi,err error) {
	var sqlStr =`select * from user where user_name=? and user_password=?`
	r, err := Db.Query(sqlStr,u.UserName,u.UserPassword)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := User_lixi{}
		err = r.Scan(&todo.UserId, &todo.UserName, &todo.UserPassword)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select login user succ:",todos)
	return
}

//查询修改用户的信息
var arrUser[] User_lixi
func (u *User_lixi) Selectupdateuser() (todos []User_lixi,err error) {
	var sqlStr =`select * from user where user_id=?`
	r, err := Db.Query(sqlStr,u.UserId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := User_lixi{}
		err = r.Scan(&todo.UserId, &todo.UserName, &todo.UserPassword)
		if err != nil {
			return
		}
		todos = append(todos, todo)
		arrUser=todos
	}
	fmt.Println("select user succ:",todos)
	return
}

//返回用户信息到修改页面
func Backupdateuser() (todos []User_lixi,err error) {
	return arrUser,nil
}

//增加图书
func (b *Book_lixi)Insertbook()(err error)  {
	r, err := Db.Exec("insert into book(book_name,book_author,book_price,book_total,book_position)values(?, ?,?,?,?)",b.Bookname,b.Bookauthor,b.Bookprice,b.Booktotal,b.Bookposition)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert book succ:", id)
	return
}

//删除图书
func (b *Book_lixi) Deletebook()(err error)  {
	r, err := Db.Exec("delete from book where book_id=?",  b.BookId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("delete book succ:", r)
	return
}

//修改图书
func (b *Book_lixi)Updatebook()(err error)  {
	r, err := Db.Exec("update book set book_name=?,book_author=?,book_price=?,book_total=?,book_position=? where book_id=?", b.Bookname,b.Bookauthor,b.Bookprice,b.Booktotal,b.Bookposition,b.BookId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("update book succ:", id)
	return
}
//补给图书
func (b *Book_lixi) Addbook()(err error)  {
	r, err := Db.Exec("update book set book_total=book_total+? where book_name=?", b.Booktotal,b.Bookname)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("add book succ:", id)
	return
}

//查询所有图书
func Selectbooks() (todos []Book_lixi,err error) {
	var sqlStr =`select * from book`

	r, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("1:",err)
		return
	}
	for r.Next() {
		todo := Book_lixi{}
		err = r.Scan(&todo.BookId, &todo.Bookname, &todo.Bookauthor,&todo.Bookprice,&todo.Booktotal,&todo.Bookposition)
		if err != nil {
			fmt.Println("2:",err)
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select books succ:",todos)
	return
}

//查询指定图书
func (b *Book_lixi) Selectbook()(todos []Book_lixi,err error)  {
	var sqlStr =`select * from book where book_name=?`
	r, err := Db.Query(sqlStr,b.Bookname)
	if err != nil {
		return
	}
	for r.Next() {
		todo := Book_lixi{}
		err = r.Scan(&todo.BookId, &todo.Bookname, &todo.Bookauthor,&todo.Bookprice,&todo.Booktotal,&todo.Bookposition)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select book succ:",todos)
	return
}

//查询修改图书的信息
var arrBook[] Book_lixi
func (b *Book_lixi) Selectupdatebook() (todos []Book_lixi,err error) {
	var sqlStr =`select * from book where book_id=?`
	r, err := Db.Query(sqlStr,b.BookId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := Book_lixi{}
		err = r.Scan(&todo.BookId, &todo.Bookname, &todo.Bookauthor,&todo.Bookprice,&todo.Booktotal,&todo.Bookposition)
		if err != nil {
			return
		}
		todos = append(todos, todo)
		arrBook=todos
	}
	fmt.Println("select book succ:",todos)
	return
}

//返回图书信息到修改页面
func Backupdatebook() (todos []Book_lixi,err error) {
	return arrBook,nil
}

//查询补给图书的信息
var arrTotal[] Book_lixi
func (b *Book_lixi) Selectaddbook() (todos []Book_lixi,err error) {
	var sqlStr =`select * from book where book_id=?`
	r, err := Db.Query(sqlStr,b.BookId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := Book_lixi{}
		err = r.Scan(&todo.BookId, &todo.Bookname, &todo.Bookauthor,&todo.Bookprice,&todo.Booktotal,&todo.Bookposition)
		if err != nil {
			return
		}
		todos = append(todos, todo)
		arrTotal=todos
	}
	fmt.Println("select addbook succ:",todos)
	return
}

//返回补给图书信息到修改页面
func Backaddbook() (todos []Book_lixi,err error) {
	return arrTotal,nil
}

//增加借阅
func (b *Borrow_lixi)Insertborrow()(err error)  {
	r, err := Db.Exec("insert into borrow(book_name,user_name)values(?, ?)",b.Bookname,b.Username)
	r1,err := Db.Exec("update book set book_total=book_total-1 where book_name=?",b.Bookname)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert borrow succ:", id)
	fmt.Println(r1)
	return
}

//删除借阅
func (b *Borrow_lixi) Deleteborrow() (err error) {
	r, err := Db.Exec("delete from borrow where book_name=? and user_name=?",b.Bookname,b.Username)
	r1,err := Db.Exec("update book set book_total=book_total+1 where book_name=?",b.Bookname)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("delete borrow succ:",r)
	fmt.Println(r1)
	return
}

//修改借阅
func (b *Borrow_lixi) Updateborrow()(err error)  {
	r, err := Db.Exec("update borrow set book_name=?,user_name=? where borrow_id=?",b.Bookname,b.Username,b.BorrowId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("update borrow succ:", id)
	return
}

//查询所有借阅
func Selectborrows() (todos []Borrow_lixi,err error) {
	var sqlStr =`select * from borrow`
	r, err := Db.Query(sqlStr)
	if err != nil {
		return
	}
	for r.Next() {
		todo := Borrow_lixi{}
		err = r.Scan(&todo.BorrowId, &todo.Bookname, &todo.Username)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select borrows succ:",todos)
	return
}

//查询指定借阅
func (b *Borrow_lixi)Selectborrow()(todos []Borrow_lixi,err error)  {
	var sqlStr =`select * from borrow where user_name=?`
	r, err := Db.Query(sqlStr,b.Username)
	if err != nil {
		return
	}
	for r.Next() {
		todo := Borrow_lixi{}
		err = r.Scan(&todo.BorrowId, &todo.Bookname, &todo.Username)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select borrow succ:",todos)
	return
}

//查询修改借阅的信息
var arrBorrow[] Borrow_lixi
func (b *Borrow_lixi) Selectupdateborrow() (todos []Borrow_lixi,err error) {
	var sqlStr =`select * from borrow where borrow_id=?`
	r, err := Db.Query(sqlStr,b.BorrowId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := Borrow_lixi{}
		err = r.Scan(&todo.BorrowId, &todo.Bookname, &todo.Username)
		if err != nil {
			return
		}
		todos = append(todos, todo)
		arrBorrow=todos
	}
	fmt.Println("select book succ:",todos)
	return
}

//返回借阅信息到修改页面
func Backupdateborrow() (todos []Borrow_lixi,err error) {
	return arrBorrow,nil
}

//借阅人是否已经有借阅
func (b *Borrow_lixi) Selectborrowuser() (todos []Borrow_lixi,err error) {
	var sqlStr =`select * from borrow where user_name=?`
	r, err := Db.Query(sqlStr,b.Username)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for r.Next() {
		todo := Borrow_lixi{}
		err = r.Scan(&todo.BorrowId, &todo.Bookname, &todo.Username)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select login user succ:",todos)
	return
}