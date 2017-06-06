package models

import "github.com/astaxie/beego/orm"

func AddUser(username string) error {
	o := orm.NewOrm()
	o.Using("default")

	err := o.Raw("INSERT INTO users SET username = ?", username).QueryRow()

	return err
}

func CheckUser(username string) error {
	o := orm.NewOrm()
	o.Using("default")

	err := o.Raw("SELECT EXISTS(SELECT * FROM users WHERE username = ?)", username).QueryRow()

	// var p []byte
	// err := db.QueryRow("SELECT data FROM message WHERE data->>'id'=$1", id).Scan(&p)

	return err
}

// func rowExists(query string, args ...interface{}) bool {
//
// 	var exists bool
// 	query = fmt.Sprintf("SELECT exists (%s)", query)
// 	err := db.QueryRow(query, args...).Scan(&exists)
// 	if err != nil && err != sql.ErrNoRows {
// 		glog.Fatalf("error checking if row exists '%s' %v", args, err)
// 	}
// 	return exists
// }
