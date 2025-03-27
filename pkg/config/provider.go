package config

// 定义平台类型（Provider）
type MediaXVendor string
type VendorApp string

const (
	// 微信生态（WeChat）平台
	WechatMediaXVendor MediaXVendor = "WeChat"

	// 微信相关应用
	WechatOfficialAccount VendorApp = "WechatOfficialAccount" // 公众号
	WechatMiniProgram     VendorApp = "WechatMiniProgram"     // 小程序
	WechatChannel         VendorApp = "WechatChannel"         // 视频号
	WechatWork            VendorApp = "WechatWork"            // 企业微信

	// 抖音生态（DouYin）平台
	DouYinMediaXVendor MediaXVendor = "DouYin"

	// 抖音相关应用
	DouYin      VendorApp = "DouYin"      // 抖音
	DouYinStore VendorApp = "DouYinStore" // 抖音小店
	DouYinLive  VendorApp = "DouYinLive"  // 抖音直播
	DouYinAd    VendorApp = "DouYinAd"    // 抖音广告

	// 小红书生态（RedBook）平台
	RedBookMediaXVendor MediaXVendor = "RedBook"

	// 小红书相关应用
	RedBook      VendorApp = "RedBook"      // 小红书
	RedBookStore VendorApp = "RedBookStore" // 小红书商城
	RedBookAd    VendorApp = "RedBookAd"    // 小红书广告

	// 快手生态（Kuaishou）平台
	KuaishouMediaXVendor MediaXVendor = "Kuaishou"

	// 快手相关应用
	Kuaishou      VendorApp = "Kuaishou"      // 快手
	KuaishouLive  VendorApp = "KuaishouLive"  // 快手直播
	KuaishouStore VendorApp = "KuaishouStore" // 快手小店
	KuaishouAd    VendorApp = "KuaishouAd"    // 快手广告

	// Bilibili 生态（Bilibili）平台
	BilibiliMediaXVendor MediaXVendor = "Bilibili"

	// Bilibili 相关应用
	Bilibili     VendorApp = "Bilibili"     // Bilibili
	BilibiliLive VendorApp = "BilibiliLive" // Bilibili 直播
	BilibiliAd   VendorApp = "BilibiliAd"   // Bilibili 广告

	// 知乎生态（Zhihu）平台
	ZhihuMediaXVendor MediaXVendor = "Zhihu"

	// 知乎相关应用
	Zhihu     VendorApp = "Zhihu"     // 知乎
	ZhihuAd   VendorApp = "ZhihuAd"   // 知乎广告
	ZhihuLive VendorApp = "ZhihuLive" // 知乎 Live

	// 微博生态（Weibo）平台
	WeiboMediaXVendor MediaXVendor = "Weibo"

	// 微博相关应用
	Weibo   VendorApp = "Weibo"   // 微博
	WeiboAd VendorApp = "WeiboAd" // 微博广告

	// YouTube 平台
	GoogleMediaXVendor MediaXVendor = "Google"

	// YouTube 相关应用
	YouTube       VendorApp = "YouTube"       // YouTube
	YouTubeAd     VendorApp = "YouTubeAd"     // YouTube 广告
	YouTubeShorts VendorApp = "YouTubeShorts" // YouTube Shorts
	Blogger       VendorApp = "Blogger"       // YouTube Shorts

	// TikTok 平台
	TikTokMediaXVendor MediaXVendor = "TikTok"

	// TikTok 相关应用
	TikTok     VendorApp = "TikTok"     // TikTok
	TikTokAd   VendorApp = "TikTokAd"   // TikTok 广告
	TikTokShop VendorApp = "TikTokShop" // TikTok 小店
	TikTokLive VendorApp = "TikTokLive" // TikTok 直播

	// Facebook 平台
	FacebookMediaXVendor MediaXVendor = "Facebook"

	// Facebook 相关应用
	Facebook      VendorApp = "Facebook"      // Facebook
	FacebookAd    VendorApp = "FacebookAd"    // Facebook 广告
	FacebookLive  VendorApp = "FacebookLive"  // Facebook 直播
	FacebookReels VendorApp = "FacebookReels" // Facebook Reels 短视频

	// Instagram 平台
	InstagramMediaXVendor MediaXVendor = "Instagram"

	// Instagram 相关应用
	Instagram      VendorApp = "Instagram"      // Instagram
	InstagramAd    VendorApp = "InstagramAd"    // Instagram 广告
	InstagramReels VendorApp = "InstagramReels" // Instagram Reels

	// Twitter（X）平台
	TwitterMediaXVendor MediaXVendor = "Twitter"

	// Twitter 相关应用
	Twitter       VendorApp = "Twitter"       // Twitter
	TwitterAd     VendorApp = "TwitterAd"     // Twitter 广告
	TwitterSpaces VendorApp = "TwitterSpaces" // Twitter Spaces（音频直播）

	// LinkedIn 平台
	LinkedInMediaXVendor MediaXVendor = "LinkedIn"

	// LinkedIn 相关应用
	LinkedIn     VendorApp = "LinkedIn"     // LinkedIn
	LinkedInAd   VendorApp = "LinkedInAd"   // LinkedIn 广告
	LinkedInLive VendorApp = "LinkedInLive" // LinkedIn 直播

	// Twitch 平台
	TwitchMediaXVendor MediaXVendor = "Twitch"

	// Twitch 相关应用
	Twitch     VendorApp = "Twitch"     // Twitch
	TwitchAd   VendorApp = "TwitchAd"   // Twitch 广告
	TwitchLive VendorApp = "TwitchLive" // Twitch 直播

	// Snapchat 平台
	SnapchatMediaXVendor MediaXVendor = "Snapchat"

	// Snapchat 相关应用
	Snapchat          VendorApp = "Snapchat"          // Snapchat
	SnapchatAd        VendorApp = "SnapchatAd"        // Snapchat 广告
	SnapchatSpotlight VendorApp = "SnapchatSpotlight" // Snapchat Spotlight 短视频
)
