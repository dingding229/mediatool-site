# Media Tool Go 站点模板

站点模板由Go类型与构造函数编写：HDSKY、OURBITS、CHDBITS、SSD、TTG、UBITS。主程序通过固定Go模块版本编译引入，启动时自动登记全部模板；运行时不拉取/编译远程源码，不提供JSON导入或站点选择导入。

每个站点定义内置域名、实际认证字段、请求解析规则与RSS获取方式。当前六站需要Cookie和User-Agent；用户ID由登录页面获取，RSS从真实页面链接或站点getrss生成接口读取。无法识别的RSS不会猜测生成。

内部Snapshot只用于数据库持久化及测试，不是对外JSON模板。改站点代码后先运行go test ./...，再发布提交，由主程序更新固定模块版本并经Actions构建部署。

来源选择器改写自mantou568/pre-dessert-sites提交37bf5c5064ae437f54e26b57406ca84d480daab5，保留原署名和条款于LICENSE.pre-dessert-sites。TTG参数参考MoviePilot对应测试，未复制其运行代码。模板格式不兼容原配置库；保留许可不代表旧格式兼容。

## Template metadata

Each compiled template carries an independent semantic `Version` (currently 1.1.0), `Domain`, and `IconURL`. The host displays the full catalog but loads only administrator-selected templates. Shared request builders in `common.go` are grouped by login, user details, search, and seeding statistics; per-site files override only site-specific rules. Parser snapshot format version is separate from template release version.
