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

// ---------------------------------------------
// 定义Provider
// ---------------------------------------------

// PluginMediaXProvider 枚举定义
type PluginMediaXProvider string

const (
	PluginMediaX PluginMediaXProvider = "PluginMediaX"

	// WeChat
	WechatOfficialAccount PluginMediaXProvider = "WechatOfficialAccount"
	WechatMiniProgram     PluginMediaXProvider = "WechatMiniProgram"
	WechatMoments         PluginMediaXProvider = "WechatMoments"
	WechatVideo           PluginMediaXProvider = "WechatVideo"
	WechatLive            PluginMediaXProvider = "WechatLive"

	// YouTube
	YouTubeChannel  PluginMediaXProvider = "YouTubeChannel"
	YouTubeVideo    PluginMediaXProvider = "YouTubeVideo"
	YouTubePlaylist PluginMediaXProvider = "YouTubePlaylist"
	YouTubeLive     PluginMediaXProvider = "YouTubeLive"

	// Instagram
	InstagramFeed  PluginMediaXProvider = "InstagramFeed"
	InstagramStory PluginMediaXProvider = "InstagramStory"
	InstagramPost  PluginMediaXProvider = "InstagramPost"
	InstagramVideo PluginMediaXProvider = "InstagramVideo"

	// Facebook
	FacebookPage  PluginMediaXProvider = "FacebookPage"
	FacebookPost  PluginMediaXProvider = "FacebookPost"
	FacebookLive  PluginMediaXProvider = "FacebookLive"
	FacebookGroup PluginMediaXProvider = "FacebookGroup"

	// TikTok
	TikTokVideo PluginMediaXProvider = "TikTokVideo"
	TikTokLive  PluginMediaXProvider = "TikTokLive"
	TikTokDuet  PluginMediaXProvider = "TikTokDuet"

	// LinkedIn
	LinkedInPost    PluginMediaXProvider = "LinkedInPost"
	LinkedInArticle PluginMediaXProvider = "LinkedInArticle"
	LinkedInVideo   PluginMediaXProvider = "LinkedInVideo"

	// Snapchat
	SnapchatStory     PluginMediaXProvider = "SnapchatStory"
	SnapchatSpotlight PluginMediaXProvider = "SnapchatSpotlight"

	// Pinterest
	PinterestPin   PluginMediaXProvider = "PinterestPin"
	PinterestBoard PluginMediaXProvider = "PinterestBoard"

	// Vimeo
	VimeoVideo PluginMediaXProvider = "VimeoVideo"
	VimeoLive  PluginMediaXProvider = "VimeoLive"

	// Reddit
	RedditPost    PluginMediaXProvider = "RedditPost"
	RedditComment PluginMediaXProvider = "RedditComment"
)
