package core

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func sendStartMessage(c tele.Context) error {
	message := `
Hi! I'm moe_sticker_bot!

This bot originated in the <a href="https://github.com/star-39/moe-sticker-bot">original project</a>.
The community-maintained fork is <a href="https://github.com/Nununene/moe-sticker-bot">here</a>.

Please:
• Send <b>LINE/Kakao sticker share link</b> to import or download.
• Send <b>Telegram sticker/link/GIF</b> to download or export to WhatsApp.
• Send <b>keywords</b> to search sticker sets.
• Tap to <b>/create</b> or <b>/manage</b> sticker set and CustomEmoji.
• Tap to check all available <b>/command_list</b>.

你好！我是萌萌貼圖 BOT。
本機器人源自<a href="https://github.com/star-39/moe-sticker-bot">原始專案</a>，目前由社群維護的版本在<a href="https://github.com/Nununene/moe-sticker-bot">這裡</a>。

你可以：
• 傳送<b>LINE/kakao貼圖包的分享連結</b>來匯入或下載.
• 傳送<b>Telegram貼圖/連結/GIF</b>來下載.
• 傳送<b>關鍵字</b>來搜尋貼圖包.
• 傳送 <b>/create</b> 或 <b>/manage</b> 來創建或管理貼圖包和表情貼。
• 傳送 <b>/command_list</b> 檢視所有可用指令.
`
	return c.Send(message, tele.ModeHTML, tele.NoPreview)
}

func sendCommandList(c tele.Context) error {
	message := `
<b>/import</b>  <b>/search</b> LINE/Kakao stickers.<code>
匯入或搜尋LINE/Kaka貼圖包.</code>
<b>/download</b>  <b>/create</b>  <b>/manage</b> Telegram stickers.<code>
下載、創建、管理Telegram貼圖包.</code>
<b>/faq  /about  /changelog  /privacy</b><code>
常見問題/關於/更新紀錄/私隱</code>
`

	return c.Send(message, tele.ModeHTML, tele.NoPreview)
}

func sendAboutMessage(c tele.Context) {
	c.Send(fmt.Sprintf(`
<b>About moe_sticker_bot / 關於萌萌貼圖 BOT</b>

Original project: <a href="https://github.com/star-39/moe-sticker-bot">star-39/moe-sticker-bot</a>
Community-maintained version: <a href="https://github.com/Nununene/moe-sticker-bot">Nununene/moe-sticker-bot</a>

This is a community-maintained fork of the original project.
本專案是原始專案的社群維護版本。

<code>
This free software is released under the GPLv3 License.
It comes with absolutely no warranty.
本 BOT 依 GPLv3 授權發布，並不提供任何保證。
</code><b>
Please send /start to start using
請傳送 /start 來開始
</b><code>
Version:版本: %s
</code>
`, BOT_VERSION), tele.ModeHTML)
}

func sendFAQ(c tele.Context) {
	c.Send(fmt.Sprintf(`
<b>FAQ / 常見問題</b>
Original project: <a href="https://github.com/star-39/moe-sticker-bot">star-39/moe-sticker-bot</a>
Community-maintained version: <a href="https://github.com/Nununene/moe-sticker-bot">Nununene/moe-sticker-bot</a>

<b>Q: I'm stuck and can't exit a command. What should I do?
我卡在操作中，無法退出指令，該怎麼辦？</b>
A: Send /quit to cancel the current operation.
請傳送 /quit 中斷目前操作。

<b>Q: Why ID has suffix: _by_%s ?
為甚麼ID的末尾有: _by_%s ?</b>
A: Telegram requires bots to use this suffix when creating sticker sets.
這是 Telegram 對機器人建立貼圖包的命名要求。

<b>Q: Who owns the sticker sets the bot created?
    BOT創造的貼圖包由誰所有?</b>
A: It's you of course. You can manage them through /manage or Telegram's official @Stickers bot.
    當然是您. 您可以通過 /manage 指令或者Telegram官方的 @Stickers 管理您的貼圖包.

<b>Q: Where can I find the community-maintained source code?
要在哪裡查看社群維護版的程式碼？</b>
A: Visit <a href="https://github.com/Nununene/moe-sticker-bot">Nununene/moe-sticker-bot</a>. The original project is <a href="https://github.com/star-39/moe-sticker-bot">star-39/moe-sticker-bot</a>.
請前往 <a href="https://github.com/Nununene/moe-sticker-bot">社群維護版 GitHub</a>；原始專案也保留在 <a href="https://github.com/star-39/moe-sticker-bot">這裡</a>。
`, botName, botName), tele.ModeHTML)
}

