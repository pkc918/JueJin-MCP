package juejin

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

var (
	LOGIN_URL    = "https://juejin.cn/"
	LOGIN_BUTTON = `//*[@id="juejin"]/div[1]/div/header/div/nav/ul/ul/li[2]/div/button`
	USERAVATAR   = `#juejin > div.container.index-container > div > header > div > nav > ul > ul > li.nav-item.menu > div > div > img`
)

func Login(page *rod.Page, ctx context.Context) error {
	p := page.Context(ctx)
	p.MustNavigate(LOGIN_URL).MustWaitLoad()
	page.MustElementX(LOGIN_BUTTON).MustClick()

	time.Sleep(5 * time.Second)

	if exist, el, _ := p.Has(USERAVATAR); exist {
		fmt.Println(el.Describe(0, true))
		return nil
	}

	return errors.New("login failed")
}
