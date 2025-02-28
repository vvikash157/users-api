package service

import (
	"Login/db"
	"Login/models"
	"errors"
)

func CreateTask(task models.Task) (*models.Task, error) {
	if err := db.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func GetTask(page, pageSize int, status string, task models.Task) ([]models.Task, int64, error) {
	var tasks []models.Task
	var totalRecords int64

	query := db.DB.Model(&models.Task{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&totalRecords)

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}

	return tasks, totalRecords, nil
}

func GetTaskById(taskID string, task models.Task) (*models.Task, error) {
	if err := db.DB.First(&task, taskID).Error; err != nil {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

func UpdateTask(taskID string, updatedTask models.Task) (*models.Task, error) {
	var task models.Task
	if err := db.DB.First(&task, taskID).Error; err != nil {
		return nil, errors.New("task not found")
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status

	if err := db.DB.Save(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func DeleteTask(taskID int) error {
	if err := db.DB.Delete(&models.Task{}, taskID).Error; err != nil {
		return err
	}
	log.Println("Task deleted successfully, ID:", taskID)
	return nil
}
