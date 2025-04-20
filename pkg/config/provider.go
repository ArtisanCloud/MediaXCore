package config

// 定义平台类型（Provider）
type MediaXVendor string
type VendorApp string

const (
	// 微信生态（WeChat）平台
	WechatMediaXVendor MediaXVendor = "we_chat"

	// 微信相关应用
	WechatOfficialAccount VendorApp = "wechat_official_account" // 公众号
	WechatMiniProgram     VendorApp = "wechat_mini_program"     // 小程序
	WechatChannel         VendorApp = "wechat_channel"          // 视频号
	WechatWork            VendorApp = "wechat_work"             // 企业微信

	// 抖音生态（DouYin）平台
	DouYinMediaXVendor MediaXVendor = "dou_yin"

	// 抖音相关应用
	DouYin      VendorApp = "dou_yin"       // 抖音
	DouYinStore VendorApp = "dou_yin_store" // 抖音小店
	DouYinLive  VendorApp = "dou_yin_live"  // 抖音直播
	DouYinAd    VendorApp = "dou_yin_ad"    // 抖音广告

	// 小红书生态（RedBook）平台
	RedBookMediaXVendor MediaXVendor = "red_book"

	// 小红书相关应用
	RedBook      VendorApp = "red_book"       // 小红书
	RedBookStore VendorApp = "red_book_store" // 小红书商城
	RedBookAd    VendorApp = "red_book_ad"    // 小红书广告

	// 快手生态（Kuaishou）平台
	KuaishouMediaXVendor MediaXVendor = "kuaishou"

	// 快手相关应用
	Kuaishou      VendorApp = "kuaishou"       // 快手
	KuaishouLive  VendorApp = "kuaishou_live"  // 快手直播
	KuaishouStore VendorApp = "kuaishou_store" // 快手小店
	KuaishouAd    VendorApp = "kuaishou_ad"    // 快手广告

	// Bilibili 生态（Bilibili）平台
	BilibiliMediaXVendor MediaXVendor = "bilibili"

	// Bilibili 相关应用
	Bilibili     VendorApp = "bilibili"      // Bilibili
	BilibiliLive VendorApp = "bilibili_live" // Bilibili 直播
	BilibiliAd   VendorApp = "bilibili_ad"   // Bilibili 广告

	// 知乎生态（Zhihu）平台
	ZhihuMediaXVendor MediaXVendor = "zhihu"

	// 知乎相关应用
	Zhihu     VendorApp = "zhihu"      // 知乎
	ZhihuAd   VendorApp = "zhihu_ad"   // 知乎广告
	ZhihuLive VendorApp = "zhihu_live" // 知乎 Live

	// 微博生态（Weibo）平台
	WeiboMediaXVendor MediaXVendor = "weibo"

	// 微博相关应用
	Weibo   VendorApp = "weibo"    // 微博
	WeiboAd VendorApp = "weibo_ad" // 微博广告

	// YouTube 平台
	GoogleMediaXVendor MediaXVendor = "google"

	// YouTube 相关应用
	YouTube       VendorApp = "youtube"        // YouTube
	YouTubeAd     VendorApp = "youtube_ad"     // YouTube 广告
	YouTubeShorts VendorApp = "youtube_shorts" // YouTube Shorts
	Blogger       VendorApp = "blogger"        // 博客平台

	// TikTok 平台
	TikTokMediaXVendor MediaXVendor = "tiktok"

	// TikTok 相关应用
	TikTok     VendorApp = "tiktok"      // TikTok
	TikTokAd   VendorApp = "tiktok_ad"   // TikTok 广告
	TikTokShop VendorApp = "tiktok_shop" // TikTok 小店
	TikTokLive VendorApp = "tiktok_live" // TikTok 直播

	// Facebook 平台
	FacebookMediaXVendor MediaXVendor = "facebook"

	// Facebook 相关应用
	Facebook      VendorApp = "facebook"       // Facebook
	FacebookAd    VendorApp = "facebook_ad"    // Facebook 广告
	FacebookLive  VendorApp = "facebook_live"  // Facebook 直播
	FacebookReels VendorApp = "facebook_reels" // Facebook Reels 短视频

	// Instagram 平台
	InstagramMediaXVendor MediaXVendor = "instagram"

	// Instagram 相关应用
	Instagram      VendorApp = "instagram"       // Instagram
	InstagramAd    VendorApp = "instagram_ad"    // Instagram 广告
	InstagramReels VendorApp = "instagram_reels" // Instagram Reels

	// Twitter（X）平台
	TwitterMediaXVendor MediaXVendor = "twitter"

	// Twitter 相关应用
	Twitter       VendorApp = "twitter"        // Twitter
	TwitterAd     VendorApp = "twitter_ad"     // Twitter 广告
	TwitterSpaces VendorApp = "twitter_spaces" // Twitter Spaces（音频直播）

	// LinkedIn 平台
	LinkedInMediaXVendor MediaXVendor = "linked_in"

	// LinkedIn 相关应用
	LinkedIn     VendorApp = "linked_in"      // LinkedIn
	LinkedInAd   VendorApp = "linked_in_ad"   // LinkedIn 广告
	LinkedInLive VendorApp = "linked_in_live" // LinkedIn 直播

	// Twitch 平台
	TwitchMediaXVendor MediaXVendor = "twitch"

	// Twitch 相关应用
	Twitch     VendorApp = "twitch"      // Twitch
	TwitchAd   VendorApp = "twitch_ad"   // Twitch 广告
	TwitchLive VendorApp = "twitch_live" // Twitch 直播

	// Snapchat 平台
	SnapchatMediaXVendor MediaXVendor = "snapchat"

	// Snapchat 相关应用
	Snapchat          VendorApp = "snapchat"           // Snapchat
	SnapchatAd        VendorApp = "snapchat_ad"        // Snapchat 广告
	SnapchatSpotlight VendorApp = "snapchat_spotlight" // Snapchat Spotlight 短视频
)
