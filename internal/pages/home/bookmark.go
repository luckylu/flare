package home

import (
	"html/template"
	"strings"

	FlareData "github.com/luckylu/flare/internal/data"
	FlareModel "github.com/luckylu/flare/internal/model"
	FlareMDI "github.com/luckylu/flare/internal/resources/mdi"
	FlareState "github.com/luckylu/flare/internal/state"
)

func GenerateBookmarkTemplate(filter string) template.HTML {
	options := FlareData.GetAllSettingsOptions()
	bookmarksData := FlareData.LoadNormalBookmarks()
	tpl := ""

	var bookmarks []FlareModel.Bookmark

	if filter != "" {
		filterLower := strings.ToLower(filter)

		for _, bookmark := range bookmarksData.Items {
			if strings.Contains(strings.ToLower(bookmark.Name), filterLower) || strings.Contains(strings.ToLower(bookmark.URL), filterLower) {
				bookmarks = append(bookmarks, bookmark)
			}
		}
	} else {
		bookmarks = bookmarksData.Items
	}

	if len(bookmarksData.Categories) > 0 {
		defaultCategory := bookmarksData.Categories[0]
		for _, category := range bookmarksData.Categories {
			tpl += renderBookmarksWithCategories(&bookmarks, &category, &defaultCategory, options.OpenBookmarkNewTab, options.EnableEncryptedLink, options.IconMode)
		}
	} else {
		tpl += renderBookmarksWithoutCategories(&bookmarks, options.OpenBookmarkNewTab, options.EnableEncryptedLink, options.IconMode)
	}

	return template.HTML(tpl)
}

func renderBookmarksWithoutCategories(bookmarks *[]FlareModel.Bookmark, OpenBookmarkNewTab bool, EnableEncryptedLink bool, IconMode string) string {
	tpl := ""
	for _, bookmark := range *bookmarks {

		// 如果以 chrome-extension:// 协议开头
		// 则使用服务端 Location 方式打开链接
		templateURL := ""
		if strings.HasPrefix(bookmark.URL, "chrome-extension://") {
			templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(bookmark.URL)
		} else {
			if EnableEncryptedLink {
				templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(bookmark.URL)
			} else {
				templateURL = bookmark.URL
			}
		}

		templateIcon := ""
		if strings.HasPrefix(bookmark.Icon, "http://") || strings.HasPrefix(bookmark.Icon, "https://") {
			templateIcon = `<img src="` + bookmark.Icon + `"/>`
		} else if bookmark.Icon != "" {
			templateIcon = FlareMDI.GetIconByName(bookmark.Icon)
		} else {
			if IconMode == "FILLING" {
				templateIcon = FlareState.GetYandexFavicon(bookmark.URL, FlareMDI.GetIconByName(bookmark.Icon))
			} else {
				templateIcon = FlareMDI.GetIconByName(bookmark.Icon)
			}
		}

		if OpenBookmarkNewTab {
			tpl += `<li><a target="_blank" rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
		} else {
			tpl += `<li><a rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
		}
	}
	return `<div class="bookmark-group-container pull-left"><ul class="bookmark-list">` + tpl + `</ul></div>`
}

func renderBookmarksWithCategories(bookmarks *[]FlareModel.Bookmark, category *FlareModel.Category, defaultCategory *FlareModel.Category, OpenBookmarkNewTab bool, EnableEncryptedLink bool, IconMode string) string {
	tpl := ""
	isEmpty := true

	for _, bookmark := range *bookmarks {

		// 如果以 chrome-extension:// 协议开头
		// 则使用服务端 Location 方式打开链接
		templateURL := ""
		if strings.HasPrefix(bookmark.URL, "chrome-extension://") {
			templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(bookmark.URL)
		} else {
			if EnableEncryptedLink {
				templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(bookmark.URL)
			} else {
				templateURL = bookmark.URL
			}
		}

		templateIcon := ""
		if strings.HasPrefix(bookmark.Icon, "http://") || strings.HasPrefix(bookmark.Icon, "https://") {
			templateIcon = `<img src="` + bookmark.Icon + `"/>`
		} else if bookmark.Icon != "" {
			templateIcon = FlareMDI.GetIconByName(bookmark.Icon)
		} else {
			if IconMode == "FILLING" {
				templateIcon = FlareState.GetYandexFavicon(bookmark.URL, FlareMDI.GetIconByName(bookmark.Icon))
			} else {
				templateIcon = FlareMDI.GetIconByName(bookmark.Icon)
			}
		}

		if bookmark.Category != "" {
			if bookmark.Category == category.ID {
				if OpenBookmarkNewTab {
					tpl += `<li><a target="_blank" rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
				} else {
					tpl += `<li><a rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
				}
				isEmpty = false
			}
		} else {
			if category.ID == defaultCategory.ID {
				if OpenBookmarkNewTab {
					tpl += `<li><a target="_blank" rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
				} else {
					tpl += `<li><a rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
				}
				isEmpty = false
			}
		}
	}

	if isEmpty {
		return ``
	}

	return `<div class="bookmark-group-container pull-left"><h3 class="bookmark-group-title">` + category.Name + `</h3><ul class="bookmark-list">` + tpl + `</ul></div>`
}