func sendChangelog(c tele.Context) error {
	return c.Send(`

Community-maintained version updates / 社群維護版本更新：
Original project: https://github.com/star-39/moe-sticker-bot
Community-maintained version: https://github.com/Nununene/moe-sticker-bot

2026-11-07
* Fix Telegram WebApp initData validation.
* Add a Dockerfile to build the bot from the community-maintained source.
* 修正 Telegram WebApp initData 驗證。
* 新增從社群維護版原始碼建置 Bot 的 Dockerfile。

2026-11-08
* Parameterize sticker search and remove bot tokens and database credentials from logs.
* Register Telegram bot commands.
* Improve import progress messages, clean up automatic import work directories, and improve sticker conversion quality.
* Validate and limit downloads; report static WEBP conversion failures.
* Fix per-sticker emoji assignment by installing its runtime dependency; show the title prompt after preparation finishes.
* Refresh bot help/about messages and document both project repositories.
* 將貼圖搜尋改為參數化查詢，並避免在日誌中記錄 Bot Token 與資料庫密碼。
* 啟動時註冊 Telegram 機器人指令。
* 改善匯入進度訊息、清理自動匯入暫存目錄並提升貼圖轉檔畫質。
* 加入下載驗證與大小限制，並回報靜態 WEBP 轉檔錯誤。
* 補上個別 emoji 設定所需的 runtime 套件，並在貼圖準備完成後才顯示標題選項。
* 更新 Bot 說明訊息，並並列原始專案與社群維護版連結。
	`, tele.NoPreview)
}

func sendPrivacy(c tele.Context) error {
	return c.Send(`
<b>Privacy Notice:</b>
None of your usage or behaviour will be stored or analyzed.
None of your user identifier or information will be collected or stored if you did not use /import or /create command and succeded.

If you used /create or /import feature and upon success,
your user identifier will be associated to the sticker set you create and stored to database on the bot server,
which will only be used to tell which sticker set is owned by you only when you use /manage command.
No one else could see or use the stored user identifier.

All the data being stored is encrypted.
This bot will never share any of the stored data to anyone or to anywhere else.
The bot server is physically located at Osaka,Japan. Local laws might apply.
This bot is free and open source software. The database implementation is available in the original project at
https://github.com/star-39/moe-sticker-bot/blob/master/core/database.go and in the community-maintained version at
https://github.com/Nununene/moe-sticker-bot/blob/master/core/database.go.
Review the source to see how the bot stores and processes data.

<b>私隱聲明:</b>
本bot不會存儲或分析您的使用情況或行為。
本bot不會採集或儲存任何用戶資訊，除非您使用了 /import 或 /create 指令且成功完成。

如果您使用了 /import 或 /create 指令並且成功完成，
您的用戶識別子(user_id)會與您創建的貼圖包關聯並存檔入bot伺服器的資料庫中。
此識別子只會用來讓您本人通過 /manage 指令查詢您所擁有的貼圖包，不作其他任何用途。
任何其他用戶無法看見或使用此識別子。

本bot儲存的所有資訊均經過加密。
本bot不會分享任何儲存的資訊給任何人或實體或到任何地方。
本bot伺服器的物理位置位於日本大阪。 當地法律可能適用。
本 bot 為自由開放原始碼軟體。資料庫實作可參閱原始專案
https://github.com/star-39/moe-sticker-bot/blob/master/core/database.go，或社群維護版
https://github.com/Nununene/moe-sticker-bot/blob/master/core/database.go，了解 bot 如何儲存和處理資料。
`, tele.ModeHTML, tele.NoPreview)
}

func sendAskEmoji(c tele.Context) error {
	selector := &tele.ReplyMarkup{}
	btnManu := selector.Data("Assign separately/分別設定", "manual")
	btnRand := selector.Data(`Batch assign as/一併設定為 "⭐"`, "random")
	selector.Inline(selector.Row(btnManu), selector.Row(btnRand))

	return c.Send(`
Telegram requires emoji to and keywords for each sticker:
• Press "Assign separately" to assign emoji and keywords one by one.
• Send an emoji to do batch assign.
Telegram要求為每張貼圖分別設定emoji和關鍵字:
• 按下"分別設定"來為每個貼圖分別設定相應的emoji和關鍵字.
• 傳送一個emoji來為全部貼圖設定成一樣的.
`, selector)
}

func sendConfirmExportToWA(c tele.Context, sn string, hex string) error {
	selector := &tele.ReplyMarkup{}
	baseUrl, _ := url.JoinPath(msbconf.WebappUrl, "export")
	webAppUrl := fmt.Sprintf("%s?sn=%s&hex=%s", baseUrl, sn, hex)
	log.Debugln("webapp export link is:", webAppUrl)
	webapp := tele.WebApp{URL: webAppUrl}
	btnExport := selector.WebApp("Continue export/繼續匯出 →", &webapp)
	selector.Inline(selector.Row(btnExport))

	return c.Reply(`
Exporting to WhatsApp requires <a href="https://github.com/star-39/msb_app">Msb App</a> due to their restrictions, then press "Continue export".
匯出到WhatsApp需要手機上安裝<a href="https://github.com/star-39/msb_app">Msb App</a>, 然後按下"繼續匯出".

Download:下載:
<b>iPhone:</b> AppStore(N/A.暫無), <a href="https://github.com/star-39/msb_app/releases/latest/download/msb_app.ipa">IPA</a>
<b>Android:</b> GooglePlay(N/A.暫無), <a href="https://github.com/star-39/msb_app/releases/latest/download/msb_app.apk">APK</a>
`, tele.ModeHTML, tele.NoPreview, selector)
}

