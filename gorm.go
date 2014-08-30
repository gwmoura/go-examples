package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/spf13/viper"
    "strings"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"Welcome":"Go API"})
    })
    var wp WpUser
    r.GET("/users",wp.getUsers)
    h := viper.GetString("host")
    //server.host = "localhost"
    fmt.Println("Run server "+h);
    // Listen and server on 0.0.0.0:8080
    r.Run(":8080")

}

func connect() gorm.DB{
	h := viper.GetString("host")
	port := viper.GetString("port")
	u := viper.GetString("user")
	p := viper.GetString("pass")
	dbname := viper.GetString("dbname")
	db, err := gorm.Open("mysql", u+":"+p+"@tcp("+h+":"+port+")/"+dbname+"?charset=utf8")
	if err != nil {
	    panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	db.LogMode(true)
	db.DB()
	return db
}

func getDataBases() []string {
	db := connect()
	var dbs []string
	//db.Raw("SHOW databases").Scan(&dbs)
	
	rows, err := db.Raw("SHOW databases").Rows()
	if err != nil{
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
	  Database := ""
	  rows.Scan(&Database)
	  if strings.HasPrefix(Database,"sj_"){
	  	if !strings.Contains(Database,"join") && Database!="sj_gip" && Database!="sj_padrao"{
	  		dbs = append(dbs,Database)
	  	}
	  }
	}

	return dbs
}

type WpUser struct{
	Id int64
	UserLogin string
	UserEmail string
}

func (wp WpUser) getUsers(c *gin.Context){
	db := connect()
	var users []WpUser
	for _,dbs := range getDataBases(){
		//fmt.Println(db)
		db.Raw("SELECT * FROM `"+dbs+"`.wp_users;").Scan(&users)
	}
	//db.Find(&users)
    c.JSON(200,users)
}