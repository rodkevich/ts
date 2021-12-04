package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/rodkevich/ts/photo/internal/models"
)

type photoPG struct {
	db *pgxpool.Pool
}

func NewPhotoPG(db *pgxpool.Pool) *photoPG {
	return &photoPG{db: db}
}

func (p *photoPG) CreatePhoto(ctx context.Context, arg models.CreatePhotoParams) (*models.Photo, error) {
	const createPhoto = `
	INSERT INTO photos
	(type, size_kb, upload_name, description, image_url, owner_id)
	VALUES
	($1, $2, $3, $4, $5, $6)
	RETURNING
	id, type, size_kb, upload_name, image_url, description, owner_id, created_at, updated_at, deleted
	`
	row := p.db.QueryRow(
		ctx, createPhoto,
		arg.Type, arg.SizeKb, arg.UploadName,
		arg.Description, arg.ImageUrl, arg.OwnerID,
	)

	var i models.Photo
	err := row.Scan(
		&i.ID, &i.Type, &i.SizeKb, &i.UploadName,
		&i.ImageUrl, &i.Description, &i.OwnerID,
		&i.CreatedAt, &i.UpdatedAt, &i.Deleted,
	)
	return &i, err
}

func (p *photoPG) GetPhoto(ctx context.Context, id uuid.UUID) (*models.Photo, error) {
	const getPhoto = `
	SELECT
	id, type, size_kb, upload_name, image_url, description, owner_id, created_at, updated_at, deleted
	FROM photos
	WHERE id = $1
	LIMIT 1
	`
	row := p.db.QueryRow(ctx, getPhoto, id)

	var i models.Photo
	err := row.Scan(
		&i.ID, &i.Type, &i.SizeKb, &i.UploadName,
		&i.ImageUrl, &i.Description, &i.OwnerID,
		&i.CreatedAt, &i.UpdatedAt, &i.Deleted,
	)
	return &i, err
}

func (p *photoPG) ListPhotos(ctx context.Context) (*models.PhotoList, error) {
	const listPhotos = `
	SELECT
	id, type, size_kb, upload_name, image_url, description, owner_id, created_at, updated_at, deleted
	FROM photos
	ORDER BY updated_at DESC
	`
	rows, err := p.db.Query(ctx, listPhotos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rtn := make([]*models.Photo, 0)
	for rows.Next() {
		var each models.Photo
		if err := rows.Scan(
			&each.ID, &each.Type, &each.SizeKb, &each.UploadName, &each.ImageUrl,
			&each.Description, &each.OwnerID, &each.CreatedAt, &each.UpdatedAt, &each.Deleted,
		); err != nil {
			return nil, err
		}
		rtn = append(rtn, &each)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &models.PhotoList{Photos: rtn}, nil
}

func (p *photoPG) UpdatePhoto(ctx context.Context, arg models.UpdatePhotoParams, id uuid.UUID) (*models.Photo, error) {
	const updatePhoto = `
	UPDATE photos
	SET type=$1::enum_photos_type,
		size_kb=$2,
		upload_name=$3::citext,
		image_url=$4,
		description=$5::citext,
		owner_id=$6
	WHERE id = $7
	RETURNING
	id, type, size_kb, upload_name, image_url, description, owner_id, created_at, updated_at, deleted
	`
	row := p.db.QueryRow(
		ctx, updatePhoto,
		arg.Type, arg.SizeKb, arg.UploadName,
		arg.ImageUrl, arg.Description, arg.OwnerID,
		id,
	)

	var i models.Photo
	err := row.Scan(
		&i.ID, &i.Type, &i.SizeKb, &i.UploadName, &i.ImageUrl,
		&i.Description, &i.OwnerID, &i.CreatedAt, &i.UpdatedAt,
		&i.Deleted,
	)
	return &i, err
}

func (p *photoPG) MarkPhotoAsDeleted(ctx context.Context, deleted bool, id uuid.UUID) (*models.Photo, error) {
	const markPhotoAsDeleted = `
	UPDATE photos
	SET deleted=$1
	WHERE id = $2
	RETURNING
	id, type, size_kb, upload_name, image_url, description, owner_id, created_at, updated_at, deleted
	`
	row := p.db.QueryRow(
		ctx, markPhotoAsDeleted,
		deleted, id,
	)

	var i models.Photo
	err := row.Scan(
		&i.ID, &i.Type, &i.SizeKb, &i.UploadName,
		&i.ImageUrl, &i.Description, &i.OwnerID,
		&i.CreatedAt, &i.UpdatedAt, &i.Deleted,
	)
	return &i, err
}

func (p *photoPG) DeletePhoto(ctx context.Context, id uuid.UUID) error {
	const deletePhoto = `
	DELETE
	FROM photos
	WHERE id = $1
	`
	_, err := p.db.Exec(ctx, deletePhoto, id)
	return err
}