func genSDnMnEInline(canManage bool, isTGS bool, sn string) *tele.ReplyMarkup {
	selector := &tele.ReplyMarkup{}
	btnSingle := selector.Data("Download this sticker/下載這張貼圖", CB_DN_SINGLE)
	btnAll := selector.Data("Download sticker set/下載整個貼圖包", CB_DN_WHOLE)
	btnMan := selector.Data("Manage sticker set/管理這個貼圖包", CB_MANAGE)
	// btnExport := selector.Data("Export to WhatsApp/匯出到WhatsApp", CB_EXPORT_WA)
	if canManage {
		selector.Inline(selector.Row(btnSingle), selector.Row(btnAll), selector.Row(btnMan))
	} else {
		if isTGS {
			//If is TGS, do not support export to WA.
			selector.Inline(selector.Row(btnSingle), selector.Row(btnAll))
		} else {
			selector.Inline(selector.Row(btnSingle), selector.Row(btnAll))
		}
	}
	return selector
}

func sendAskSDownloadChoice(c tele.Context, s *tele.Sticker) error {
	selector := genSDnMnEInline(false, s.Animated, s.SetName)
	return c.Reply(`
You can download this sticker or the whole sticker set, please select below.
您可以下載這個貼圖或者其所屬的整個貼圖包, 請選擇:
`, selector)
}

func sendAskSChoice(c tele.Context, sn string) error {
	selector := genSDnMnEInline(true, false, sn)
	return c.Reply(`
You own this sticker set. You can download or manage this sticker set, please select below.
您擁有這個貼圖包. 您可以下載或者管理這個貼圖包, 請選擇:
`, selector)
}

func sendAskTGLinkChoice(c tele.Context) error {
	selector := &tele.ReplyMarkup{}
	btnManu := selector.Data("Download sticker set/下載整個貼圖包", CB_DN_WHOLE)
	btnMan := selector.Data("Manage sticker set/管理這個貼圖包", CB_MANAGE)
	selector.Inline(selector.Row(btnManu), selector.Row(btnMan))
	return c.Reply(`
You own this sticker set. You can download or manage this sticker set, please select below.
您擁有這個貼圖包. 您可以下載或者管理這個貼圖包, 請選擇:
`, selector)
}

func sendAskWantSDown(c tele.Context) error {
	selector := &tele.ReplyMarkup{}
	btn1 := selector.Data("Yes", CB_DN_WHOLE)
	btnNo := selector.Data("No", CB_BYE)
	selector.Inline(selector.Row(btn1), selector.Row(btnNo))
	return c.Reply(`
You can download this sticker set. Press Yes to continue.
您可以下載這個貼圖包, 按下Yes來繼續.
`, selector)
}

func sendAskWantImportOrDownload(c tele.Context, avalAsEmoji bool) error {
	msg := ""
	selector := &tele.ReplyMarkup{}
	btnImportSticker := selector.Data("Import as sticker set/作為普通貼圖包匯入", CB_OK_IMPORT)
	btnImportEmoji := selector.Data("Import as CustomEmoji/作為表情貼匯入", CB_OK_IMPORT_EMOJI)
	btnDownload := selector.Data("Download/下載", CB_OK_DN)
	if avalAsEmoji {
		selector.Inline(selector.Row(btnImportSticker), selector.Row(btnImportEmoji), selector.Row(btnDownload))
		msg = `
You can import this sticker set to Telegram or download it.
Import as Custom Emoji is also available, however you will need Telegram Premium to send them.
您可以下載或匯入這個貼圖包到Telegram.
也可以作為表情貼匯入，但是傳送需要Telegram會員。`
	} else {
		selector.Inline(selector.Row(btnImportSticker), selector.Row(btnDownload))
		msg = `
You can import this sticker set to Telegram or download it.
您可以下載或匯入這個貼圖包到Telegram.`
	}

	return c.Reply(msg, selector)
}

func sendAskWhatToDownload(c tele.Context) error {
	return c.Send("Please send a sticker that you want to download, or its share link(can be either Telegram or LINE ones)\n" +
		"請傳送想要下載的貼圖, 或者是貼圖包的分享連結(可以是Telegram或LINE連結).")
}

