package juejin

import (
	"context"
	"fmt"

	"github.com/go-rod/rod"
)

const (
	CATEGORY_LIST = `//*[@id="juejin-web-editor"]/div[2]/div/header/div[2]/div[3]/div/div[2]/div[2]`
)

var CategoryMap = map[int]string{
	0: "后端",
	1: "前端",
	2: "Android",
	3: "iOS",
	4: "人工智能",
	5: "开发工具",
	6: "代码人生",
	7: "阅读",
}

func SelectorCategoryItem(page *rod.Page, _ctx context.Context, categoryIndex int) error {
	// 1. 等待分类列表容器出现并可见
	categoryList := page.MustElementX(CATEGORY_LIST)
	categoryList.MustWaitVisible()
	categoryList.MustWaitStable()

	// 2. 直接获取分类项子元素
	items := categoryList.MustElements(".item")

	if len(items) == 0 {
		return fmt.Errorf("分类列表为空，无法选择分类")
	}

	// 3. 检查索引是否越界
	if categoryIndex < 0 || categoryIndex >= len(items) {
		return fmt.Errorf("分类索引 %d 越界，可用范围: 0-%d", categoryIndex, len(items)-1)
	}

	// 4. 等待目标分类项可见并稳定，然后点击
	targetItem := items[categoryIndex]
	targetItem.MustWaitVisible()
	targetItem.MustWaitStable()
	targetItem.MustClick()

	return nil
}
