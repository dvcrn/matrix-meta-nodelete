package table

type LSInsertStickerAttachment struct {
	PlayableUrl                      string `index:"0" json:",omitempty"`
	PlayableUrlFallback              string `index:"1" json:",omitempty"`
	PlayableUrlExpirationTimestampMs int64  `index:"2" json:",omitempty"`
	PlayableUrlMimeType              string `index:"3" json:",omitempty"`
	PreviewUrl                       string `index:"4" json:",omitempty"`
	PreviewUrlFallback               string `index:"5" json:",omitempty"`
	PreviewUrlExpirationTimestampMs  int64  `index:"6" json:",omitempty"`
	PreviewUrlMimeType               string `index:"7" json:",omitempty"`
	PreviewWidth                     int64  `index:"9" json:",omitempty"`
	PreviewHeight                    int64  `index:"10" json:",omitempty"`
	ImageUrlMimeType                 string `index:"11" json:",omitempty"`
	AttachmentIndex                  int64  `index:"12" json:",omitempty"`
	AccessibilitySummaryText         string `index:"13" json:",omitempty"`
	ThreadKey                        int64  `index:"14" json:",omitempty"`
	TimestampMs                      int64  `index:"17" json:",omitempty"`
	MessageId                        string `index:"18" json:",omitempty"`
	AttachmentFbid                   string `index:"19" json:",omitempty"`
	ImageUrl                         string `index:"20" json:",omitempty"`
	ImageUrlFallback                 string `index:"21" json:",omitempty"`
	ImageUrlExpirationTimestampMs    int64  `index:"22" json:",omitempty"`
	FaviconUrlExpirationTimestampMs  int64  `index:"23" json:",omitempty"`
	AvatarViewSize                   int64  `index:"25" json:",omitempty"`
	AvatarCount                      int64  `index:"26" json:",omitempty"`
	TargetId                         int64  `index:"27" json:",omitempty"`
	MustacheText                     string `index:"30" json:",omitempty"`
	AuthorityLevel                   int64  `index:"31" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSInsertXmaAttachment struct {
	Filename                                   string         `index:"1" json:",omitempty"`
	Filesize                                   int64          `index:"2" json:",omitempty"`
	IsSharable                                 bool           `index:"3" json:",omitempty"`
	PlayableUrl                                string         `index:"4" json:",omitempty"`
	PlayableUrlFallback                        string         `index:"5" json:",omitempty"`
	PlayableUrlExpirationTimestampMs           int64          `index:"6" json:",omitempty"`
	PlayableUrlMimeType                        string         `index:"7" json:",omitempty"`
	PreviewUrl                                 string         `index:"8" json:",omitempty"`
	PreviewUrlFallback                         string         `index:"9" json:",omitempty"`
	PreviewUrlExpirationTimestampMs            int64          `index:"10" json:",omitempty"`
	PreviewUrlMimeType                         string         `index:"11" json:",omitempty"`
	PreviewWidth                               int64          `index:"13" json:",omitempty"`
	PreviewHeight                              int64          `index:"14" json:",omitempty"`
	AttributionAppId                           int64          `index:"15" json:",omitempty"`
	AttributionAppName                         string         `index:"16" json:",omitempty"`
	AttributionAppIcon                         int64          `index:"17" json:",omitempty"`
	AttributionAppIconFallback                 string         `index:"18" json:",omitempty"`
	AttributionAppIconUrlExpirationTimestampMs int64          `index:"19" json:",omitempty"`
	AttachmentIndex                            int64          `index:"20" json:",omitempty"`
	AccessibilitySummaryText                   string         `index:"21" json:",omitempty"`
	ShouldRespectServerPreviewSize             bool           `index:"22" json:",omitempty"`
	SubtitleIconUrl                            string         `index:"23" json:",omitempty"`
	ShouldAutoplayVideo                        bool           `index:"24" json:",omitempty"`
	ThreadKey                                  int64          `index:"25" json:",omitempty"`
	AttachmentType                             AttachmentType `index:"27" json:",omitempty"`
	TimestampMs                                int64          `index:"29" json:",omitempty"`
	MessageId                                  string         `index:"30" json:",omitempty"`
	OfflineAttachmentId                        int64          `index:"31" json:",omitempty"`
	AttachmentFbid                             string         `index:"32" json:",omitempty"`
	XmaLayoutType                              int64          `index:"33" json:",omitempty"`
	XmasTemplateType                           int64          `index:"34" json:",omitempty"`
	CollapsibleId                              int64          `index:"35" json:",omitempty"`
	DefaultCtaId                               int64          `index:"36" json:",omitempty"`
	DefaultCtaTitle                            string         `index:"37" json:",omitempty"`
	DefaultCtaType                             int64          `index:"38" json:",omitempty"`
	AttachmentCta1Id                           int64          `index:"40" json:",omitempty"`
	Cta1Title                                  int64          `index:"41" json:",omitempty"`
	Cta1IconType                               int64          `index:"42" json:",omitempty"`
	Cta1Type                                   int64          `index:"43" json:",omitempty"`
	AttachmentCta2Id                           int64          `index:"45" json:",omitempty"`
	Cta2Title                                  int64          `index:"46" json:",omitempty"`
	Cta2IconType                               int64          `index:"47" json:",omitempty"`
	Cta2Type                                   int64          `index:"48" json:",omitempty"`
	AttachmentCta3Id                           int64          `index:"50" json:",omitempty"`
	Cta3Title                                  string         `index:"51" json:",omitempty"`
	Cta3IconType                               int64          `index:"52" json:",omitempty"`
	Cta3Type                                   int64          `index:"53" json:",omitempty"`
	ImageUrl                                   string         `index:"54" json:",omitempty"`
	ImageUrlFallback                           string         `index:"55" json:",omitempty"`
	ImageUrlExpirationTimestampMs              int64          `index:"56" json:",omitempty"`
	ActionUrl                                  string         `index:"57" json:",omitempty"`
	TitleText                                  string         `index:"58" json:",omitempty"`
	SubtitleText                               string         `index:"59" json:",omitempty"`
	SubtitleDecorationType                     int64          `index:"60" json:",omitempty"`
	MaxTitleNumOfLines                         int64          `index:"61" json:",omitempty"`
	MaxSubtitleNumOfLines                      int64          `index:"62" json:",omitempty"`
	DescriptionText                            string         `index:"63" json:",omitempty"`
	SourceText                                 string         `index:"64" json:",omitempty"`
	FaviconUrl                                 string         `index:"65" json:",omitempty"`
	FaviconUrlFallback                         string         `index:"66" json:",omitempty"`
	FaviconUrlExpirationTimestampMs            int64          `index:"67" json:",omitempty"`
	ListItemsId                                int64          `index:"69" json:",omitempty"`
	ListItemsDescriptionText                   string         `index:"70" json:",omitempty"`
	ListItemsDescriptionSubtitleText           string         `index:"71" json:",omitempty"`
	ListItemsSecondaryDescriptionText          string         `index:"72" json:",omitempty"`
	ListItemId1                                int64          `index:"73" json:",omitempty"`
	ListItemTitleText1                         string         `index:"74" json:",omitempty"`
	ListItemContactUrlList1                    int64          `index:"75" json:",omitempty"`
	ListItemProgressBarFilledPercentage1       int64          `index:"76" json:",omitempty"`
	ListItemContactUrlExpirationTimestampList1 int64          `index:"77" json:",omitempty"`
	ListItemContactUrlFallbackList1            int64          `index:"78" json:",omitempty"`
	ListItemAccessibilityText1                 string         `index:"79" json:",omitempty"`
	ListItemTotalCount1                        int64          `index:"80" json:",omitempty"`
	ListItemId2                                int64          `index:"81" json:",omitempty"`
	ListItemTitleText2                         int64          `index:"82" json:",omitempty"`
	ListItemContactUrlList2                    int64          `index:"83" json:",omitempty"`
	ListItemProgressBarFilledPercentage2       int64          `index:"84" json:",omitempty"`
	ListItemContactUrlExpirationTimestampList2 int64          `index:"85" json:",omitempty"`
	ListItemContactUrlFallbackList2            int64          `index:"86" json:",omitempty"`
	ListItemAccessibilityText2                 int64          `index:"87" json:",omitempty"`
	ListItemTotalCount2                        int64          `index:"88" json:",omitempty"`
	ListItemId3                                int64          `index:"89" json:",omitempty"`
	ListItemTitleText3                         int64          `index:"90" json:",omitempty"`
	ListItemContactUrlList3                    int64          `index:"91" json:",omitempty"`
	ListItemProgressBarFilledPercentage3       int64          `index:"92" json:",omitempty"`
	ListItemContactUrlExpirationTimestampList3 int64          `index:"93" json:",omitempty"`
	ListItemContactUrlFallbackList3            int64          `index:"94" json:",omitempty"`
	ListItemAccessibilityText3                 int64          `index:"95" json:",omitempty"`
	ListItemTotalCount3                        int64          `index:"96" json:",omitempty"`
	IsBorderless                               bool           `index:"100" json:",omitempty"`
	HeaderImageUrlMimeType                     string         `index:"101" json:",omitempty"`
	HeaderTitle                                string         `index:"102" json:",omitempty"`
	HeaderSubtitleText                         string         `index:"103" json:",omitempty"`
	HeaderImageUrl                             string         `index:"104" json:",omitempty"`
	HeaderImageUrlFallback                     string         `index:"105" json:",omitempty"`
	HeaderImageUrlExpirationTimestampMs        int64          `index:"106" json:",omitempty"`
	PreviewImageDecorationType                 int64          `index:"107" json:",omitempty"`
	ShouldHighlightHeaderTitleInTitle          bool           `index:"108" json:",omitempty"`
	TargetId                                   int64          `index:"109" json:",omitempty"`
	XMATypeOne                                 string         `index:"110" json:",omitempty"`
	XMATypeTwo                                 string         `index:"111" json:",omitempty"`
	AttachmentLoggingType                      int64          `index:"112" json:",omitempty"`
	PreviewUrlLarge                            string         `index:"114" json:",omitempty"`
	GatingType                                 int64          `index:"115" json:",omitempty"`
	GatingTitle                                string         `index:"116" json:",omitempty"`
	TargetExpiryTimestampMs                    int64          `index:"117" json:",omitempty"`
	CountdownTimestampMs                       int64          `index:"118" json:",omitempty"`
	ShouldBlurSubattachments                   int64          `index:"119" json:",omitempty"`
	VerifiedType                               int64          `index:"120" json:",omitempty"`
	CaptionBodyText                            int64          `index:"121" json:",omitempty"`
	IsPublicXma                                bool           `index:"122" json:",omitempty"`
	ReplyCount                                 int64          `index:"123" json:",omitempty"`
	AuthorityLevel                             int64          `index:"124" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSInsertBlobAttachment struct {
	Filename                                   string         `index:"0" json:",omitempty"`
	Filesize                                   int64          `index:"1" json:",omitempty"`
	HasMedia                                   bool           `index:"2" json:",omitempty"`
	PlayableUrl                                string         `index:"3" json:",omitempty"`
	PlayableUrlFallback                        string         `index:"4" json:",omitempty"`
	PlayableUrlExpirationTimestampMs           int64          `index:"5" json:",omitempty"`
	PlayableUrlMimeType                        string         `index:"6" json:",omitempty"`
	DashManifest                               string         `index:"7" json:",omitempty"`
	PreviewUrl                                 string         `index:"8" json:",omitempty"`
	PreviewUrlFallback                         string         `index:"9" json:",omitempty"`
	PreviewUrlExpirationTimestampMs            int64          `index:"10" json:",omitempty"`
	PreviewUrlMimeType                         string         `index:"11" json:",omitempty"`
	UnknownExpirationTimestamp                 int64          `index:"12" json:",omitempty"`
	MiniPreview                                int64          `index:"13" json:",omitempty"`
	PreviewWidth                               int64          `index:"14" json:",omitempty"`
	PreviewHeight                              int64          `index:"15" json:",omitempty"`
	AttributionAppId                           int64          `index:"16" json:",omitempty"`
	AttributionAppName                         string         `index:"17" json:",omitempty"`
	AttributionAppIcon                         string         `index:"18" json:",omitempty"`
	AttributionAppIconFallback                 int64          `index:"19" json:",omitempty"`
	AttributionAppIconUrlExpirationTimestampMs int64          `index:"20" json:",omitempty"`
	LocalPlayableUrl                           int64          `index:"21" json:",omitempty"`
	PlayableDurationMs                         int64          `index:"22" json:",omitempty"`
	AttachmentIndex                            int64          `index:"23" json:",omitempty"`
	AccessibilitySummaryText                   int64          `index:"24" json:",omitempty"`
	IsPreviewImage                             bool           `index:"25" json:",omitempty"`
	OriginalFileHash                           string         `index:"26" json:",omitempty"`
	ThreadKey                                  int64          `index:"27" json:",omitempty"`
	AttachmentType                             AttachmentType `index:"29" json:",omitempty"`
	AttachmentMimeType                         string         `index:"30" json:",omitempty"`
	TimestampMs                                int64          `index:"31" json:",omitempty"`
	MessageId                                  string         `index:"32" json:",omitempty"`
	OfflineAttachmentId                        string         `index:"33" json:",omitempty"`
	AttachmentFbid                             string         `index:"34" json:",omitempty"`
	HasXma                                     bool           `index:"35" json:",omitempty"`
	XmaLayoutType                              int64          `index:"36" json:",omitempty"`
	XmasTemplateType                           int64          `index:"37" json:",omitempty"`
	TitleText                                  string         `index:"38" json:",omitempty"`
	SubtitleText                               string         `index:"39" json:",omitempty"`
	DescriptionText                            string         `index:"40" json:",omitempty"`
	SourceText                                 string         `index:"41" json:",omitempty"`
	FaviconUrlExpirationTimestampMs            int64          `index:"42" json:",omitempty"`
	IsBorderless                               bool           `index:"44" json:",omitempty"`
	PreviewUrlLarge                            string         `index:"45" json:",omitempty"`
	SamplingFrequencyHz                        int64          `index:"46" json:",omitempty"`
	WaveformData                               string         `index:"47" json:",omitempty"`
	AuthorityLevel                             int64          `index:"48" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSInsertAttachmentItem struct {
	AttachmentFbid                   string `index:"0" json:",omitempty"`
	AttachmentIndex                  int64  `index:"1" json:",omitempty"`
	ThreadKey                        int64  `index:"2" json:",omitempty"`
	MessageId                        string `index:"4" json:",omitempty"`
	OriginalPageSenderId             int64  `index:"7" json:",omitempty"`
	TitleText                        string `index:"8" json:",omitempty"`
	SubtitleText                     string `index:"9" json:",omitempty"`
	PlayableUrl                      string `index:"12" json:",omitempty"`
	PlayableUrlFallback              string `index:"13" json:",omitempty"`
	PlayableUrlExpirationTimestampMs int64  `index:"14" json:",omitempty"`
	PlayableUrlMimeType              string `index:"15" json:",omitempty"`
	DashManifest                     string `index:"16" json:",omitempty"`
	PreviewUrl                       string `index:"17" json:",omitempty"`
	PreviewUrlFallback               string `index:"18" json:",omitempty"`
	PreviewUrlExpirationTimestampMs  int64  `index:"19" json:",omitempty"`
	PreviewUrlMimeType               string `index:"20" json:",omitempty"`
	PreviewWidth                     int64  `index:"21" json:",omitempty"`
	PreviewHeight                    int64  `index:"22" json:",omitempty"`
	ImageUrl                         string `index:"23" json:",omitempty"`
	DefaultCtaId                     int64  `index:"24" json:",omitempty"`
	DefaultCtaTitle                  string `index:"25" json:",omitempty"`
	DefaultCtaType                   int64  `index:"26" json:",omitempty"`
	DefaultButtonType                int64  `index:"28" json:",omitempty"`
	DefaultActionUrl                 string `index:"29" json:",omitempty"`
	DefaultActionEnableExtensions    bool   `index:"30" json:",omitempty"`
	DefaultWebviewHeightRatio        int64  `index:"32" json:",omitempty"`
	AttachmentCta1Id                 int64  `index:"34" json:",omitempty"`
	Cta1Title                        string `index:"35" json:",omitempty"`
	Cta1IconType                     int64  `index:"36" json:",omitempty"`
	Cta1Type                         int64  `index:"37" json:",omitempty"`
	AttachmentCta2Id                 int64  `index:"39" json:",omitempty"`
	Cta2Title                        string `index:"40" json:",omitempty"`
	Cta2IconType                     int64  `index:"41" json:",omitempty"`
	Cta2Type                         int64  `index:"42" json:",omitempty"`
	AttachmentCta3Id                 int64  `index:"44" json:",omitempty"`
	Cta3Title                        string `index:"45" json:",omitempty"`
	Cta3IconType                     int64  `index:"46" json:",omitempty"`
	Cta3Type                         int64  `index:"47" json:",omitempty"`
	FaviconUrl                       string `index:"48" json:",omitempty"`
	FaviconUrlFallback               string `index:"49" json:",omitempty"`
	FaviconUrlExpirationTimestampMs  int64  `index:"50" json:",omitempty"`
	PreviewUrlLarge                  string `index:"51" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSGetFirstAvailableAttachmentCTAID struct {
	Unrecognized map[int]any `json:",omitempty"`
}

type LSInsertAttachmentCta struct {
	CtaId               int64  `index:"0" json:",omitempty"`
	AttachmentFbid      string `index:"1" json:",omitempty"`
	AttachmentIndex     int64  `index:"2" json:",omitempty"`
	ThreadKey           int64  `index:"3" json:",omitempty"`
	MessageId           string `index:"5" json:",omitempty"`
	Title               string `index:"6" json:",omitempty"`
	Type_               string `index:"7" json:",omitempty"`
	PlatformToken       int64  `index:"8" json:",omitempty"`
	ActionUrl           string `index:"9" json:",omitempty"`
	NativeUrl           string `index:"10" json:",omitempty"`
	UrlWebviewType      int64  `index:"11" json:",omitempty"`
	ActionContentBlob   string `index:"12" json:",omitempty"`
	EnableExtensions    bool   `index:"13" json:",omitempty"`
	ExtensionHeightType int64  `index:"14" json:",omitempty"`
	TargetId            int64  `index:"15" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSUpdateAttachmentItemCtaAtIndex struct {
	AttachmentFbid  string `index:"0" json:",omitempty"`
	Unknown         int64  `index:"1" json:",omitempty"`
	AttachmentCtaId int64  `index:"2" json:",omitempty"`
	CtaTitle        string `index:"3" json:",omitempty"`
	CtaType         string `index:"4" json:",omitempty"`
	Index           int64  `index:"5" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSUpdateAttachmentCtaAtIndexIgnoringAuthority struct {
	ThreadKey       int64  `index:"0" json:",omitempty"`
	MessageId       string `index:"1" json:",omitempty"`
	AttachmentFbid  string `index:"2" json:",omitempty"`
	AttachmentCtaId int64  `index:"3" json:",omitempty"`
	CtaTitle        string `index:"4" json:",omitempty"`
	CtaType         string `index:"5" json:",omitempty"`
	Index           int64  `index:"6" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSHasMatchingAttachmentCTA struct {
	ThreadKey      int64  `index:"0" json:",omitempty"`
	AttachmentFbid string `index:"1" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSUpsertLiveLocationSharer struct {
	ThreadKey        int64   `index:"0" json:",omitempty"`
	Sender           int64   `index:"1" json:",omitempty"`
	Latitude         float64 `index:"2" json:",omitempty"`
	Longitude        float64 `index:"3" json:",omitempty"`
	StartTimestampMS int64   `index:"4" json:",omitempty"`
	EndTimestampMS   int64   `index:"5" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSDeleteLiveLocationSharer struct {
	ThreadKey int64 `index:"0" json:",omitempty"`
	Sender    int64 `index:"1" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}

type LSUpdateSharedAlbumOnMessageRecall struct {
	ThreadKey int64  `index:"0" json:",omitempty"`
	MessageId string `index:"1" json:",omitempty"`

	Unrecognized map[int]any `json:",omitempty"`
}