func sendAskTitle_Import(c tele.Context) error {
	ld := users.data[c.Sender().ID].lineData
	ld.TitleWg.Wait()
	log.Debug("titles are::")
	log.Debugln(ld.I18nTitles)
	selector := &tele.ReplyMarkup{}

	var titleButtons []tele.Row
	var titleText string
	for i, t := range ld.I18nTitles {
		if t == "" {
			continue
		}
		title := escapeTagMark(t) + " @" + botName
		btn := selector.Data(title, strconv.Itoa(i))
		row := selector.Row(btn)
		titleButtons = append(titleButtons, row)
		titleText = titleText + "\n<code>" + title + "</code>"
	}

	if len(titleButtons) == 0 {
		btnDefault := selector.Data(escapeTagMark(ld.Title)+" @"+botName, CB_DEFAULT_TITLE)
		titleButtons = []tele.Row{selector.Row(btnDefault)}
	}
	selector.Inline(titleButtons...)

	return c.Send("Please send a title for this sticker set. You can also select an original title below:\n"+
		"請傳送貼圖包的標題.您也可以按下面的按鈕自動填上合適的原版標題:\n"+
		titleText, selector, tele.ModeHTML)
}

func sendAskTitle(c tele.Context) error {
	return c.Send("Please send a title for this sticker set.\n" +
		"請傳送貼圖包的標題.")
}

func sendAskID(c tele.Context) error {
	selector := &tele.ReplyMarkup{}
	btnAuto := selector.Data("Auto Generate/自動生成", "auto")
	selector.Inline(selector.Row(btnAuto))
	return c.Send(`
Please send an ID for sticker set, used in share link.
Can contain alphanum and underscore only.
請設定貼圖包的ID, 用於分享連結.
只可以含有英文,數字,下劃線.
For example: 例如:
<code>My_favSticker21</code>

ID is usually not important, you can press Auto Generate.
ID通常不重要, 您可以按下"自動生成".`, selector, tele.ModeHTML)
}

func sendAskImportLink(c tele.Context) error {
	return c.Send(`
Please send LINE/kakao store link of the sticker set. You can obtain this link from App by going to sticker store and tapping Share->Copy Link.
請傳送貼圖包的LINE/kakao Store連結. 您可以在App裡的貼圖商店按右上角的分享->複製連結來取得連結.
For example: 例如:
<code>https://store.line.me/stickershop/product/7673/ja</code>
<code>https://e.kakao.com/t/pretty-all-friends</code>
<code>https://emoticon.kakao.com/items/lV6K2fWmU7CpXlHcP9-ysQJx9rg=?referer=share_link</code>
`, tele.ModeHTML)
}

func sendNotifySExist(c tele.Context, lineID string) bool {
	lines := queryLineS(lineID)
	if len(lines) == 0 {
		return false
	}
	message := "This sticker set exists in our database, you can continue import or just use them if you want.\n" +
		"此套貼圖包已經存在於資料庫中, 您可以繼續匯入, 或者使用下列現成的貼圖包\n\n"

	var entries []string
	for _, l := range lines {
		if l.Ae {
			entries = append(entries, fmt.Sprintf(`<a href="%s">%s</a>`, "https://t.me/addstickers/"+l.Tg_id, l.Tg_title))
		} else {
			// append to top
			entries = append([]string{fmt.Sprintf(`★ <a href="%s">%s</a>`, "https://t.me/addstickers/"+l.Tg_id, l.Tg_title)}, entries...)
		}
	}
	if len(entries) > 5 {
		entries = entries[:5]
	}
	message += strings.Join(entries, "\n")
	c.Send(message, tele.ModeHTML)
	return true
}

func sendSearchResult(entriesWant int, lines []LineStickerQ, c tele.Context) error {
	var entries []string
	message := "Search Results: 搜尋結果：\n"

	for _, l := range lines {
		l.Tg_title = strings.TrimSuffix(l.Tg_title, " @"+botName)
		if l.Ae {
			entries = append(entries, fmt.Sprintf(`<a href="%s">%s</a>`, "https://t.me/addstickers/"+l.Tg_id, l.Tg_title))
		} else {
			// append to top
			entries = append([]string{fmt.Sprintf(`★ <a href="%s">%s</a>`, "https://t.me/addstickers/"+l.Tg_id, l.Tg_title)}, entries...)
		}
	}

	if entriesWant == -1 && len(entries) > 120 {
		c.Send("Too many results, please narrow your keyword, truncated to 120 entries.\n" +
			"搜尋結果過多，已縮減到120個，請使用更準確的搜尋關鍵字。")
		entries = entries[:120]
	}
	if entriesWant != -1 && len(entries) > entriesWant {
		entries = entries[:entriesWant]
	}
	if len(entries) > 30 {
		eChunks := chunkSlice(entries, 30)
		for _, eChunk := range eChunks {
			msgToSend := message + strings.Join(eChunk, "\n")
			c.Send(msgToSend, tele.ModeHTML)
		}
	} else {
		message += strings.Join(entries, "\n")
		c.Send(message, tele.ModeHTML)
	}

	return nil
}

