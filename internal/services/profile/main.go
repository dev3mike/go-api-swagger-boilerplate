package profileservice

import (
	"context"
	"database/sql"
	"time"

	"github.com/dev3mike/go-api-swagger-boilerplate/internal/database"
	"github.com/dev3mike/go-api-swagger-boilerplate/internal/entities"
	internalerrors "github.com/dev3mike/go-api-swagger-boilerplate/pkg/constants/internal_errors"
)

func GetProfile(username string) (*entities.ProfileEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var profile entities.ProfileEntity
	err := database.DB.
		NewSelect().
		Model(&profile).
		Where("username = ?", username).
		Limit(1).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return &entities.ProfileEntity{}, internalerrors.ErrNotFound
		}
		return &entities.ProfileEntity{}, err
	}

	return &profile, nil
}

func CreateProfile(profileEntity *entities.ProfileEntity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.DB.
		NewInsert().
		Model(profileEntity).
		Exec(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return internalerrors.ErrNotFound
		}
		return err
	}

	return nil
}

func UpdateProfile(username string, profileEntity *entities.ProfileEntity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.DB.NewUpdate().Model(profileEntity).Where("username = ?", username).Returning("NULL").Exec(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return internalerrors.ErrNotFound
		}
		return err
	}

	return nil
}

func UpdateProfilePartiallyByExternalId(externalId string, profileEntity *entities.ProfileEntity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	timeNow := time.Now()
	profileEntity.Updated = &timeNow

	_, err := database.DB.
		NewUpdate().
		Column("private_email", "profile_image", "fullname", "updated").
		Model(profileEntity).
		Where("external_id = ?", externalId).
		Returning("NULL").
		Exec(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return internalerrors.ErrNotFound
		}
		return err
	}

	return nil
}
