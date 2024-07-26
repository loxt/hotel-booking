package guest

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loxt/hotel-booking/booking/internal/domain/entities"
)

type GuestRepository struct {
	db *pgxpool.Pool
}

func NewGuestRepository(db *pgxpool.Pool) *GuestRepository {
	return &GuestRepository{
		db: db,
	}
}

func (r *GuestRepository) Get(id string) (*entities.Guest, error) {
	sql := `SELECT id, name, surname, email, document_id, document_type FROM guests WHERE id = $1`

	var guest entities.Guest

	if err := r.db.QueryRow(context.Background(), sql, id).Scan(&guest.Id, &guest.Name, &guest.Surname, &guest.Email, &guest.DocumentID.Value, &guest.DocumentID.DocumentType); err != nil {
		return nil, err
	}

	return &guest, nil
}

func (r *GuestRepository) Create(guest *entities.Guest) error {
	sql := `INSERT INTO guests (id, name, surname, email, document_id, document_type) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id string

	if err := r.db.QueryRow(context.Background(), sql, guest.Id, guest.Name, guest.Surname, guest.Email, guest.DocumentID.Value, guest.DocumentID.DocumentType).Scan(&id); err != nil {
		return err
	}

	return nil
}
