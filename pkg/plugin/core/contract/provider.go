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
	WechatOfficialAccount PluginMediaXProvider = "PluginMediaXWechatOfficialAccount"
	WechatMiniProgram     PluginMediaXProvider = "PluginMediaXWechatMiniProgram"
	WechatMoments         PluginMediaXProvider = "PluginMediaXWechatMoments"
	WechatVideo           PluginMediaXProvider = "PluginMediaXWechatVideo"
	WechatLive            PluginMediaXProvider = "PluginMediaXWechatLive"

	// YouTube
	YouTubeChannel  PluginMediaXProvider = "PluginMediaXYouTubeChannel"
	YouTubeVideo    PluginMediaXProvider = "PluginMediaXYouTubeVideo"
	YouTubePlaylist PluginMediaXProvider = "PluginMediaXYouTubePlaylist"
	YouTubeLive     PluginMediaXProvider = "PluginMediaXYouTubeLive"

	// Instagram
	InstagramFeed  PluginMediaXProvider = "PluginMediaXInstagramFeed"
	InstagramStory PluginMediaXProvider = "PluginMediaXInstagramStory"
	InstagramPost  PluginMediaXProvider = "PluginMediaXInstagramPost"
	InstagramVideo PluginMediaXProvider = "PluginMediaXInstagramVideo"

	// Facebook
	FacebookPage  PluginMediaXProvider = "PluginMediaXFacebookPage"
	FacebookPost  PluginMediaXProvider = "PluginMediaXFacebookPost"
	FacebookLive  PluginMediaXProvider = "PluginMediaXFacebookLive"
	FacebookGroup PluginMediaXProvider = "PluginMediaXFacebookGroup"

	// TikTok
	TikTokVideo PluginMediaXProvider = "PluginMediaXTikTokVideo"
	TikTokLive  PluginMediaXProvider = "PluginMediaXTikTokLive"
	TikTokDuet  PluginMediaXProvider = "PluginMediaXTikTokDuet"

	// LinkedIn
	LinkedInPost    PluginMediaXProvider = "PluginMediaXLinkedInPost"
	LinkedInArticle PluginMediaXProvider = "PluginMediaXLinkedInArticle"
	LinkedInVideo   PluginMediaXProvider = "PluginMediaXLinkedInVideo"

	// Snapchat
	SnapchatStory     PluginMediaXProvider = "PluginMediaXSnapchatStory"
	SnapchatSpotlight PluginMediaXProvider = "PluginMediaXSnapchatSpotlight"

	// Pinterest
	PinterestPin   PluginMediaXProvider = "PluginMediaXPinterestPin"
	PinterestBoard PluginMediaXProvider = "PluginMediaXPinterestBoard"

	// Vimeo
	VimeoVideo PluginMediaXProvider = "PluginMediaXVimeoVideo"
	VimeoLive  PluginMediaXProvider = "PluginMediaXVimeoLive"

	// Reddit
	RedditPost    PluginMediaXProvider = "PluginMediaXRedditPost"
	RedditComment PluginMediaXProvider = "PluginMediaXRedditComment"
)