func sendAskStickerFile(c tele.Context) error {
	return c.Send("Please send images/photos/stickers(less than 120 in total),\n" +
		"or send an archive containing image files,\n" +
		"wait until upload complete, then tap 'Done adding'.\n\n" +
		"請傳送任意格式的圖片/影片/貼圖(少於120張)\n" +
		"或者傳送內有貼圖檔案的歸檔,\n" +
		"等候所有檔案上載完成, 然後按下「停止增添」\n")
}

func sendInStateWarning(c tele.Context) error {
	command := users.data[c.Sender().ID].command
	state := users.data[c.Sender().ID].state

	return c.Send(fmt.Sprintf(`
Please send content according to instructions.
請按照bot提示傳送相應內容.
Current command: %s
Current state: %s
You can also send /quit to terminate session.
您也可以傳送 /quit 來中斷對話.
`, command, state))
}

func sendNoSessionWarning(c tele.Context) error {
	return c.Send("Please use /start or send LINE/kakao/Telegram links or stickers.\n請使用 /start 或者傳送LINE/kakao/Telegram連結或貼圖.")
}

func sendAskSTypeToCreate(c tele.Context) error {
	selector := &tele.ReplyMarkup{}
	btnRegular := selector.Data("Regular sticker set/普通貼圖包", CB_REGULAR_STICKER)
	btnCustomEmoji := selector.Data("Custom Emoji/表情貼", CB_CUSTOM_EMOJI)

	selector.Inline(selector.Row(btnRegular), selector.Row(btnCustomEmoji))
	return c.Send("What kind of sticker set you want to create?\nNote that custom emoji can only be sent by Telegram Premium member."+
		"您想要創建何種類型的貼圖包?\n請注意只有Telgram會員可以傳送表情貼。", selector)
}

func sendAskEmojiAssign(c tele.Context) error {
	sd := users.data[c.Sender().ID].stickerData
	sf := sd.stickers[sd.pos]
	sf.wg.Wait()
	caption := fmt.Sprintf(`
Send emoji(s) representing this sticker.
請傳送代表這個貼圖的emoji(可以多個).

%d of %d
`, sd.pos+1, sd.lAmount)

	if sf.fileID != "" {
		msg, _ := c.Bot().Send(c.Sender(), &tele.Sticker{
			File: tele.File{FileID: sf.fileID},
		})
		_, err := c.Bot().Reply(msg, caption)
		return err
	}

	err := c.Send(&tele.Video{
		File:    tele.FromDisk(sf.oPath),
		Caption: caption,
	})
	if err != nil {
		err2 := c.Send(&tele.Video{
			File:    tele.FromDisk(sd.stickers[sd.pos].oPath),
			Caption: caption,
		})
		if err2 != nil {
			err3 := c.Send(&tele.Document{
				File:     tele.FromDisk(sd.stickers[sd.pos].oPath),
				FileName: filepath.Base(sd.stickers[sd.pos].oPath),
				Caption:  caption,
			})
			if err3 != nil {
				err4 := c.Send(&tele.Sticker{File: tele.File{FileID: sd.stickers[sd.pos].oPath}})
				if err4 != nil {
					return err4
				}
			}
		}
	}
	return nil
}

func sendFatalError(err error, c tele.Context) {
	if c == nil {
		return
	}
	var errMsg string
	if err != nil {
		errMsg = err.Error()
		errMsg = strings.ReplaceAll(errMsg, msbconf.BotToken, "***")
		if strings.Contains(errMsg, "500") {
			errMsg += "\nThis is an internal error of Telegram server, we could do nothing but wait for its recover. Please try again later.\n" +
				"此錯誤為Telegram伺服器之內部錯誤, 無法由bot解決, 只能等候官方修復. 建議您稍後再嘗試一次.\n"
		}
	}

	c.Send("<b>Fatal error encounterd. Please try again. /start\n"+
		"發生嚴重錯誤. 請您從頭再試一次. /start </b>\n\n"+
		"Original project issues: https://github.com/star-39/moe-sticker-bot/issues\n"+
		"Community-maintained version issues: https://github.com/Nununene/moe-sticker-bot/issues\n\n"+
		"<code>"+errMsg+"</code>", tele.ModeHTML, tele.NoPreview)
}

func sendExecEmojiAssignFinished(c tele.Context) error {
	ud := users.data[c.Sender().ID]
	msg := fmt.Sprintf(`
LINE Cat: <code>%s</code>
LINE ID: <code>%s</code>
TG ID: <code>%s</code>
TG Title: <a href="%s">%s</a>

Success. 成功完成. /start
	`, ud.lineData.Category,
		ud.lineData.Id,
		ud.stickerData.id,
		"https://t.me/addstickers/"+ud.stickerData.id,
		escapeTagMark(ud.stickerData.title),
	)
	return c.Send(msg, tele.ModeHTML)
}

