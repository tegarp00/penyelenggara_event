package main

import (
	"fmt"
  "github.com/spf13/viper"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
)



func main()  {


    // set config untuk membaca variabel 
    viper.SetConfigType("json")
    viper.AddConfigPath(".")
    viper.SetConfigName("app.config")

    err := viper.ReadInConfig()
    if err != nil {
      fmt.Println("Config not reading")
    }

    // mengkoneksikan database pakai gorm
    dsn := viper.GetString("server.DB_URL")
    _, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
      log.Fatal(err.Error())
    }
    fmt.Println("db Connected!")
}
