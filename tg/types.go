package tg

type User struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code,omitempty"`
	IsPremium    bool   `json:"is_premium,omitempty"`
}

type Chat struct {
	Id        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsForum   bool   `json:"is_forum,omitempty"`
}

type MessageOrigin struct {
	// Type of the message origin, always “user”
	Type string `json:"type"`
	// Date the message was sent originally in Unix time
	Date int64 `json:"date"`
	// User that sent the message originally
	SenderUser User `json:"sender_user"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	Url          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard []InlineKeyboardButton `json:"inline_keyboard"`
}

type Message struct {
	// Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageId int64 `json:"message_id"`
	// Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	MessageThreadId int64 `json:"message_thread_id"`
	// Optional. Sender of the message; may be empty for messages sent to channels. For backward compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in non-channel chats
	From User `json:"from"`
	// Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	Date int64 `json:"date"`
	// Chat the message belongs to
	Chat Chat `json:"chat"`
	// Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type MessagePayload struct {
	ChatId              int64                `json:"chat_id"`
	MessageThreadId     string               `json:"message_thread_id,omitempty"`
	Text                string               `json:"text"`
	ParseMode           string               `json:"parse_mode,omitempty"` // HTML or MarkdownV2
	DisableNotification bool                 `json:"disable_notification,omitempty"`
	ProtectContent      bool                 `json:"protect_content,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type MessageResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int    `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}
