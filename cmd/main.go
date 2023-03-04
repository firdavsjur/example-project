package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsonDb"
	"log"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)

	// Product(c)
	controller.GetAllShopCart(c) //12.Shop cart boyicha default holati time sort DESC
	// controller.GetAllCategoryShopCars(c) //8.Qaysi category larda qancha mahsulot sotilgan boyicha jadval
	// controller.ActiveClient(c)           //9. Qaysi Client eng Active xaridor. Bitta ma'lumot chiqsa yetarli.
	// controller.Top10SellProducts(c) //5.Top 10 ta sotilayotgan mahsulotlarni royxati.
	// controller.Down10SellProducts(c) //6.Top 10 ta  eng past sotilayotgan mahsulotlarni royxati.
	// controller.MostSaledByDate(c) //7.Qaysi Sanada eng kop mahsulot sotilganligi boyicha jadval Count Sort DESC
	// fmt.Println("-------------------")
	// controller.DateFilterOfShopCarts("2022-02-02", "2022-02-04", c) //1. Shop cartlar Date boyicha filter qoyish kerak
	// Category(c)
	// controller.ClientHistory("Shokhrukh", c) //2. Client history chiqish kerak. Ya'ni sotib olgan mahsulotlari korsatish kerak
	// fmt.Println("-------------------")
	// controller.ClientTotalPrice("Shokhrukh", c) //3. Client qancha pul mahsulot sotib olganligi haqida hisobot.
	// controller.ProductSellCount("Shampoo", c)   //4. Productlarni Qancha sotilgan boyicha hisobot
}

// ShopCart

func Product(c *controller.Controller) {

	c.CreateProduct(&models.CreateProduct{
		Name:       "Smartfon vivo V25 8/256 GB",
		Price:      4_860_000,
		CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
	})
}

func Category(c *controller.Controller) {
	c.CreateCategory(&models.CreateCategory{
		Name:     "Boshqa",
		ParentID: "",
	})
}

func User(c *controller.Controller) {

	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
	err := c.MoneyTransfer(sender, receiver, 500_000)
	if err != nil {
		log.Println(err)
	}
}