// Return:
// string: Text of the message.
// *tele.Message: The pointer of the message.
// error: error
func sendProcessStarted(ud *UserData, c tele.Context, optMsg string) (string, *tele.Message, error) {
	message := fmt.Sprintf(`
Preparing stickers, please wait...
正在準備貼圖, 請稍後...

LINE Cat: <code>%s</code>
LINE ID: <code>%s</code>
TG ID: <code>%s</code>
TG TYPE: <code>%s</code>
TG Title: <a href="%s">%s</a>

<b>Progress / 進展</b>
<code>%s</code>
`, ud.lineData.Category,
		ud.lineData.Id,
		ud.stickerData.id,
		ud.stickerData.stickerSetType,
		"https://t.me/addstickers/"+ud.stickerData.id,
		escapeTagMark(ud.stickerData.title),
		optMsg)
	ud.progress = message

	teleMsg, err := c.Bot().Send(c.Recipient(), message, tele.ModeHTML)
	ud.progressMsg = teleMsg
	return message, teleMsg, err
}

// if progressText is empty, a progress bar will be generated based on cur and total.
func editProgressMsg(cur int, total int, progressText string, originalText string, teleMsg *tele.Message, c tele.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Errorln("editProgressMsg encountered panic! ignoring...", string(debug.Stack()))
		}
	}()

	header := originalText[:strings.LastIndex(originalText, "<code>")]
	prog := ""

	if progressText != "" {
		prog = progressText
		goto SEND
	}
	cur = cur + 1
	if cur == 1 {
		prog = fmt.Sprintf("<code>[=>                  ]\n       %d of %d</code>", cur, total)
	} else if cur == int(float64(0.25)*float64(total)) {
		prog = fmt.Sprintf("<code>[====>               ]\n       %d of %d</code>", cur, total)
	} else if cur == int(float64(0.5)*float64(total)) {
		prog = fmt.Sprintf("<code>[=========>          ]\n       %d of %d</code>", cur, total)
	} else if cur == int(float64(0.75)*float64(total)) {
		prog = fmt.Sprintf("<code>[==============>     ]\n       %d of %d</code>", cur, total)
	} else if cur == total {
		prog = fmt.Sprintf("<code>[====================]\n       %d of %d</code>", cur, total)
	} else {
		return nil
	}
SEND:
	messageText := header + prog
	c.Bot().Edit(teleMsg, messageText, tele.ModeHTML)
	return nil
}

func sendAskSToManage(c tele.Context) error {
	return c.Send("Send a sticker from the sticker set that want to edit,\n" +
		"or send its share link.\n\n" +
		"您想要修改哪個貼圖包? 請傳送那個貼圖包內任意一張貼圖,\n" +
		"或者是它的分享連結.")
}

func sendUserOwnedS(c tele.Context) error {
	usq := queryUserS(c.Sender().ID)
	if usq == nil {
		return errors.New("no sticker owned")
	}

	var entries []string

	for _, us := range usq {
		date := time.Unix(us.timestamp, 0).Format("2006-01-02 15:04")
		title := strings.TrimSuffix(us.tg_title, " @"+botName)
		//workaround for empty title.
		if title == "" || title == " " {
			title = "_"
		}
		entry := fmt.Sprintf(`<a href="https://t.me/addstickers/%s">%s</a>`, us.tg_id, title)
		entry += " | " + date
		entries = append(entries, entry)
	}

	if len(entries) > 30 {
		eChunks := chunkSlice(entries, 30)
		for _, eChunk := range eChunks {
			message := "You own following stickers:\n"
			message += strings.Join(eChunk, "\n")
			c.Send(message, tele.ModeHTML)
		}
	} else {
		message := "You own following stickers:\n"
		message += strings.Join(entries, "\n")
		c.Send(message, tele.ModeHTML)
	}
	return nil
}

