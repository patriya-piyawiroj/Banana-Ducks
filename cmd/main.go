//go:generate swagger generate spec
package main

import (
	"github.com/patriya-piyawiroj/banana-ducks/internal/app/bnnd"
	"github.com/patriya-piyawiroj/banana-ducks/internal/app/shop"

	_ "github.com/patriya-piyawiroj/banana-ducks/docs"

	"github.com/labstack/echo/v4"
)

// Main function
func main() {
	e := echo.New()

	// ===================== init configs =============================
	// f, err := os.Open("configs/config.yml")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()

	// var conf configs.Configs
	// decoder := yaml.NewDecoder(f)
	// err = decoder.Decode(&conf)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// log.withFields(log.Fields{
	// 	"state": state,
	// 	"port": conf.Server.Port,
	// 	"config": "",
	// }).Info("Starting Server")

	// ===================== init redis =============================
	// app.InitRedis(&conf)

	// ===================== Middleware =============================
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// ===================== init services =============================
	db := bnnd.NewBNNDRepo()
	svc := bnnd.NewBNNDManager(db)
	// ===================== connect db =============================

	//  ===================== define routes =============================
	shop := shop.NewShop(svc)
	e.GET("/v1/banana-ducks", shop.GetAllBananaDucks)
	e.POST("/v1/banana-ducks", shop.CreateBananaDuck)
	// e.GET("/v1/orders", shop.GetAllOrders)
	// e.POST("/v1/orders", shop.CreateOrder)

	e.Logger.Fatal(e.Start(":1234"))
}
