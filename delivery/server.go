package delivery

import (
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	f "fmt"

	"github.com/enigma-mart/config"
	"github.com/enigma-mart/delivery/cli"
	"github.com/enigma-mart/manager"
)

type appServer struct {
	ucManager manager.UseCaseManager
}

func Server() *appServer {
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repoManager := manager.NewRepositoryManager(infra)
	ucManager := manager.NewUseCaseManager(repoManager)
	return &appServer{
		ucManager: ucManager,
	}
}

func (a *appServer) initCli() {
	var menu string
	var isInt = regexp.MustCompile("^[0-9]+$")
	// var isString = regexp.MustCompile(`[a-z]+`)
	f.Println(strings.Repeat("*", 40))
	f.Println("\tMain Menu [Enigma Mart]")
	f.Println(strings.Repeat("*", 40))
	f.Println("   1. Produk Barang")
	f.Println("   2. Transaksi Penjualan")
	f.Println("   3. Laporan Penjualan")
	f.Println("   4. Logout")
	f.Println(strings.Repeat("*", 40))
	for {
		f.Printf("Masukan menu yang dipilih (1-4): ")
		f.Scan(&menu)
		if isInt.MatchString(menu) {
			menu, _ := strconv.Atoi(menu)
			if menu > 0 && menu < 6 {
				break
			}
		}
		f.Println("Error: menu salah!")
		clearScreen()
	}

	switch menu {
	case "1":
		cli.NewProductCli(a.ucManager.ProductUseCase())
	case "4":
		f.Println("Logout Success!!")
		os.Exit(1)
	}
}

func (a *appServer) Run() {
	a.initCli()
}

func clearScreen() {
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osRunning == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
