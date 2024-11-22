package main

import (
	"fmt"
	"mrs-go/service"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	dsn := viper.GetString("mysql.dsn")
	dialetor := mysql.Open(dsn)
	db, err := gorm.Open(dialetor)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Success")

	service := service.NewshowDataService(db)
	landmark, err := service.GetLandmarkByName("ปากช่อง")
	if err != nil {
		panic(err)
	}
	for _, item := range *landmark {
		fmt.Printf("%v\n", item)
	}

	result := service.UpdatelandmarkByNameAndByDetail(1, "เป็นคนปากช่อง", "และปากช่องเป็นประเทศ")
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("Update completed")
}
