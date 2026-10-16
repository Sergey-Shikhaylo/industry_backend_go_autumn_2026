package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
	"uuid"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type TaskRepo interface {
	Create(title string) (Task, error)
	Get(id string) (Task, bool)
	List(done bool) []Task
	SetDone(id string, done bool) (Task, error)
}
type Clock interface{ Now() time.Time }

var (
	ErrNotFound     = errors.New("task not found")
	ErrInvalidTitle = errors.New("invalid title")
)

type inMemoryTaskRepo struct {
	mu    sync.RWMutex
	clock Clock
	seq   uint64
	tasks map[string]Task
}

func NewInMemoryTaskRepo(clock Clock) TaskRepo {
	return &inMemoryTaskRepo{
		clock: clock,
		tasks: make(map[string]Task),
	}
}
func (r *inMemoryTaskRepo) Create(title string) (Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	trimTitle := strings.TrimSpace(title)
	if trimTitle == "" {
		return Task{}, ErrInvalidTitle
	}
	newTask := Task{
		ID:        uuid.New().String(),
		Title:     trimTitle,
		Done:      false,
		UpdatedAt: r.clock.Now(),
	}
	r.tasks[newTask.ID] = newTask
	return newTask, nil
}

func (r *inMemoryTaskRepo) Get(id string) (Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	t, ok := r.tasks[id]
	if !ok {
		return Task{}, false
	}
	return t, true
}

func (r *inMemoryTaskRepo) List(done bool) []Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Task, 0, len(r.tasks))
	for _, val := range r.tasks {
		result = append(result, val)
	}
	return result
}

func (r *inMemoryTaskRepo) SetDone(id string, done bool) (Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	if t.Done != done {
		t.Done = done
		t.UpdatedAt = r.clock.Now()
		r.tasks[id] = t
	}
	return t, nil
}

type httpHandler struct{ repo TaskRepo }

func NewHTTPHandler(repo TaskRepo) http.Handler {
	panic("TODO: implement")
}
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}
func (h *httpHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}
func (h *httpHandler) handleGet(w http.ResponseWriter, r *http.Request, id string) {
	panic("TODO: implement")
}
func (h *httpHandler) handleList(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement")
}
func (h *httpHandler) handlePatch(w http.ResponseWriter, r *http.Request, id string) {
	panic("TODO: implement")
}
func decodeStrictJSON(r io.Reader, v any) error {
	panic("TODO: implement")
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	panic("TODO: implement")
}
