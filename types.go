package telebot

import "encoding/json"

type WebhookUpdate struct {
	UpdateID           int            `json:"update_id"`
	Message            *Message       `json:"message"`
	EditedMessage      *Message       `json:"edited_message"`
	ChannelPost        *Message       `json:"channel_post"`
	EditedChannelPost  *Message       `json:"edited_channel_post"`
	InlineQuery        interface{}    `json:"inline_query"`
	ChosenInlineResult interface{}    `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery `json:"callback_query"`
	ShippingQuery      interface{}    `json:"shipping_query"`
	PreCheckoutQuery   interface{}    `json:"pre_checkout_query"`
}

type SetWebhook struct {
	Url            string   `json:"url"`
	Certificate    []byte   `json:"certificate"`
	MaxConnections int      `json:"max_connections"`
	AllowedUpdates []string `json:"allowed_updates"`
}

type WebhookInfo struct {
	Url                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        int      `json:"last_error_date"`
	LastErrorMessage     string   `json:"last_error_message"`
	MaxConnections       int      `json:"max_connections"`
	AllowedUpdates       []string `json:"allowed_updates"`
}

type User struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	Id                          int        `json:"id"`
	Type                        string     `json:"type"`
	Title                       string     `json:"title"`
	Username                    string     `json:"username"`
	FirstName                   string     `json:"first_name"`
	LastName                    string     `json:"last_name"`
	AllMembersAreAdministrators bool       `json:"all_members_are_administrators"`
	Photo                       *ChatPhoto `json:"photo"`
	Description                 string     `json:"description"`
	InviteLink                  string     `json:"invite_link"`
	PinnedMessage               *Message   `json:"pinned_message"`
	StickerSetName              string     `json:"sticker_set_name"`
	CanSetStickerSet            bool       `json:"can_set_sticker_set"`
}

type Message struct {
	MessageID             int              `json:"message_id"`
	From                  *User            `json:"from"`
	Date                  int              `json:"date"`
	Chat                  *Chat            `json:"chat"`
	ForwardFrom           *User            `json:"forward_from"`
	ForwardFromChat       *Chat            `json:"forward_from_chat"`
	ForwardFromMessageID  int              `json:"forward_from_message_id"`
	ForwardSignature      string           `json:"forward_signature"`
	ForwardDate           int              `json:"forward_date"`
	ReplyToMessage        *Message         `json:"reply_to_message"`
	EditDate              int              `json:"edit_date"`
	MediaGroupID          string           `json:"media_group_id"`
	AuthorSignature       string           `json:"author_signature"`
	Text                  string           `json:"text"`
	Entities              []*MessageEntity `json:"entities"`
	CaptionEntities       []*MessageEntity `json:"caption_entities"`
	Audio                 *Audio           `json:"audio"`
	Document              *Document        `json:"document"`
	Animation             *Animation       `json:"animation"`
	Game                  interface{}      `json:"game"`
	Photo                 []*PhotoSize     `json:"photo"`
	Sticker               *Sticker         `json:"sticker"`
	Video                 *Video           `json:"video"`
	Voice                 *Voice           `json:"voice"`
	VideoNote             *VideoNote       `json:"video_note"`
	Caption               string           `json:"caption"`
	Contact               *Contact         `json:"contact"`
	Location              *Location        `json:"location"`
	Venue                 *Venue           `json:"venue"`
	NewChatMemebers       []*User          `json:"new_chat_memebers"`
	LeftChatMember        *User            `json:"left_chat_member"`
	NewChatTitle          string           `json:"new_chat_title"`
	NewChatPhoto          []*PhotoSize     `json:"new_chat_photo"`
	DeleteChatPhoto       bool             `json:"delete_chat_photo"`
	GroupChatCreated      bool             `json:"group_chat_created"`
	SuperGroupChatCreated bool             `json:"supergroup_chat_created"`
	ChannelChatCreated    bool             `json:"channel_chat_created"`
	MigrateToChatID       int              `json:"migrate_to_chat_id"`
	MigrateFromChatID     int              `json:"migrate_from_chat_id"`
	PinnedMessage         *Message         `json:"pinned_message"`
	Invoice               interface{}      `json:"invoice"`
	SuccessfullPayment    interface{}      `json:"successfull_payment"`
	ConnectedWebSite      string           `json:"connected_web_site"`
	PassportData          interface{}      `json:"passport_data"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	OffSet int    `json:"off_set"`
	Lenght int    `json:"lenght"`
	Url    string `json:"url"`
	User   *User  `json:"user"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type Audio struct {
	FileID    string     `json:"file_id"`
	Duration  int        `json:"duration"`
	Performer string     `json:"performer"`
	Title     string     `json:"title"`
	MimeType  string     `json:"mime_type"`
	FileSize  int        `json:"file_size"`
	Thumb     *PhotoSize `json:"thumb"`
}

type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

type Animation struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}
type VideoNote struct {
	FileID   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	FileSize int        `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
	VCard       string `json:"vcard"`
}

