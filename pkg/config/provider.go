package config

// 定义平台类型（Provider）
type MediaXVendor string
type VendorApp string

const (
	// 微信生态（WeChat）平台
	WechatMediaXVendor MediaXVendor = "WeChat"

	// 微信相关插件
	WechatOfficialAccount VendorApp = "WechatOfficialAccount" // 公众号
	WechatMiniProgram     VendorApp = "WechatMiniProgram"     // 小程序
	WechatChannel         VendorApp = "WechatChannel"         // 视频号
	WechatWork            VendorApp = "WechatWork"            // 企业微信

	// 抖音生态（DouYin）平台
	DouYinMediaXVendor MediaXVendor = "DouYin"

	// 抖音插件
	DouYin VendorApp = "DouYin" // 抖音

	// 小红书生态（RedBook）平台
	RedBookMediaXVendor MediaXVendor = "RedBook"

	// 小红书插件
	RedBook VendorApp = "RedBook" // 小红书

	// YouTube平台
	YouTubeMediaXVendor MediaXVendor = "YouTube"

	// YouTube插件
	YouTube VendorApp = "YouTube" // YouTube
)
