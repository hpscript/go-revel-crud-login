package controllers
 
import (
    "github.com/revel/revel"
    "github.com/jinzhu/gorm"
    "strings"
    "fmt"
    _"github.com/go-sql-driver/mysql"
)
 
var DB *gorm.DB
 
 
func InitDB(){
    db, err := gorm.Open("mysql", getConnectionString())
 
    if err != nil {
        panic(err.Error())
    }
 
    db.DB()
    DB = db
}
 
 
func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            fmt.Sprintf("Could not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}
 
 
func getConnectionString() string {
    host := getParamString("db.host", "localhost")
    port := getParamString("db.port", "3306")
    user := getParamString("db.user", "hoge")
    pass := getParamString("db.password", "hogehoge")
    dbname := getParamString("db.name", "test")
    protocol := getParamString("db.protocol", "tcp")
    dbargs := getParamString("dbargs", "    ")
    timezone := getParamString("db.timezone", "parseTime=True")
 
    if strings.Trim(dbargs, " ") != ""{
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
 
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s", user, pass, protocol, host, port, dbname, timezone)
}