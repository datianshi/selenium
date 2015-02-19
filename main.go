package main

import (
	"fmt"
	. "github.com/datianshi/pivotal-download/pages"
	agouti "github.com/sclevine/agouti/core"
)

func main() {
	err := drive()
	if err != nil {
		fmt.Println(err)
	}
}

func drive() (err error) {
	driver, err := agouti.Selenium()
	if err != nil {
		return
	}
	err = driver.Start()
	defer driver.Stop()
	if err != nil {
		return
	}
	page, err := driver.Page(agouti.Use().Browser("chrome"))
	if err != nil {
		return
	}
	nav := GoToNetwork.Next(GoToLoginPage).Next(Login)
	_, err = nav(page)
	return
}
