package repository

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/models"
	"fmt"
	"github.com/satori/go.uuid"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	selectComplaintQuery         = "SELECT * FROM complaints WHERE id = $1 FOR UPDATE"
	updateComplaintStatusQuery   = "UPDATE complaints SET stage = $1, updated_at = NOW() WHERE id = $2"
	insertHistoryQuery           = "INSERT INTO reports_history (report_id, old_stage, new_stage, admin_comment) VALUES (:report_id, :old_stage, :new_stage, :admin_comment)"
	deleteCommentQuery           = "DELETE FROM comments WHERE id = ? AND complaint_id = ?"
	updateComplaintPriorityQuery = `UPDATE complaints SET priority = $1, updated_at = NOW() WHERE id = $2 RETURNING updated_at`
)

const (
	defaultOffset = 0
	defaultLimit  = 10
)

type ComplaintsRepository struct {
	db            *sqlx.DB
	Authorization Authorization
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{
		db:            db,
		Authorization: NewAuthPostgres(db),
	}
}

func (r *ComplaintsRepository) FindUsers(UserUUID uuid.UUID) (*entity.Users, error) {
	var user entity.Users

	if UserUUID == uuid.Nil {
		return nil, fmt.Errorf("user_uuid is required")
	}

	const query = `SELECT user_uuid, username, email, role, phone FROM users WHERE user_uuid = $1`
	row := r.db.QueryRow(query, UserUUID.String())

	err := row.Scan(
		&user.UserUUID,
		&user.UserName,
		&user.Email,
		&user.Role,
		&user.Phone,
	)

	if err != nil {
		return nil, fmt.Errorf("user_uuid not found")
	}

	return &user, nil // Возвращаем указатель на найденного пользователя
}

func (r *ComplaintsRepository) UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error) {
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
func (r *ComplaintsRepository) DeleteComment(complaintID string, commentID string) error {
	result, err := r.db.Exec(deleteCommentQuery, commentID, complaintID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return models.ErrCommentNotFound
	}

	return nil
}

func (r *ComplaintsRepository) UpdateComplaintPriority(id string, priority string) (time.Time, error) {
	var updatedAt time.Time
	err := r.db.QueryRow(updateComplaintPriorityQuery, priority, id).Scan(&updatedAt)
	if err != nil {
		return time.Time{}, err
	}
	return updatedAt, nil
}
