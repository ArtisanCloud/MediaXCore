package config

// **内容类型（Content Type）**
type MediaXContentType string

const (
	ContentTypeVideo  MediaXContentType = "Video"  // 视频
	ContentTypeLive   MediaXContentType = "Live"   // 直播
	ContentTypePost   MediaXContentType = "Post"   // 文章/图文
	ContentTypeStory  MediaXContentType = "Story"  // 短时动态（如 Instagram Story）
	ContentTypeReels  MediaXContentType = "Reels"  // 短视频（如 Facebook Reels）
	ContentTypeDuet   MediaXContentType = "Duet"   // TikTok/抖音对拍
	ContentTypeStatus MediaXContentType = "Status" // 类似 WhatsApp 状态，微信 Moment（朋友圈）
)
