# Media Tool 精简站点模板

更新：2026-10-13。六个独立JSON模板仅保留主程序需要的功能，无需公共配置继承；旧格式common参数不再使用。模板不是插件，不包含Cookie、Passkey、用户ID或API密钥。索引为`index.json`。

## 站点与实际差异

| 站点 | 文件 | 程序站点ID | 搜索 | 做种统计 |
|---|---|---|---|---|
| HDSKY | hdsky.json | hdsky | torrents.php；下载form action | AJAX列表，第3列大小 |
| CHDBITS（彩虹岛） | ptchdbits.json | ptchdbits | b标题、circle-text H&R、subtitle/tag | AJAX列表，第3列大小 |
| SSD（春天） | ssd.json | ssd | span促销、描述/标签；魔力取title | AJAX汇总：两个b提供数量和总体积 |
| OURBITS（我堡） | ourbits.json | ourbits | rowfollow资源大小、hitandrun、tags标签 | AJAX列表，第3列大小 |
| UBITS | ubits.json | ubits | 描述/标签、hitandrun | AJAX汇总：“条记录/总大小”容器 |
| TTG | ttg.json | ttg | browse.php?c=M，关键字为search_field，torrent_table/name_left/dl_a | userdetails当前上传列表，第4列大小 |

CHDBITS沿用上游真实ID `ptchdbits`，文件名可读但不编造上游ID。TTG保留独立`ttg`框架标识；静态兼容报告仍提示需专用站点实测，不能用改名NexusPHP掩盖限制。TTG未自动把标题转换成IMDb搜索，也未添加分类词；原始关键字交给站点。

## 每个模板仅有4个请求

1. `user_basic_info`：登录验证及id/name/uploaded/downloaded/ratio/bonus。
2. `user_details`：等级。
3. `search`：搜索/订阅资源发现，保留id/title/details/download/size/seeders/leechers/downloadvolumefactor/description/labels，以及有证据的H&R字段。
4. `seeding_statistics`：做种数量与总体积；列表型累加经验证的完整分页，汇总型直接解析明确的总计。

RSS URL放在PT账号中，使用通用RSS模块，不复制生成RSS、签到/验证码、站内信、公告、字幕、聊天、图片抓取等接口。分辨率、来源、编码、标题/年份、季/集由主程序0day解析/筛选处理，不复制上游搜索表单、封面、完成数、日期、IMDb/豆瓣等未使用字段。

## 安全与数据语义

- 只使用搜索页给出的真实下载href/action，不构造Passkey或猜测取种地址。TTG保留自己的下载路径，不套用Nexus download.php。
- 未证实的促销保持未知，不用上游30%规则的默认`1`把没有徽标的资源标成确定付费。
- CHDBITS/TTG的H&R天数保留完整数字（14不能截成1）；OURBITS/UBITS仅在真实hitandrun徽标存在时证明有HR，不捏造固定2/3天。没有徽标不证明无HR；SSD没有可靠逐资源HR规则，保持未知。
- 搜索必须存在唯一的声明结果容器；登录/挑战/错误页不是成功空结果。经验证的空表才可表示该请求为空。
- 仅跟随响应中的真实next链接，必须同源、同路径、保留搜索/用户身份；只允许数值分页变化。最多5页/500资源，异常、循环和上限保留partial，不当作完整结果。
- 列表做种响应缺失/空白/重复/体积不合法，保持未知；只有登录与用户/路径匹配的Nexus AJAX明确`没有记录`才证明0。TTG缺少列表不证明0。
- SSD/UBITS显式`summary`模式要求唯一容器、合法非负整数和完整总体积；页码参数、身份/路径变化、认证表单、负数、溢出、缺失字段、多个汇总或数量/大小矛盾均拒绝。SSD仅接受两处b的完整格式，UBITS只接受完整“数量条记录/总大小”文本；未知页面格式应适配而不是猜测总数。
- 数据未知时仍遵循既有订阅规则；不会为了通过订阅而自动允许未知HR。

## 原生仓库与格式

Media Tool内置的唯一仓库为dingding229/mediatool-site。模板位于sites/<id>.json（CHDBITS为ptchdbits.json），不再使用原有公共配置继承。根字段templateVersion必须为数字1；每份模板完整声明requests。旧模板、reuse_schema、disabled_extends和非空common均不支持。

index.json的catalogVersion为1，repository固定本仓库；每项具有id/name/file/templateVersion/SHA256。摘要针对原始UTF-8文件字节（含尾换行），修改模板必须同时更新index。主程序固定Git提交后读取索引与模板，拒绝摘要、ID、版本或安全路径不符，整批验证成功后原子保存。索引不能指定外部URL或../路径。

模板中不保存Cookie/Passkey/用户ID/API密钥。后台在PT站点页面选择仓库目录、确认许可后导入；旧来源记录不能按新仓库的相同SHA冒充恢复。导入改变解析指纹，旧搜索/统计失效，仍须真实账号验证。不会自动安装或修改媒体。

## 验收与来源

运行`go test ./internal/modules/sites -run 'TestMinimalTemplates|TestSummary|TestTemplateIndex|TestHDSKYMinimal'`。样本是自有合成HTML和受控transport；覆盖真实模板加载、登录拒绝、电影/剧集标题、促销/H&R、取种、分页、空结果/挑战、账号指标、列表和汇总统计、异常数字与请求边界。不访问PT站点、不下载真实种子。

这五个新模板本轮**没有真实账号实测**；不将合成页面验证称为站点验收完成。HDSKY原站点已有前轮只读实测，但精简模板本身也不冒充新实测。

选择器精简自mantou568/pre-dessert-sites提交`37bf5c5064ae437f54e26b57406ca84d480daab5`的site_config/sites以及NexusPHP公共配置，原署名/条款见LICENSE.pre-dessert-sites。TTG search_field/c=M交叉核对MoviePilot提交`6ef8408f432dde643aee2ba7078f59d2776f84ce`的test_indexer_spider_search_url.py；未复制其运行代码。模板由本仓库发布，主程序只内置仓库和解析能力；保留来源许可是授权义务，不是旧格式兼容。
