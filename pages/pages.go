package pages

import (
	"os"
	"time"

	"github.com/sclevine/agouti/core"
)

type NavigateFunc func(core.Page) (core.Page, error)

func (current NavigateFunc) Next(next NavigateFunc) NavigateFunc {
	return func(page core.Page) (core.Page, error) {
		p, err := current(page)
		time.Sleep(3 * time.Second)
		if err != nil {
			return p, err
		}
		return next(p)
	}
}

var GoToNetwork NavigateFunc = func(page core.Page) (core.Page, error) {
	err := page.Navigate("http://network.pivotal.io")
	return page, err
}

var GoToLoginPage NavigateFunc = func(page core.Page) (core.Page, error) {
	err := page.FirstByLink("Sign in").Click()
	if err != nil {
		return page, err
	}
	return page, err
}

var Login NavigateFunc = func(page core.Page) (core.Page, error) {
	err := page.Find("input[name='username']").Fill(os.Getenv("USERNAME"))
	if err != nil {
		return page, err
	}
	err = page.Find("input[name='password']").Fill(os.Getenv("PASSWORD"))
	if err != nil {
		return page, err
	}
	err = page.FindByButton("Sign in").Click()
	return page, err
}

var GoToPivotalCF NavigateFunc = func(page core.Page) (core.Page, error) {
	err := page.Find("a[href='/products/pivotal-cf']").Click()
	return page, err
}
