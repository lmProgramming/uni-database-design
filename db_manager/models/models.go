package models

import (
	"time"

	"gorm.io/gorm"
)

// create your tables here...

type Author struct {
	gorm.Model
	AuthorType AuthorType `gorm:"type:author_type_enum;not null"`

	Comments      []Comment      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Posts         []Post         `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Reactions     []Reaction     `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Messages      []Message      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Conversations []Conversation `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Reels         []Reel         `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Events        []Event        `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Groups        []Group        `gorm:"many2many:group_members;constraint:OnDelete:CASCADE"`
}

type Group struct {
	gorm.Model
	Name    string    `gorm:"not null;size:100"`
	Members []*Author `gorm:"many2many:group_members;constraint:OnDelete:CASCADE"`
}

type User struct {
	AuthorID          uint                `gorm:"primaryKey"`
	FirstName         string              `gorm:"not null;size:50"`
	SecondName        string              `gorm:"not null;size:50"`
	Email             string              `gorm:"not null;size:100;unique"`
	Password          string              `gorm:"not null"`
	PictureUrl        *string             `gorm:"size:255"`
	BackgroundUrl     *string             `gorm:"size:255"`
	Birthday          *time.Time          `gorm:"type:date"`
	IsVerified        bool                `gorm:"default:false"`
	Bio               string              `gorm:"default:'Edit bio';size:160"`
	ExternalUserLinks []ExternalUserLinks `gorm:"foreignKey:AuthorID"`
	Friends           []User              `gorm:"many2many:user_friends;"`
	FriendRequests    []FriendRequest     `gorm:"foreignKey:ReceiverID"`
	UserPrivilegeID   uint                `gorm:"not null"`

	Author Author `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
}

type UserPrivilege struct {
	gorm.Model
	Users         []User `gorm:"foreignKey:UserPrivilegeID"`
	PrivilegeName string `gorm:"not null;unique;size:40"`
}

type ExternalUserLinks struct {
	gorm.Model
	AuthorID uint   `gorm:"not null"`
	Platform string `gorm:"not null"`
	Link     string `gorm:"not null"`
}

type Tag struct {
	gorm.Model
	PageID  uint
	TagName string  `gorm:"not null;unique;size:100"`
	Pages   []*Page `gorm:"many2many:page_tags"`
}

type FriendRequest struct {
	gorm.Model
	SenderID   uint   `gorm:"not null"`
	ReceiverID uint   `gorm:"not null"`
	Status     string `gorm:"type:friend_request_status;default:'pending'"`

	Sender   User `gorm:"foreignKey:SenderID;references:AuthorID;constraint:OnDelete:CASCADE"`
	Receiver User `gorm:"foreignKey:ReceiverID;references:AuthorID;constraint:OnDelete:CASCADE"`
}

type Comment struct {
	gorm.Model
	AuthorID uint
	Author   Author `gorm:"foreignKey:AuthorID"`

	Content string `gorm:"not null"`

	Hashtags []*Hashtag `gorm:"many2many:comment_hashtags"`
}

type Post struct {
	gorm.Model
	AuthorID uint `gorm:"not null"`
	Title    string
	Content  string `gorm:"not null"`
	IsPublic bool   `gorm:"default:true"`

	Location   *Location `gorm:"foreignKey:LocationID;references:ID"`
	LocationID uint

	Hashtags []*Hashtag `gorm:"many2many:post_hashtags"`
}

type Location struct {
	gorm.Model
	City       string `gorm:"size:100"`
	Country    string `gorm:"size:100"`
	PostalCode string `gorm:"size:20"`

	Geolocation   *Geolocation `gorm:"foreignKey:GeolocationID;references:ID"`
	GeolocationID uint

	Address   *Address `gorm:"foreignKey:AddressID;references:ID"`
	AddressID uint
}

type Geolocation struct {
	gorm.Model
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`
}

type Address struct {
	gorm.Model
	StreetName string `gorm:"not null;size:255"`
	Building   string `gorm:"size:20"`
	Gate       string `gorm:"size:20"`
	Floor      string `gorm:"size:20"`
	Apartment  string `gorm:"size:20"`
}

type Event struct {
	gorm.Model
	AuthorID    uint       `gorm:"not null"`
	Name        string     `gorm:"not null;size:300"`
	Description string     `gorm:"size:1024"`
	StartDate   *time.Time `gorm:"not null;type:date"`
	EndDate     *time.Time `gorm:"not null;type:date"`

	Members []*Author `gorm:"many2many:event_members"`

	Location   *Location `gorm:"foreignKey:LocationID;references:ID"`
	LocationID uint
}

type Reel struct {
	gorm.Model
	AuthorID uint   `gorm:"not null"`
	Content  string `gorm:"not null"`
}

type Message struct {
	gorm.Model
	Content        string `gorm:"not null"`
	AuthorID       uint   `gorm:"not null"`
	ConversationID uint   `gorm:"not null"`
}

type Conversation struct {
	gorm.Model
	Title    string `gorm:"not null"`
	IconUrl  string
	AuthorID uint      `gorm:"not null"`
	Members  []*Author `gorm:"many2many:conversation_members"`
}

type Page struct {
	gorm.Model
	AuthorID uint
	Author   Author `gorm:"foreignKey:AuthorID"`

	Title          string `gorm:"not null;size:100"`
	Tags           []*Tag `gorm:"many2many:page_tags"`
	Advertisements []*Advertisement
	Views          uint `gorm:"default:0"`
	Likes          uint `gorm:"default:0"`
}

type Advertisement struct {
	gorm.Model
	Content string `gorm:"not null"`
	AdLink  string `gorm:"not null"`
	PageID  uint   `gorm:"null"`
}

type Reaction struct {
	gorm.Model
	AuthorID uint   `gorm:"primaryKey"`
	PostID   uint   `gorm:"primaryKey"`
	Reaction string `gorm:"size:20;not null"`
}

type Hashtag struct {
	gorm.Model
	TagName string `gorm:"not null;unique;size:200"`
}
