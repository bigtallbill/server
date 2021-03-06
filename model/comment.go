package model

import "time"

type Comment struct {
	Model
	TopicID     uint         `json:"topicId" sql:"index"`
	Topic       Topic        `json:"topic" gorm:"association_autoupdate:false;association_autocreate:false"`
	UserID      uint         `json:"userId" sql:"index"`
	User        User         `json:"user" gorm:"association_autoupdate:false;association_autocreate:false"`
	Body        string       `json:"body" sql:"type:text"`
	Attachments []Attachment `json:"attachments" gorm:"association_autoupdate:false;association_autocreate:false"`
	Rank        uint         `json:"rank"`
	Solution    bool         `json:"solution"`
}

type Attachment struct {
	Model
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Thumbnail   string `json:"thumbnail"`
	Html        string `json:"html"`
}

type CommentStatus struct {
	Model
	CommentID  uint       `json:"commentId"`
	Comment    Topic      `json:"comment" gorm:"association_autoupdate:false;association_autocreate:false"`
	UserID     uint       `json:"userId" sql:"index"`
	User       User       `json:"user" gorm:"association_autoupdate:false;association_autocreate:false"`
	ReadAt     *time.Time `json:"readAt"`
	NotifiedAt *time.Time `json:"notifiedAt"`
	Vote       int        `json:"vote"`
}