type Location struct {
	Longtitude float64 `json:"longtitude"`
	Latitude   float64 `json:"latitude"`
}

type Venue struct {
	Location       *Location `json:"location"`
	Title          string    `json:"title"`
	Address        string    `json:"address"`
	FourSquareId   string    `json:"four_square_id"`
	FourSquareType string    `json:"four_square_type"`
}

type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}

type ReplyKeyboardMarkup struct {
	Keyboard       [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard bool                `json:"resize_keyboard"`
	OnTimeKeyboard bool                `json:"on_time_keyboard"`
	Selective      bool                `json:"selective"`
}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

type InlineKeyboardButton struct {
	Text                         string      `json:"text"`
	Url                          string      `json:"url"`
	CallbackData                 string      `json:"callback_data"`
	SwitchInlineQuery            string      `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string      `json:"switch_inline_query_current_chat"`
	CallbackGame                 interface{} `json:"callback_game"`
	Pay                          bool        `json:"pay"`
}

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

type ForceReplay struct {
	ForceReplay bool `json:"force_replay"`
	Selective   bool `json:"selective"`
}

type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	UntilDate             int    `json:"until_date"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanRestrictMemberes   bool   `json:"can_restrict_memberes"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
}

type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}

type InputMediaPhoto struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

type InputMediaVideo struct {
	Type              string `json:"type"`
	Media             string `json:"media"`
	Thumb             []byte `json:"thumb"`
	Caption           string `json:"caption"`
	ParseMode         string `json:"parse_mode"`
	Width             int    `json:"width"`
	Height            int    `json:"height"`
	Duration          int    `json:"duration"`
	SupportsStreaming bool   `json:"supports_streaming"`
}

type InputMediaAnimation struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Thumb     []byte `json:"thumb"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Duration  int    `json:"duration"`
}

type InputMediaAudio struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Thumb     []byte `json:"thumb"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
}

type InputMediaDocument struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Thumb     []byte `json:"thumb"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

type Sticker struct {
	FileID       string        `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        string        `json:"emoji"`
	SetName      string        `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int           `json:"file_size"`
}

type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []*Sticker `json:"stickers"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type APIResponse struct {
	Ok          bool            `json:"ok"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
	Result      json.RawMessage `json:"result"`
}

//
type (
	FilePath struct {
		FileId   string `json:"file_id"`
		FileSize int    `json:"file_size"`
		FilePath string `json:"file_path"`
	}
	InlineKeyboardMarkup struct {
		InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	}

	EditMessageText struct {
		ChatID                int         `json:"chat_id"`
		MessageID             int         `json:"message_id"`
		InlineMessageID       string      `json:"inline_message_id"`
		Text                  string      `json:"text"`
		ParseMode             string      `json:"parse_mode"`
		DisableWebPagePreview bool        `json:"disable_web_page_preview"`
		ReplyMarkup           interface{} `json:"reply_markup"`
	}
	DeleteMessage struct {
		ChatID    int `json:"chat_id"`
		MessageID int `json:"message_id"`
	}
)
