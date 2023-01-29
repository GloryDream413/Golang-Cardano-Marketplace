package models

import (
	"time"
	"github.com/google/uuid"
)

type Collection struct {
	ID      			uuid.UUID	`gorm:"type:uuid;primary_key"`
	Name	 			string		`gorm:"type:string"`
	Description			string		`gorm:"type:string"`
	Creator 			string		`gorm:"type:string"`
	RecipientAddr		string		`gorm:"type:string"`
	FeeType				string		`gorm:"type:string"`
	Fee					string		`gorm:"type:string"`
	Type				string		`gorm:"type:string"`
	CollectionIpfs		string		`gorm:"type:string"`
	RepresentativeIpfs	string		`gorm:"type:string"`
	Count				string		`gorm:"type:string"`
	MintedInfo			string		`gorm:"type:string"`
	Price				string		`gorm:"type:string"`
	CreatedAt			time.Time	`gorm:"not null"`
}