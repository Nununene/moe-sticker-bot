# [@nununene_sticker_bot](https://t.me/nununene_sticker_bot)

[![Go Reference](https://pkg.go.dev/badge/github.com/star-39/moe-sticker-bot.svg)](https://pkg.go.dev/github.com/star-39/moe-sticker-bot) ![Go Report](https://goreportcard.com/badge/github.com/star-39/moe-sticker-bot) ![CI](https://github.com/Nununene/moe-sticker-bot/actions/workflows/msb_oci.yml/badge.svg) ![CI](https://github.com/Nununene/moe-sticker-bot/actions/workflows/build_binaries.yml/badge.svg)

[![MSB](https://user-images.githubusercontent.com/75669297/222379608-1359ac0f-18ed-4a25-a91e-32974994d27b.png)](https://t.me/nununene_sticker_bot)

---

Use moe-sticker-bot, a Telegram bot, to easily import or download LINE/Kakaotalk/Telegram stickers, use your own image or video to create Telegram sticker set or CustomEmoji and manage it.

This is a community-maintained fork of [star-39/moe-sticker-bot](https://github.com/star-39/moe-sticker-bot), maintained at [Nununene/moe-sticker-bot](https://github.com/Nununene/moe-sticker-bot). For community updates, visit [@nununenenununeno](https://t.me/nununenenununeno). It contains maintenance fixes and deployment-related improvements.

---

Telegram用萌萌貼圖BOT。

匯入或下載LINE和kakaotalk貼圖包到Telegram. 使用自己的圖片和影片創建Telegram貼圖包或表情貼並管理.

下載Telegram貼圖包/GIF。

本專案是 [star-39/moe-sticker-bot](https://github.com/star-39/moe-sticker-bot) 的社群維護版本，維護於 [Nununene/moe-sticker-bot](https://github.com/Nununene/moe-sticker-bot)，持續整理及修正程式。

## Features/功能

- Import LINE or kakao stickers to Telegram without effort, you can batch or separately assign emojis.
- Create your own sticker set or CustomEmoji with your own images or videos in any format.
- Support mixed-format sticker set. You can put animated and static stickers in the same set.
- Batch download and convert Telegram stickers or GIFs to original or common formats.
- Export Telegram stickers to WhatsApp (requires [Msb App](https://github.com/star-39/msb_app), supports iPhone and Android).
- Manage your sticker set interactively through WebApp: add/move/remove/edit sticker and emoji.
- Provides a CLI app [msbimport](https://github.com/Nununene/moe-sticker-bot/tree/master/pkg/msbimport) for downloading LINE/Kakaotalk stickers.

- 輕鬆匯入LINE/kakao貼圖包到Telegram, 可以統一或分開指定emoji.
- 輕鬆使用自己任意格式的圖片和影片來創建自己的貼圖包或表情貼.
- 支援混合貼圖包。動態靜態貼圖可以放在同一個包內。
- 下載Telegram/LINE/kakao貼圖包和GIF, 自動變換為常用格式, 並且保留原檔.
- 匯出Telegram的貼圖包至WhatsApp（需要安裝[Msb App](https://github.com/star-39/msb_app), 支援iPhone和Android）。
- 互動式WebApp可以輕鬆管理自己的貼圖包: 可以新增/刪除貼圖, 移動位置或修改emoji.
- 提供名為[msbimport](https://github.com/Nununene/moe-sticker-bot/tree/master/pkg/msbimport)的終端機程式， 用於下載LINE/kakao貼圖。

## Screenshots

[![MSB](https://img.shields.io/badge/-%40nununene__sticker__bot-blue?style=plastic&logo=telegram)](https://t.me/nununene_sticker_bot)

![Screenshot](https://user-images.githubusercontent.com/75669297/221539624-c0cc32a9-477c-425f-8e98-6566326385b4.png)

![Screenshot](https://user-images.githubusercontent.com/75669297/221538927-526a878a-5d86-4b45-ab9a-d324743e3b91.png)

![Screenshot](https://user-images.githubusercontent.com/75669297/221541547-4618c9ef-9be3-4d50-b7da-fdc5b25e64b8.png)

![Screenshot](https://user-images.githubusercontent.com/75669297/206968889-1fe25c05-6071-422b-9e1b-549d56f5d351.png)

![Screenshot](https://user-images.githubusercontent.com/75669297/218156358-0145264f-ab11-4010-bfcd-2e38621d7381.png)

![Screenshot](https://user-images.githubusercontent.com/75669297/218153727-5fb1d3e0-3770-4dc8-a2b5-3e0ecd89a003.png)

![Screenshot](https://user-images.githubusercontent.com/75669297/221529085-2581bcca-fe49-46b0-8123-5614e90a838c.png)

## Deployment

It is recommended to deploy moe-sticker-bot using containers.

### Build a container image from this fork

This fork provides a custom Dockerfile:

[`deployments/docker/bot.Dockerfile`](deployments/docker/bot.Dockerfile)

The Dockerfile compiles the bot executable from this repository and uses the upstream runtime image as its base image.

Build the image from the repository root:

```bash
git clone https://github.com/Nununene/moe-sticker-bot.git
cd moe-sticker-bot

docker build \
  -f deployments/docker/bot.Dockerfile \
  -t moe-sticker-bot:local .
```

The resulting image is:

```text
moe-sticker-bot:local
```

**Runtime image note:**

The current Dockerfile uses `ghcr.io/star-39/moe-sticker-bot:latest` as its runtime base image. The bot executable itself is compiled from this fork's source code, while the runtime environment and its system dependencies are inherited from the upstream image.

Therefore, building this image still requires access to the upstream GHCR image.

### Upstream pre-built container

The original upstream project provides a pre-built OCI container:

[star-39/moe-sticker-bot on GitHub Packages](https://github.com/users/star-39/packages/container/package/moe-sticker-bot)

This image is maintained separately from this fork and does not automatically contain the changes made in this repository.

If you want to use the changes in this fork, build the custom image described above.

### Kubernetes deployment

See the original deployment example:

[`deployments/kubernetes_msb.yaml`](deployments/kubernetes_msb.yaml)

This is an upstream deployment example. Review the image name, credentials and environment-specific settings before using it.

### System Dependencies

For deployments outside the provided container image, the following dependencies must be installed and accessible through `PATH`:

- ImageMagick (6 or 7, both fine)
- bsdtar (libarchive-tools)
- ffmpeg (Requires a lot of components, use a fat static build or package manager's ones pls!)
- curl
- gifsicle (for converting GIF)
- python3 (for following tools)
- [msb_emoji.py](tools/msb_emoji.py) (for extracting emoji)
- [msb_kakao_decrypt.py](tools/msb_kakao_decrypt.py) (for decrypting animated kakao)
- [msb_rlottie.py](tools/msb_rlottie.py) (for converting TGS)
- mariadb-server (optional, for database)
- nginx (optional, for WebApp)

Dependencies above must be accessible through PATH.

## Build

### Build Dependencies

- Go (the project uses Go modules; Go 1.23 is used by the provided Dockerfile)
- nodejs v18+ (optional, for WebApp)
- react-js v18+ (optional, for WebApp)

Clone this fork:

```bash
git clone https://github.com/Nununene/moe-sticker-bot.git
cd moe-sticker-bot
```

Build the bot executable:

```bash
go build -o moe-sticker-bot cmd/moe-sticker-bot/main.go
```

### WebApp

Since version 2.0 of moe-sticker-bot, managing sticker set order and emoji is achieved using Telegram's WebApp technology.

See details in [web/webapp3](web/webapp3).

## CHANGELOG

### 2026-11-07

- **Fix Telegram WebApp initData validation.**
  - Validate all submitted initData fields except `hash`.
  - Sort fields by key and construct the data-check string according to Telegram's WebApp validation format.
  - Fix the HMAC validation failure that prevented the WebApp from opening when managing sticker sets.

- **Add a custom Dockerfile for building this fork.**
  - Add `deployments/docker/bot.Dockerfile`.
  - Compile the bot executable from this repository's source code.
  - Use the upstream runtime image as the base image.

---
<details>
<summary>Old changelogs</summary>
## Old Changelogs

The following changelogs are retained from the original upstream project.

### v2.5.0-RC1 (20240531)

- Support mix-typed sticker set.
- You can add video to static set and vice versa.
- Removed WhatsApp export temporarily.
- Many bug fixes.

### v2.4.0-RC3 (20240226)

- Remove deduplication during database curation.
- LINE emoji can be imported to either sticker or CustomEmoji.
- Support creating sticker set with under 50 stickers in batch.
- Import speed is increased significantly.

### v2.4.0-RC1-RC2 (20240207)

- Support importing LINE Emoji into CustomEmoji.
- Support creating CustomEmoji.
- Support editing sticker emoji and title.

### v2.3.14-2.3.15 (20230228)

- Fix missing libwebp in OCI.
- Support TGS export.
- Improve GIF efficiency.

### v2.3.11-v2.3.13 (20230227)

- Fix flood limit by ignoring it.
- Fix managing video stickers.
- Improved WhatsApp export.
- Support region locked LINE Message sticker.
- 修復flood limit匯入錯誤。
- 修復動態貼圖管理。
- 改善WhatsApp匯出。
- 支援有區域鎖的LINE訊息貼圖。

### v2.3.8-2.3.10 (20230217)

- Fix kakao import fatal, support more animated kakao.
- 修復KAKAO匯入錯誤，支援更多KAKAO動態貼圖。
- Fix static webp might oversize (256KB).
- Fix panic during assign: invalid sticker emojis.
- Fix some kakao animated treated as static.
- Improved kakao static sticker quality.
- Improved user experience.

### v2.3.6-2.3.7 (20230213)

- Support webhook.
- Support animated webp for user sticker.
- Add change sticker set title feature.
- Fix "sticker order mismatch" when using WebApp to sort.
- Fix error on emoji assign.
- Fix too large animated kakao sticker.

### v2.3.1-2.3.5 (20230209)

- Fix i18n titles.
- Fix flood limit by implementing channel to limit autocommit concurrency.
- Fix error on WebApp.
- Fix import hang.
- Fix fatal error not being reported to user.
- Fix typos.

### v2.3.0 (20230207)

- Fix flood limit by using local API server.
- Support webhook and local API server.
- Huge performance gain.
- Fix `/search` panic.

### v2.2.1 (20230204)

- Fix webm too big.
- Fix id too long.
- Tuned flood limit algorithm.

### v2.2.0 (20230131)

- Support animated kakao sticker.
- 支援動態kakao貼圖。

### v2.1.0 (20230129)

- Support exporting sticker to WhatsApp.
- 支援匯出貼圖到WhatsApp。

### 2.0.1 (20230106)

- Fix special LINE officialaccount sticker.
- Fix `--log_level` command-line parsing.
- Thank you all! This project has reached 100 stars!

### 2.0.0 (20230105)

- Use new WebApp from `/manage` command to edit sticker set with ease.
- Send text or use `/search` command to search imported LINE/kakao sticker sets by all users.
- Auto import now happens in background.
- Downloading sticker set is now a lot faster.
- Fix many LINE import issues.
- 通過/manage指令使用新的WebApp輕鬆管理貼圖包。
- 直接傳送文字或使用/search指令來搜尋所有用戶匯入的LINE/KAKAO貼圖包。
- 自動匯入現在會在背景處理。
- 下載整個貼圖包的速度現在會快許多。
- 修復了許多LINE貼圖匯入的問題。

### Detailed 2.0 Changelogs

#### 2.0.0 (20230104)

- Improve flood limit handling.
- Auto LINE import now happens in background.
- Improve GIF download.

#### 2.0.0 RC-7 (20221230)

- Support `/search` in group chat.
- Fix search result.
- Fix empty sticker title.
- Sticker download is now parallel.

#### 2.0.0 RC-6 (20221220)

- Fix line APNG with unexpected `tEXt` chunk.
- Changed length of webm from 2.9s to 3s.
- Minor improvements.

#### 2.0.0 RC-5 (20221211)

- Fix potential panic when onError.
- Warn user sticker set is full.
- Fix LINE message sticker with region lock.

#### 2.0.0 RC-4 (20221211)

- Fix edit sticker on iOS.
- Fix error editing multiple emojis.

#### 2.0.0 RC-3 (20221210)

- Complies to LINE store's new UA requirements.
- Fix animated sticker in WebApp.
- Fixed sticker download.
- Fixed WebApp image aspect ratio.

#### 2.0.0 RC-2 (20221210)

- Fix fatal panic on WebApp.
- Add `/search` functionality.
- Removed gin TLS support.
- Auto database curation.

#### 2.0.0 RC-1 (20221206)

- WebApp support for editing stickers.
- Code structure refactored.
- Now accepts options from command line instead of environment variables.
- Support parallel sticker download.
- Fix LINE officialaccount/event/sticker.
- Fix kakao link with queries.

### 1.2.4 (20221111)

- Minor improvements.
- Fixed (almost) flood limit.
- Fixed kakao link with queries.

### 1.2.2 (20220523)

- Improved user experience.

### 1.2.1 (20220520)

- Improved emoji edit.

### 1.2 (20220518)

- Fix import error for LINE ID < 775.
- Improved UX during `/import`.
- Warn user if sticker set is already full.

### 1.1 (20220517)

- Code refactors.
- UX improvements.
- Skip error on TGS to GIF due to lottie issues.

### 1.0 (20220513)

- First stable release in Go version.
- Added support for downloading TGS and converting to GIF.
- Backing database for `@moe_sticker_bot` has gone complete sanitization.

### 1.0 RC-9 (20220512)

- Add an administrative command to sanitize database, which purges duplicated stickers.
- Add an advanced command `/register`, to register your sticker to database.
- Minor bug fixes.
- This is the REAL final RC release, next one is stable!

### 1.0 RC-8 GO (20220512)

- Fix random number in ID.
- Major code refactor.
- Download sticker now happens in background.
- Better documentation.
- This release should be the final RC... hopefully.

### 1.0 RC-7 GO (20220511)

- You can specify custom ID when `/create`.
- Changed import ID naming scheme for cleaner look.
- Die immediately if FloodLimit exceeds 4.
- If everything looks good, this should be the last RC for 1.0.

### 1.0 RC-6 GO (20220511)

- New feature: Change sticker order.
- New feature: Edit sticker emoji.
- New import support: kakaotalk emoticon.
- Fix possible panic when editMessage.
- We are closing to a stable release! RC should not exceed 8.

### 1.0 RC-5 GO (20220510)

- New feature: download raw LINE stickers to zip.
- FatalError now prints stack trace.
- zh-Hant is now default in auto LINE title.
- Quality of video sticker should improve by a bit.
- Fix possible slice out of range panic.
- If user experiences FloodLimit over 3 times, terminate process.

### 1.0 RC-4 GO (20220509)

- Use my custom fork of telebot.
- User sent sticker now supports any file.

### 1.0 RC-3 GO (20220509)

- Split large zip to under 50MB chunks.
- Split long message to chunks.
- GIF downloaded is now in original 512px resolution.
- You can press "No" now when sent link or sticker.
- If error is not HTTP400, bot will retry for 3 times.
- Other minor improvements.

### 1.0 RC-2 GO (20220508)

- Fix SEGV when user requested `/quit`.
- Ignore FloodLimit by default since API will do retry at TDLib level.
- Fix emoji in database.
- Fix video sticker when `/manage`.
- Support LINE message sticker and emoticon (emoji).

### 1.0 RC-1 GO (20220506)

- Completely rewritten whole project to Go.
- Countless bug fixes.
- You can send sticker or link without a command now.
- Performance gained by a lot thanks to goroutine and worker pool.

### Older Python versions

#### 5.1 RC-4 (20220423)

- Fix duplicated sticker.
- Fix alpha channel converting GIF.

#### 5.1 RC-3 (20220416)

- Do not use joblib due to their bugs.
- `/download_telegram_sticker` now converts to standard GIFs.

#### 5.1 RC-2 (20220326)

- Significantly improved performance by using parallel loop.
- Sanitize kakao id starting with `-`.

#### 5.1 RC-1 (20220309)

- Support kakaotalk emoticon.
- Add more checks for telegram sticker id.

#### 5.0 RC-12 (20220305)

- Database now records LINE link and emoji settings.
- Fix issue when LINE name has `<` and `>` marks.
- Fix issue adding video to static set.
- Fix hang handling CallbackQuery.

#### 5.0 RC-11 (20220303)

- You can now delete a sticker set.
- `/manage_sticker_set` will show sets you created.
- Fix missing sticker during USER_STICKER.

#### 5.0 RC-10 (20220226)

- Performance is now significantly improved.
- Fix issue converting LINE message stickers.
- Bypass some regional block by LINE.

#### 5.0 RC-9 (20220223)

- Split LINE popups to two categories, one keeping animated only.
- Bot now has a database storing "good" imported sticker sets.
- Fix duplicated stickers in sticker set.

#### 5.0 RC-8 (20220222)

- Fix user sticker parsing.
- Add support for MdIcoFlashAni_b.

#### 5.0 RC-7 (20220215)

- Fix exception if user sent nothing during USER_STICKER.
- Fix a bug where `/import_line_sticker` may have no response.
- Corrected file download limit.
- Fix animated sticon.
- Fix import hang due to missing ffmpeg `-y` parameter.

#### 5.0 RC-6 (20220215)

- Fix python-telegram-bot WebHook problem.
- Fix emoji assign.
- Fix black background video sticker.
- Fix "Sticker too big" when uploading video sticker.

#### 5.0 RC-5 (20220214)

- Allow using WebHook for better performance.
- Code refactors.
- Support LINE name sticker.

#### 5.0 RC-4 (20220212)

- Improved user sticker experience.

#### 5.0 RC-3 (20220212)

- Fix a bug where creating sticker set with one sticker will cause fatal error.
- Fix missing clean_userdata in `/download_line_sticker`.
- Tune VP9 params to avoid hitting 256K limit. This reduces video quality by a bit.

#### 5.0 RC-2 (20220211)

- Fix media_group.
- Minor bug fixes.
- Version 5.0 now enters feature freeze. No new feature will be added. Will have bug fixes only.

#### 5.0 RC-1 (20220211)

- Support LINE popup sticker without sound.
- Support AVIF.
- Many bug fixes.

#### 5.0 ALPHA-1 (20220211)

- Full support of animated (video) stickers.
- New feature: `/manage_sticker_set`, now you can add, delete, move stickers in a sticker set.
- Add support for LINE full-screen sticker (animated).

#### 4.0 ALPHA-5 (20220210)

- Bring back fake RetryAfter check since some people still having this issue.

#### 4.0 ALPHA-4 (20220210)

- Support user-uploaded animated (video) stickers. You can both create or add to a set.
- Better support sticon (LINE emoji).
- Bug fixes.

#### 4.0 ALPHA-3 (20220209)

- Supports all special LINE stickers, including effect_animation and sticon (emoji).

#### 4.0 ALPHA-1 (20220209)

- Supports animated LINE sticker import.

---
</details>

## Special Thanks

For the free 4CPU AArch64 Cloud Instance.

[貼圖群 - Sticker Group Taiwan](http://t.me/StickerGroup) for testing and reporting.

[LINE Corp](https://linecorp.com/) / [Kakao Corp](http://www.kakaocorp.com/) for cute stickers.

[https://github.com/blluv/KakaoTalkEmoticonDownloader](https://github.com/blluv/KakaoTalkEmoticonDownloader) MIT License Copyright @blluv

[https://github.com/laggykiller/rlottie-python](https://github.com/laggykiller/rlottie-python) GPL-2.0 license Copyright @laggykiller

You and all the users! ☺

## License

The GPL V3 License

![GPLv3](https://www.gnu.org/graphics/gplv3-with-text-136x68.png)
