package contract

import "context"

// ---------------------------------------------
// Provider 定义插件需要实现的接口
// ---------------------------------------------
type ProviderInterface interface {
	// Initialize 插件初始化方法
	Initialize(ctx *context.Context, config interface{}) error
	// Name 返回插件名称
	Name(ctx *context.Context) string
	// Publish 发布内容的方法
	Publish(ctx *context.Context, req interface{}) (interface{}, error)
}

// 定义平台类型（Provider）
type MediaVendor string
type AppPlugin string

const (
	// 微信生态（WeChat）平台
	WechatMediaVendor MediaVendor = "WeChat"

	// 微信相关插件
	WechatOfficialAccount AppPlugin = "WechatOfficialAccount" // 公众号
	WechatMiniProgram     AppPlugin = "WechatMiniProgram"     // 小程序
	WechatChannel         AppPlugin = "WechatChannel"         // 视频号
	WechatWork            AppPlugin = "WechatWork"            // 企业微信

	// 抖音生态（DouYin）平台
	DouYinMediaVendor MediaVendor = "DouYin"

	// 抖音插件
	DouYin AppPlugin = "DouYin" // 抖音

	// 小红书生态（RedBook）平台
	RedBookMediaVendor MediaVendor = "RedBook"

	// 小红书插件
	RedBook AppPlugin = "RedBook" // 小红书

	// YouTube平台
	YouTubeMediaVendor MediaVendor = "YouTube"

	// YouTube插件
	YouTube AppPlugin = "YouTube" // YouTube
)

// **内容类型（Content Type）**
type PluginMediaXContentType string

const (
	ContentTypeVideo  PluginMediaXContentType = "Video"  // 视频
	ContentTypeLive   PluginMediaXContentType = "Live"   // 直播
	ContentTypePost   PluginMediaXContentType = "Post"   // 文章/图文
	ContentTypeStory  PluginMediaXContentType = "Story"  // 短时动态（如 Instagram Story）
	ContentTypeReels  PluginMediaXContentType = "Reels"  // 短视频（如 Facebook Reels）
	ContentTypeDuet   PluginMediaXContentType = "Duet"   // TikTok/抖音对拍
	ContentTypeStatus PluginMediaXContentType = "Status" // 类似 WhatsApp 状态，微信 Moment（朋友圈）
)
