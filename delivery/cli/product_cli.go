package cli

import (
	"fmt"
	f "fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/enigma-mart/models"
	"github.com/enigma-mart/usecase"
	"github.com/enigma-mart/utils"
)

type ProductCli struct {
	ucProduct usecase.ProductUseCase
}

func (p *ProductCli) createNewProduct() error {
	fmt.Print("Masukan nama produk :")
	productName := utils.BufioInput().Text()

	fmt.Print("Masukan harga produk :")
	productPrice, _ := strconv.Atoi(utils.BufioInput().Text())

	fmt.Print("Masukan stok produk :")
	productStock, _ := strconv.Atoi(utils.BufioInput().Text())
	productId := utils.GenerateId()

	newProduct := models.Product{
		Id:    productId,
		Name:  productName,
		Price: productPrice,
		Stock: productStock,
	}

	err := p.ucProduct.CreateNewProduct(&newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductCli) listProduct() error {
	products, err := p.ucProduct.FindAllProduct()
	fmt.Println(products)
	if err != nil {
		return err
	}

	f.Println(strings.Repeat("-", 70))
	f.Println("\t\t\tENIGMA MART - DATA PRODUK")
	f.Println(strings.Repeat("-", 70))
	f.Println()
	thead := f.Sprintf("%s\t%3s\t%45s\t%15s\t\t%s\n", "NO", "KODE", "NAMA PRODUK", "HARGA", "STOCK")
	f.Println(thead)
	i := 0
	f.Println(strings.Repeat("-", 70))
	for _, each := range products {
		i++
		f.Printf("%v\t%-3v\t%-5v\t%15v\t\t%v\n", i, each.Id, each.Name, each.Price, each.Stock)
	}
	f.Println(strings.Repeat("-", 70))

	return nil
}

func NewProductCli(ucProduct usecase.ProductUseCase) *ProductCli {
	cli := ProductCli{
		ucProduct: ucProduct,
	}

	var menu string
	var isInt = regexp.MustCompile("^[0-9]+$")
	f.Println(strings.Repeat("*", 50))
	f.Println("\tMenu Produk Barang [Enigma Mart]")
	f.Println(strings.Repeat("*", 50))
	f.Println("   1. Tambah Produk")
	f.Println("   2. Hapus Produk")
	f.Println("   3. Detail Produk")
	f.Println("   4. Exit")
	f.Println(strings.Repeat("*", 50))

	// validasi menu
	for {
		f.Printf("Masukan menu yang dipilih (1-3): ")
		f.Scan(&menu)
		if isInt.MatchString(menu) {
			menu, _ := strconv.Atoi(menu)
			if menu > 0 && menu < 5 {
				break
			}
		}
		f.Println("Error: menu salah!")
	}
	switch menu {
	case "1":
		cli.createNewProduct()
	case "3":
		cli.listProduct()
	case "4":

	}

	return &cli
}
