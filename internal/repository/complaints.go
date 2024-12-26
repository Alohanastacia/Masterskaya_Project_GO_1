package repository

import (
	"complaint_service/internal/entity"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	selectComplaintQuery       = "SELECT * FROM complaints WHERE id = $1 FOR UPDATE"
	updateComplaintStatusQuery = "UPDATE complaints SET stage = $1, updated_at = NOW() WHERE id = $2"
	insertHistoryQuery         = "INSERT INTO reports_history (report_id, old_stage, new_stage, admin_comment) VALUES (:report_id, :old_stage, :new_stage, :admin_comment)"
)

const (
	defaultOffset = 0
	defaultLimit  = 10
)

type ComplaintsDB struct {
	db *sqlx.DB
}

type ComplaintsRepository struct {
	Authorization
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{
		Authorization: NewAuthPostgres(db),
	}
}

func (rep *ComplaintsDB) FindUsers(UserUUID string, limit, offset int) ([]*entity.Users, error) {

	var user entity.Users

	if limit <= 0 {
		limit = defaultLimit
	}
	if offset < 0 {
		offset = defaultOffset
	}

	const query = `SELECT user_uuid, username, email, role, phone
					FROM users 
					WHERE user_uuid = ?
					ORDER BY user_uuid
					LIMIT ? OFFSET ?`

	if UserUUID == "" {
		return nil, fmt.Errorf("user_uuid is required")
	}
	rows := rep.db.QueryRow(query, UserUUID, limit, offset)

	err := rows.Scan(
		&user.UserUUID,
		&user.UserName,
		&user.Email,
		&user.Role,
		&user.Phone,
	)
	if user.Role != entity.Admin {
		return nil, fmt.Errorf("access errors, insufficient rights")
	}
	if err != nil {
		return nil, fmt.Errorf("user_uuid not found")
	}

	return []*entity.Users{&user}, nil
}

func (r *ComplaintsDB) UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error) {
	var complaint entity.Complaint

	tx, err := r.db.Beginx()
	if err != nil {
		return time.Time{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Get(&complaint, selectComplaintQuery, id); err != nil {
		return time.Time{}, err
	}

	oldStage := complaint.Stage

	if _, err = tx.Exec(updateComplaintStatusQuery, status, id); err != nil {
		return time.Time{}, err
	}

	history := entity.ReportsHistory{
		ReportID:     complaint.ID,
		OldStage:     oldStage,
		NewStage:     status,
		AdminComment: adminComment,
	}

	if _, err = tx.NamedExec(insertHistoryQuery, history); err != nil {
		return time.Time{}, err
	}

	if err = tx.Commit(); err != nil {
		return time.Time{}, err
	}

	return time.Time{}, nil
}

// Ниже будут методы ComplaintsRepository, которые делают запросы в БД и отдают результат
