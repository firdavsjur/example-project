package controller

import (
	"app/models"
	"errors"
	"fmt"
	"log"
)

// getAllShopCart
func (c *Controller) GetAllShopCart(req *models.GetAllShopCartRequest) (models.GetAllShopCartsResponse, error) {
	data, err := c.store.ShopCart().GetAllShopCart(req)
	if err != nil {
		return models.GetAllShopCartsResponse{}, err
	}
	return data, nil
}

func (c *Controller) AddShopCart(req *models.Add) (string, error) {
	_, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.UserId})
	if err != nil {
		return "", err
	}

	_, err = c.store.Product().GetByID(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		return "", err
	}

	id, err := c.store.ShopCart().AddShopCart(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) RemoveShopCart(req *models.Remove) error {
	err := c.store.ShopCart().RemoveShopCart(req)
	if err != nil {
		return err
	}
	return err
}

func (c *Controller) CalculateTotal(req *models.UserPrimaryKey, status string, discount float64) (float64, error) {
	_, err := c.store.User().GetByID(req)
	if err != nil {
		return 0, err
	}

	users, err := c.store.ShopCart().GetUserShopCart(req)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, v := range users {
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return 0, err
		}
		if status == "fixed" {
			total += float64(v.Count) * (product.Price - discount)
		} else if status == "percent" {
			if discount < 0 || discount > 100 {
				return 0, errors.New("invalid discount range")
			}
			total += float64(v.Count) * (product.Price - (product.Price*discount)/100)
		} else {
			return 0, errors.New("invalid status name")
		}
	}

	if total < 0 {
		return 0, nil
	}
	return total, nil
}

//xxxxxxxxxxxxxxxxxxxxxxxxxx

func GetAllShopCart(c *Controller) {
	products, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}
	var results []models.ShopCartFull

	for i, v := range products.ShopCarts {
		getUser, err := c.GetByIdUser(&models.UserPrimaryKey{Id: v.UserId})
		if err != nil {
			log.Println(err)
			return
		}
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}

		results = append(results, models.ShopCartFull{
			Number:  i + 1,
			Name:    getUser.Name,
			Product: getProduct.Name,
			Price:   getProduct.Price,
			Total:   float64(v.Count) * getProduct.Price,
			Time:    v.Date,
		})

	}
	results = sortArrayofElement(results)
	if results != nil {

		for _, v := range results {
			fmt.Printf("%v. Name: %v, Product: %v Price: %v  Total: %v, Time: %v\n", v.Number, v.Name, v.Product, v.Price, v.Total, v.Time)
		}
	} else {
		fmt.Println("Malumot chiqmadi")
	}
}

func ClientHistory(clientName string, c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	var results []models.ShopCartFull
	for i, v := range shopCarts.ShopCarts {
		getUser, err := c.GetByIdUser(&models.UserPrimaryKey{Id: v.UserId})
		if err != nil {
			log.Println("Klient topilmadi")
			return
		}
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println("Mahsulot topilmadi")
			return
		}
		if clientName == getUser.Name {
			results = append(results, models.ShopCartFull{
				Number:  i + 1,
				Name:    getUser.Name,
				Product: getProduct.Name,
				Price:   getProduct.Price,
				Total:   float64(v.Count) * getProduct.Price,
				Time:    v.Date,
			})
		}
	}
	if results != nil {
		for _, v := range results {
			fmt.Printf("%v. Name: %v, Product: %v Price: %v  Total: %v, Time: %v\n", v.Number, v.Name, v.Product, v.Price, v.Total, v.Time)
		}
	} else {
		fmt.Println("Malumot chiqmadi")
	}

}

func ClientTotalPrice(clientName string, c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	var totalProce float64
	var totalCount int
	var theCheapestProductPrice []float64
	for _, v := range shopCarts.ShopCarts {
		getUser, err := c.GetByIdUser(&models.UserPrimaryKey{Id: v.UserId})
		if err != nil {
			log.Println(err)
			return
		}
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}
		theCheapestProductPrice = append(theCheapestProductPrice, getProduct.Price)
		if clientName == getUser.Name {
			totalProce += getProduct.Price * float64(v.Count)
			totalCount += (v.Count)
		}
	}
	sortArrayofFloat(theCheapestProductPrice)
	if totalProce != 0 {
		fmt.Printf("Count: %v\n", totalCount)
		if totalCount > 10 {
			println(totalProce)

			totalProce -= theCheapestProductPrice[0]

			println(totalProce)
			fmt.Printf("Name: %v, Total Buy Price: %v\n", clientName, totalProce)

		} else {
			fmt.Printf("Name: %v, Total Buy Price: %v\n", clientName, totalProce)

		}

	} else {
		fmt.Println("Malumot chiqmadi")
	}

}

func ProductSellCount(productName string, c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	var count int
	for _, v := range shopCarts.ShopCarts {
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}
		if productName == getProduct.Name {
			count += (v.Count)
		}
	}
	if count != 0 {
		fmt.Printf("Name: %v, Total: %v\n", productName, count)

	} else {
		fmt.Println("Malumot chiqmadi")
	}

}

