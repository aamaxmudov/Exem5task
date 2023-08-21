package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Branch struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Products struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	CategorieId int     `json:"categorieId"`
}
type Transaction struct {
	Id         int    `json:"id"`
	Branch_id  int    `json:"branch_id"`
	Product_id int    `json:"product_id"`
	Type       string `json:"type"`
	Quantity   int    `json:"quantity"`
	Created_at string `json:"created_at"`
}
type BranchProducts struct {
	Branch_id  int `json:"branch_id"`
	Product_id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}
type Count struct {
	KirganSoni  int //1 Kun  string
	ChiqqanSoni int //1 Soni int
}

func main() {
	/*	//1-task (7-masala)
		dataTransaction, err := os.ReadFile("branchPrTransaction.json")
			if err != nil {
				fmt.Println("error reading transaction json", err)
				return
			}
			var transactions []Transaction
			err = json.Unmarshal(dataTransaction, &transactions)
			if err != nil {
				fmt.Println("error unmarshaling transactions", err)
			}

			dailyCounts := make(map[string]int)

			for _, transaction := range transactions {

				createdAt, err := time.Parse("2006-01-02 15:04:05", transaction.Created_at)
				if err != nil {
					fmt.Println("error parsing created_at", err)
					continue
				}

				kun := createdAt.Format("2006-01-02 15:04:05")

				dailyCounts[kun] += transaction.Quantity
			}

			var counts []Count
			for kun, soni := range dailyCounts {
				counts = append(counts, Count{Kun: kun, Soni: soni})
			}

			sort.Slice(counts, func(i, j int) bool {
				return counts[i].Soni > counts[j].Soni
			})

			fmt.Println("Har bir porductlar kamayish tartibida")
			for i, count := range counts {
				fmt.Printf("%2d. %10s %10d\n", i+1, count.Kun, count.Soni)
			}
	*/

	/*
		//2-task(8-masala)

		dataTransaction, err := os.ReadFile("branchPrTransaction.json")
		if err != nil {
			fmt.Println("error reading transaction json", err)
			return
		}
		var transactions []Transaction
		err = json.Unmarshal(dataTransaction, &transactions)
		if err != nil {
			fmt.Println("error unmarshaling transactions", err)
		}

		dataProduct, err := os.ReadFile("products.json")
		if err != nil {
			fmt.Println("error reading products json", err)
			return
		}

		var products []Products
		err = json.Unmarshal(dataProduct, &products)
		if err != nil {
			fmt.Println("error unmarshiling products", err)
		}

		productsPM := make(map[int]Count)

		for _, transaction := range transactions {
			if transaction.Type == "plus" {
				count := productsPM[transaction.Product_id]
				count.KirganSoni += transaction.Quantity
				productsPM[transaction.Product_id] = count
			} else if transaction.Type == "minus" {
				count := productsPM[transaction.Product_id]
				count.ChiqqanSoni += transaction.Quantity
				productsPM[transaction.Product_id] = count
			}
		}

		fmt.Println("Name       Kiritilgan   Chiqarilgan")
		for _, product := range products {
			count := productsPM[product.Id]
			fmt.Printf("%-17s %-12d %-10d\n", product.Name, count.KirganSoni, count.ChiqqanSoni)
		}
	*/

	//3-task(9-masala)

	dataBranch, err := os.ReadFile("branches.json")
	if err != nil {
		fmt.Println("error reading branches json", err)
		return
	}
	var branches []Branch

	err = json.Unmarshal(dataBranch, &branches)
	if err != nil {
		fmt.Println("error unmarshiling branches json", err)
		return
	}

	dataBranchProduct, err := os.ReadFile("branch_products.json")
	if err != nil {
		fmt.Println("error reading branch products json", err)
		return
	}
	var branchProduct []BranchProducts

	err = json.Unmarshal(dataBranchProduct, &branchProduct)
	if err != nil {
		fmt.Println("error unmarshiling branch products json", err)
		return
	}

	dataProduct, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println("error reading products json", err)
		return
	}

	var products []Products
	err = json.Unmarshal(dataProduct, &products)
	if err != nil {
		fmt.Println("error unmarshiling products", err)
	}

	branchProducts := make(map[int]Branch)
	for _, branch := range branches {
		branchProducts[branch.Id] = branch
	}

	branchTV := make(map[string]int)
	for _, bp := range branchProduct {
		productPrice := getProductPrice(bp.Product_id, products)
		branchName := branchProducts[bp.Branch_id].Name
		branchTV[branchName] += int(productPrice * float64(bp.Quantity))
	}

	fmt.Println("Filialda qancha summalik product borligi jadvali:")
	fmt.Printf("%-20s %s\n", "Filial", "Summa")
	for branchName, totalValue := range branchTV {
		fmt.Printf("%-20s %s\n", branchName, formatCurrency(totalValue))
	}
}

func getProductPrice(productID int, products []Products) float64 {
	for _, product := range products {
		if product.Id == productID {
			return product.Price
		}
	}
	return 0
}

func formatCurrency(amount int) string {
	return fmt.Sprintf("%d", amount)
}
