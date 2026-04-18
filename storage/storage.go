package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AuraReaper/taskman/models"
)

type Store struct {
	filePath string
}

func NewStore(path string) (*Store, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	return &Store{
		filePath: path,
	}, nil
}

func (s *Store) writeAll(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

func (s *Store) Create(task models.Task) error {
	tasks, err := s.ReadAll()
	if err != nil {
		return err
	}

	tasks = append(tasks, task)
	return s.writeAll(tasks)
}

func (s *Store) ReadAll() ([]models.Task, error) {
	file, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if len(file) == 0 {
		return []models.Task{}, nil
	}

	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func (s *Store) Delete(id int) error {
	tasks, err := s.ReadAll()
	if err != nil {
		return err
	}

	isIDPresent := false

	var newTasks []models.Task
	for _, task := range tasks {
		if task.ID == id {
			isIDPresent = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if isIDPresent == false {
		return fmt.Errorf("the task with ID %v not present\n", id)
	}

	return s.writeAll(newTasks)
}

func (s *Store) UpdateStatus(id int) error {
	tasks, err := s.ReadAll()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "COMPLETED"
		}
	}

	return s.writeAll(tasks)
}

func (s *Store) GetNextId() (int, error) {
	file, err := os.ReadFile(s.filePath)
	if err != nil {
		return -1, err
	}

	if len(file) == 0 {
		return 1, nil
	}

	var tasks []models.Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return -1, err
	}

	maxId := 0

	for _, task := range tasks {
		if task.ID > maxId {
			maxId = task.ID
		}
	}

	return maxId + 1, nil
}
