package model
import (
	"fmt"
	db "gin/utils"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
}
/*
 *description：添加用户
 */
func AddUser(user *User) error {
	sql := "insert into user(name,password,phone) values(?,?,?)"
	fmt.Println(sql)
	_, err := db.Db.Exec(sql,  &user.Name, &user.Password, &user.Phone)
	if err != nil {
		return err
	}
	return nil
}
/*
 *description：获取单个用户
 */
func GetUserByID(id int) (*User, error) {
	sql := "select * from user where id = ?"
	row := db.Db.QueryRow(sql, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}
/*
 *description：获取所有用户
 */
func GetUsers() ([]*User, error) {
	sql := "select * from user"
	rows, err := db.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan( &user.ID,&user.Name, &user.Password, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
/*
 *description：更新用户信息
 */
func UpdateUser(user *User) error {
	sql := "update user set name=?,password=?,phone=? where id=?"
	_, err := db.Db.Exec(sql, &user.Name, &user.Password, &user.Phone,&user.ID)
	if err != nil {
		return err
	}
	return nil
}
/*
 *description：删除单个用户
 */
func DeleteUserByID(id int) error {
	sql := "delete from user where id=?"
	_, err := db.Db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}