func sendAskEditChoice(c tele.Context) error {
	ud := users.data[c.Sender().ID]
	selector := &tele.ReplyMarkup{}
	btnAdd := selector.Data("Add sticker/增添貼圖", CB_ADD_STICKER)
	btnDel := selector.Data("Delete sticker/刪除貼圖", CB_DELETE_STICKER)
	btnDelset := selector.Data("Delete sticker set/刪除貼圖包", CB_DELETE_STICKER_SET)
	btnChangeTitle := selector.Data("Change title/修改標題", CB_CHANGE_TITLE)
	btnExit := selector.Data("Exit/退出", "bye")

	if msbconf.WebappUrl != "" {
		baseUrl, _ := url.JoinPath(msbconf.WebappUrl, "edit")
		url := fmt.Sprintf("%s?ss=%s&dt=%d",
			baseUrl,
			ud.stickerData.id,
			time.Now().Unix())
		log.Debugln("WebApp URL is : ", url)
		webApp := &tele.WebApp{
			URL: url,
		}
		btnEdit := selector.WebApp("Change order or emoji/修改順序或Emoji", webApp)
		selector.Inline(
			selector.Row(btnAdd), selector.Row(btnDel), selector.Row(btnDelset), selector.Row(btnEdit), selector.Row(btnChangeTitle), selector.Row(btnExit))
	} else {
		selector.Inline(
			selector.Row(btnAdd), selector.Row(btnDel), selector.Row(btnDelset), selector.Row(btnChangeTitle), selector.Row(btnExit))
	}

	return c.Send(fmt.Sprintf(`
ID: <code>%s</code>
Title: <a href="https://t.me/addstickers/%s">%s</a>

What do you want to edit? Please select below:
您想要修改貼圖包的甚麼內容? 請選擇:`,
		users.data[c.Sender().ID].stickerData.id,
		ud.stickerData.id,
		ud.stickerData.title),
		selector, tele.ModeHTML)
}

func sendAskSDel(c tele.Context) error {
	return c.Send("Which sticker do you want to delete? Please send it.\n" +
		"您想要刪除哪一個貼圖? 請傳送那個貼圖")
}

func sendConfirmDelset(c tele.Context) error {
	selector := &tele.ReplyMarkup{}
	btnYes := selector.Data("Yes", CB_YES)
	btnNo := selector.Data("No", CB_NO)
	selector.Inline(selector.Row(btnYes), selector.Row(btnNo))

	return c.Send("You are attempting to delete the whole sticker set, please confirm.\n"+
		"您將要刪除整個貼圖包, 請確認.", selector)
}

func sendSFromSS(c tele.Context, ssid string, reply *tele.Message) error {
	ss, _ := c.Bot().StickerSet(ssid)
	if reply != nil {
		c.Bot().Reply(reply, &ss.Stickers[0])
	} else {
		c.Send(&ss.Stickers[0])
	}
	return nil
}

func sendFLWarning(c tele.Context) error {
	return c.Send(`
It might take longer to process this sticker set (2-8 minutes)... 
This warning indicates that you might triggered Telegram's flood limit, and bot is trying to re-submit.
Due to this mechanism, resulted sticker set might contains duplicate or missing sticker, please check manually after done.

此貼圖包可能需要更長時間處理(2-8分鐘)...
看到這一條警告表示Telegram可能限制了您創建貼圖包的頻度, 且bot正在自動嘗試重新製作, 因此得出的貼圖包可能會重複或缺失貼圖, 請在完成製作後再檢查一下.
`)
}

func sendTooManyFloodLimits(c tele.Context) error {
	return c.Send("Sorry, it seems that you have triggered Telegram's flood limit for too many times, it's recommended try again after a while.\n" +
		"抱歉, 您似乎觸發了Telegram的貼圖製作次數限制, 建議您過一段時間後再試一次.")
}

func sendNoCbWarn(c tele.Context) error {
	return c.Send("Please press a button! /quit\n請選擇按鈕!")
}

func sendBadIDWarn(c tele.Context) error {
	return c.Send(`
Bad ID. try again or press Auto Generate. /quit
Can contain alphanum and underscore only, must begin with alphabet, must not contain consecutive underscores.
只可以含有英文,數字,下劃線, 必須由英文字母開頭，不可以有連續下劃線.
ID錯誤, 請試多一次或按下'自動生成'按鈕. /quit`)
}

func sendIDOccupiedWarn(c tele.Context) error {
	return c.Send("ID already occupied! try another one. ID已經被占用, 請試試另一個.")
}

func sendBadImportLinkWarn(c tele.Context) error {
	return c.Send("Invalid import link, make sure its a LINE Store link or kakao store link. Try again or /quit\n"+
		"無效的連結, 請檢視是否為LINE貼圖商店的連結, 或是kakao emoticon的連結.\n\n"+
		"For example: 例如:\n"+
		"<code>https://store.line.me/stickershop/product/7673/ja</code>\n"+
		"<code>https://e.kakao.com/t/pretty-all-friends</code>", tele.ModeHTML)
}

func sendNoSToManage(c tele.Context) error {
	return c.Send("Sorry, you have not created any sticker set yet. You can use /import or /create .\n" +
		"抱歉, 您還未創建過貼圖包, 您可以使用 /create 或 /import .")
}

func sendPromptStopAdding(c tele.Context) error {
	selector := &tele.ReplyMarkup{}
	btnDone := selector.Data("Done adding/停止添加", CB_DONE_ADDING)
	selector.Inline(selector.Row(btnDone))
	return c.Send("Continue sending files or press button below to stop adding.\n"+
		"請繼續傳送檔案. 或者按下方按鈕來停止增添.", selector)
}