func Top10SellProducts(c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}
	var results []models.ClientTotalPrice
	for i, v := range shopCarts.ShopCarts {
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}
		atvet, index := ifNameInFunction(getProduct.Name, results)
		if atvet {
			results[index].Total += float64(v.Count)
		} else {
			results = append(results, models.ClientTotalPrice{
				Number: fmt.Sprint(i + 1),
				Name:   getProduct.Name,
				Total:  float64(v.Count),
			})
		}

	}

	results = sortArrayofElement1(results)
	for i := 0; i < len(results); i++ {
		fmt.Printf("Name: %v Count: %v\n", results[i].Name, results[i].Total)
	}
}
func Down10SellProducts(c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}
	var results []models.ClientTotalPrice
	for i, v := range shopCarts.ShopCarts {
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}
		atvet, index := ifNameInFunction(getProduct.Name, results)
		if atvet {
			results[index].Total += float64(v.Count)
		} else {
			results = append(results, models.ClientTotalPrice{
				Number: fmt.Sprint(i + 1),
				Name:   getProduct.Name,
				Total:  float64(v.Count),
			})
		}

	}

	results = desortArrayofElement1(results)
	for i := 0; i < len(results); i++ {
		fmt.Printf("Name: %v Count: %v\n", results[i].Name, results[i].Total)
	}
}

func MostSaledByDate(c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}
	var results []models.ClientTotalPrice
	for i, v := range shopCarts.ShopCarts {
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}
		atvet, index := ifDateInFunction(v.Date, results)
		if atvet {
			results[index].Total += float64(v.Count)
		} else {
			results = append(results, models.ClientTotalPrice{
				Number: fmt.Sprint(i + 1),
				Name:   getProduct.Name,
				Total:  float64(v.Count),
				Date:   v.Date,
			})
		}

	}

	results = sortArrayofElement1(results)
	for i := 0; i < len(results); i++ {
		fmt.Printf("Name: %v Date: %v Count: %v\n", results[i].Name, results[i].Date, results[i].Total)
	}
}

func ActiveClient(c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}
	var results []models.ClientTotalPrice
	for i, v := range shopCarts.ShopCarts {
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}
		getUser, err := c.GetByIdUser(&models.UserPrimaryKey{Id: v.UserId})
		if err != nil {
			log.Println(err)
			return
		}
		atvet, index := ifNameInFunction(getUser.Name, results)
		if atvet {
			results[index].Total += float64(v.Count) * getProduct.Price
		} else {
			results = append(results, models.ClientTotalPrice{
				Number: fmt.Sprint(i + 1),
				Name:   getUser.Name,
				Total:  float64(v.Count) * getProduct.Price,
			})
		}

	}

	results = sortArrayofElement1(results)
	fmt.Println(results[0])

}

func GetAllCategoryShopCars(c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}
	var results []models.CategoryCountList
	for _, v := range shopCarts.ShopCarts {
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}
		getCategory, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: getProduct.Category.Id})
		if err != nil {
			log.Println(err)
			return
		}
		atvet, index := ifInFunction(getCategory.Name, results)
		if atvet {
			results[index].Count += v.Count
		} else {
			results = append(results, models.CategoryCountList{
				Name:  getCategory.Name,
				Count: v.Count,
			})
		}

	}
	for _, v := range results {
		fmt.Printf("Name: %v Count: %v\n", v.Name, v.Count)
	}

}
func ifDateInFunction(date string, result []models.ClientTotalPrice) (bool, int) {
	for i, v := range result {
		if v.Date == date {
			return true, i
		}
	}
	return false, 0
}
func ifNameInFunction(name string, result []models.ClientTotalPrice) (bool, int) {
	for i, v := range result {
		if v.Name == name {
			return true, i
		}
	}
	return false, 0
}
func ifInFunction(name string, result []models.CategoryCountList) (bool, int) {
	for i, v := range result {
		if v.Name == name {
			return true, i
		}
	}
	return false, 0
}
func DateFilterOfShopCarts(from_date, to_date string, c *Controller) {
	shopCarts, err := c.GetAllShopCart(
		&models.GetAllShopCartRequest{
			Offset: 0,
			Limit:  10,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	var results []models.ShopCartFull
	for i, v := range shopCarts.ShopCarts {
		getUser, err := c.GetByIdUser(&models.UserPrimaryKey{Id: v.UserId})
		if err != nil {
			log.Println(err, "salom")
			return
		}
		getProduct, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			log.Println(err)
			return
		}
		if v.Date >= from_date && v.Date <= to_date {
			results = append(results, models.ShopCartFull{
				Number:  i + 1,
				Name:    getUser.Name,
				Product: getProduct.Name,
				Price:   getProduct.Price,
				Total:   float64(v.Count) * getProduct.Price,
				Time:    v.Date,
			})
		}
	}
	for _, v := range results {
		fmt.Printf("%v. Name: %v, Product: %v Price: %v  Total: %v, Time: %v\n", v.Number, v.Name, v.Product, v.Price, v.Total, v.Time)
	}

}

func sortArrayofElement(arr []models.ShopCartFull) []models.ShopCartFull {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i].Time <= arr[j].Time {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				continue
			}
		}
	}

	return arr
}

func sortArrayofFloat(arr []float64) []float64 {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] >= arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				continue
			}
		}
	}

	return arr
}

func sortArrayofElement1(arr []models.ClientTotalPrice) []models.ClientTotalPrice {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i].Total <= arr[j].Total {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				continue
			}
		}
	}

	return arr
}
func desortArrayofElement1(arr []models.ClientTotalPrice) []models.ClientTotalPrice {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i].Total >= arr[j].Total {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				continue
			}
		}
	}

	return arr
}
