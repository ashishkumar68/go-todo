package model

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	createIdentityError = errors.New("Could not create new Identity.")
)

type Identity struct {
	ID    int32      `gorm:"name:id;primarykey;AUTO_INCREMENT" json:"id"`
	Uuid  uuid.UUID   `gorm:"name:uuid;type:binary(16);unique_index" json:"uuid"`
}

func GetNewIdentity(newUuid []byte) (*Identity, error) {
	if len(newUuid) == 0 {
		nUuid, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(createIdentityError.Error())
			fmt.Println(err)
			return nil, err
		}
		return &Identity{Uuid: nUuid}, nil
	} else {
		nUuid, err := uuid.FromBytes(newUuid)
		if err != nil {
			fmt.Println(createIdentityError.Error())
			fmt.Println(err)
			return nil, err
		}
		return &Identity{Uuid: nUuid}, nil
	}
}

type CreatedMetaInfo struct {
	CreatedAt time.Time `gorm:"name:created_at;type:datetime;autoCreateTime:milli" json:"createdAt"`
	CreatedBy *int32 `gorm:"name:created_by"`
	CreatedByUser *User `gorm:"foreignkey:CreatedBy"`
}

type UpdatedMetaInfo struct {
	UpdatedAt time.Time `gorm:"name:updated_at;type:datetime;autoUpdateTime:milli" json:"updatedAt"`
	UpdatedBy *int32 `gorm:"name:updated_by"`
	UpdatedByUser *User `gorm:"foreignkey:UpdatedBy"`
}

type DeletedMetaInfo struct {
	DeletedAt time.Time `gorm:"name:deleted_at;type:datetime" json:"deletedAt"`
	DeletedBy *int32 `gorm:"name:deleted_by"`
	DeletedByUser *User `gorm:"foreignkey:DeletedBy"`
}

type ActionedMetaInfo struct {
	ActionedAt time.Time `gorm:"name:actioned_at;type:datetime" json:"actionedAt"`
	CycledAt time.Time `gorm:"name:cycled_at;type:datetime" json:"cycledAt"`
}

type RequestedMetaInfo struct {
	RequestedAt time.Time `gorm:"name:requested_at;type:datetime" json:"requestedAt"`
}

type VersionMetaInfo struct {
	Version int32 `gorm:"name:version;" json:"version"`
}