func replySFileOK(c tele.Context, count int) error {
	selector := &tele.ReplyMarkup{}
	btnDone := selector.Data("Done adding/停止添加", CB_DONE_ADDING)
	selector.Inline(selector.Row(btnDone))
	return c.Reply(
		fmt.Sprintf("File OK. Got %d stickers. Continue sending files or press button below to stop adding.\n"+
			"檔案OK. 已收到%d份貼圖. 請繼續傳送檔案. 或者按下方按鈕來停止增添.", count, count), selector)
}

func sendSEditOK(c tele.Context) error {
	return c.Send(
		"Successfully edited sticker set. /start\n" +
			"成功修改貼圖包. /start")
}

func sendStickerSetFullWarning(c tele.Context) error {
	return c.Send(
		"Warning: Your sticker set is already full. You cannot add new sticker.\n" +
			"提示：當前貼圖包已滿，您將不能增添貼圖。")
}

// func sendEditingEmoji(c tele.Context) error {
// 	return c.Send("Commiting changes...\n正在套用變更，請稍候...")
// }

func sendAskSearchKeyword(c tele.Context) error {
	return c.Send("Please send a word that you want to search\n請傳送想要搜尋的內容")
}

func sendSearchNoResult(c tele.Context) error {
	message := "Sorry, no result.\n抱歉, 搜尋沒有結果."
	if c.Chat().Type == tele.ChatPrivate {
		message += "\nTry again or /quit\n請試試別的關鍵字或 /quit"
	}
	return c.Send(message)
}

func sendNotifyNoSessionSearch(c tele.Context) error {
	return c.Send("Here are some search results, use /search to dig deeper or /start to see available commands.\n" +
		"這些是貼圖包搜尋結果，使用 /search 詳細搜尋或 /start 來看看可用的指令。")
}

func sendUnsupportedCommandForGroup(c tele.Context) error {
	return c.Send("This command is not supported in group chat, please chat with bot directly.\n" +
		"此指令無法於群組內使用, 請與bot直接私訊.")
}

func sendBadSearchKeyword(c tele.Context) error {
	return c.Send(fmt.Sprintf(`
Please specify keyword
請指定搜尋關鍵字.

Example: 例如:
/search@%s keyword1 keyword2 ...
/search@%s nekomimi mia
`, botName, botName))
}

func sendPreferKakaoShareLinkWarning(c tele.Context) error {
	msg := `
The link you sent is a Kakao store link.
Use a share link for improved image quality and animated sticker support,
you can obtain it from KakaoTalk app by tapping share->copy link in sticker store.

您傳送的是Kakao商店的連結.
使用分享連結才能支援動態貼圖, 靜態貼圖的畫質也更高。
您可以在KakaoTalk App內的貼圖商店點選 分享->複製連結 來取得分享連結。

eg:例如: <code>https://emoticon.kakao.com/items/lV6K2fWmU7CpXlHcP9-ysQJx9rg=?referer=share_link</code>
`
	err := c.Reply(&tele.Photo{
		File:    tele.File{FileID: FID_KAKAO_SHARE_LINK},
		Caption: msg,
	}, tele.ModeHTML)
	if err != nil {
		c.Reply(msg, tele.ModeHTML)
	}
	return nil
}

func sendUseCommandToImport(c tele.Context) error {
	return c.Send("Please use /create to create sticker set using your own photos and videos. /start\n" +
		"請使用 /create 指令來使用自己的圖片和影片和創建貼圖包. /start")
}

func sendOneStickerFailedToAdd(c tele.Context, pos int, err error) error {
	return c.Reply(fmt.Sprintf(`
Failed to add one sticker.
一張貼圖添加失敗.
Index: %d
Error: %s
`, pos, err.Error()))
}

func sendBadSNWarn(c tele.Context) error {
	return c.Reply("Wrong sticker or link!\n貼圖或連結錯誤!")
}

func sendSSTitleChanged(c tele.Context) error {
	msg := `
Successfully changed title.
新標題設定完成`
	return c.Reply(msg, tele.ModeHTML)
}
func sendSSTitleFailedToChanged(c tele.Context) error {
	msg := `
Failed to change title, please try again.
新標題設定失敗，請再試一次。`
	return c.Reply(msg, tele.ModeHTML)
}

// func sendInvalidEmojiWarn(c tele.Context) error {
// 	return c.Reply(`
// Sorry, this emoji is invalid, it has been defaulted to ⭐️, you can edit it after done by using /manage command.
// 抱歉，這個emoji無效，並且已默認設定為⭐️，你可以在完成製作後使用 /manage 來修改。
// 	`)
// }

func sendProcessingStickers(c tele.Context) error {
	return c.Send(`
Processing stickers, please wait a while...
正在製作貼圖，請稍等...
`)
}
