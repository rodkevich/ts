package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rodkevich/ts/ticket/internal/models"

	"github.com/google/uuid"
)

type tagPG struct {
	db *pgxpool.Pool
}

func NewTagPG(db *pgxpool.Pool) *tagPG {
	return &tagPG{db: db}
}

func (q *tagPG) CreateTag(ctx context.Context, name string, description *string) (models.Tag, error) {
	const createTag = `
	INSERT INTO tags (name, description)
	VALUES ($1, $2)
	RETURNING 
	id, name, description, created_at, updated_at, deleted
	`
	row := q.db.QueryRow(ctx, createTag, name, description)
	var i models.Tag
	err := row.Scan(
		&i.ID, &i.Name, &i.Description,
		&i.CreatedAt, &i.UpdatedAt, &i.Deleted,
	)
	return i, err
}

func (q *tagPG) DeleteTag(ctx context.Context, id uuid.UUID) error {
	const deleteTag = `
	DELETE FROM tags
	WHERE id = $1
	`
	_, err := q.db.Exec(ctx, deleteTag, id)
	return err
}

func (q *tagPG) GetTag(ctx context.Context, id uuid.UUID) (models.Tag, error) {
	const getTag = `
	SELECT 
	id, name, description, created_at, updated_at, deleted FROM tags
	WHERE id = $1
	LIMIT 1
	`
	row := q.db.QueryRow(ctx, getTag, id)
	var i models.Tag
	err := row.Scan(
		&i.ID, &i.Name, &i.Description,
		&i.CreatedAt, &i.UpdatedAt, &i.Deleted,
	)
	return i, err
}

func (q *tagPG) ListTags(ctx context.Context) (*models.TagList, error) {
	const listTags = `
	SELECT 
	id, name, description, created_at, updated_at, deleted FROM tags
	ORDER BY name
	`
	rows, err := q.db.Query(ctx, listTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*models.Tag, 0)
	for rows.Next() {
		var i models.Tag
		if err := rows.Scan(
			&i.ID, &i.Name, &i.Description,
			&i.CreatedAt, &i.UpdatedAt, &i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &models.TagList{Tags: items}, nil
}
