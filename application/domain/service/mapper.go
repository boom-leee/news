package service

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	inport "news-api/application/port/in"
	"time"

	"news-api/internal/db"
)

func MapUser(user db.User) *inport.User {
	return &inport.User{
		ID:        ToUUID(user.ID).String(),
		AuthID:    user.AuthID,
		Email:     user.Email.String,
		Name:      user.Name.String,
		Role:      user.Role.String,
		ImageUrl:  user.ImageUrl.String,
		CreatedAt: PGTimestampToTime(user.CreatedAt),
		UpdatedAt: PGTimestampToTime(user.UpdatedAt),
		DeletedAt: PGTimestampToTime(user.DeletedAt),
	}
}

func MapCategory(category db.Category) *inport.Category {
	return &inport.Category{
		ID:        ToUUID(category.ID).String(),
		Name:      category.Name.String,
		CreatedAt: PGTimestampToTime(category.CreatedAt),
		UpdatedAt: PGTimestampToTime(category.UpdatedAt),
		DeletedAt: PGTimestampToTime(category.DeletedAt),
	}
}

func ToUUID(id pgtype.UUID) uuid.UUID {
	if !id.Valid {
		return uuid.Nil
	}
	return id.Bytes
}

func PGTimestampToTime(timestamp pgtype.Timestamp) *time.Time {
	if !timestamp.Valid {
		return nil
	}
	return &timestamp.Time
}
