package repositories

import (
	"fmt"
	"time"

	"go-goland-api/main.go/domain"

	"github.com/jmoiron/sqlx"
)

type TaskRepository interface {
	GetTask(id uint) (*domain.Task, error)
	createTask(task *domain.Task) error
}

type taskRepository struct {
	conn *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return &taskRepository{
		conn: db,
	}
}

func (repo *taskRepository) GetTask(id uint) (*domain.Task, error) {
	book := new(domain.Task)
	err := repo.conn.Get(Task, "SELECT * FROM tasks WHERE id=?", id)
	if err != nil {
		return nil, fmt.Errorf("error getting task: %w", err)
	}

	return book, nil
}

func (repo *taskRepository) CreateTask(task *domain.Task) error {
	createdAt := time.Now()

	result, err := repo.conn.Exec(`INSERT INTO books 
		(title, author, price, stock, isbn, created_at, updated_at) 
		VALUES(?,?,?,?,?,?, ?)`, createdAt, createdAt)

	if err != nil {
		return fmt.Errorf("error inserting task: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error saving task: %w", err)
	}

	task.ID = int(id)

	return nil
}
