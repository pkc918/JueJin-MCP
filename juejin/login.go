package juejin

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func Login(page *rod.Page, ctx context.Context) error {
	p := page.Context(ctx)
	p.MustNavigate("https://juejin.cn/").MustWaitLoad()
	page.MustElementX(`//*[@id="juejin"]/div[1]/div/header/div/nav/ul/ul/li[2]/div/button`).MustClick()

	time.Sleep(5 * time.Second)

	if exist, el, _ := p.Has("#juejin > div.container.index-container > div > header > div > nav > ul > ul > li.nav-item.menu > div > div > img"); exist {
		fmt.Println(el.Describe(0, true))
		return nil
	}

	return errors.New("login failed")
}
