package repository

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

// CategoryRepository is responsible of getting/saving information from DB.
type CategoryRepository struct {
	DB *gorm.DB
}

// Get returns all categories.
func (r CategoryRepository) Get() (categories []model.Category, apiError *errors.ApiError) {
	result := r.DB.Find(&categories)
	if err := result.Error; err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.NotFoundMessage,
			Error:   err.Error(),
		}
	}

	return
}

// Find return one category from DB by ID.
func (r CategoryRepository) Find(id string) (category model.Category, apiError *errors.ApiError) {
	result := r.DB.Model(model.Category{}).Where("id = ?", id).First(&category)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return model.Category{}, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.NotFoundMessage,
				Error:   err.Error(),
			}
		}

		return model.Category{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.NotFoundMessage,
			Error:   "category not found",
		}
	}

	return
}

// Create persist a category to the DB.
func (r CategoryRepository) Create(category model.Category) (uint, *errors.ApiError) {
	result := r.DB.Create(&category)
	if err := result.Error; err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailedMessage,
			Error:   err.Error(),
		}
	}

	return category.ID, nil
}

// Update update an existent category.
func (r CategoryRepository) Update(id string, upCategory model.Category) (model.Category, *errors.ApiError) {
	category, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailedMessage
		return model.Category{}, apiError
	}

	result := r.DB.Model(&category).Updates(upCategory)
	if err := result.Error; err != nil {
		return model.Category{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailedMessage,
			Error:   err.Error(),
		}
	}

	return category, nil
}

// Delete delete an existent category from DB.
func (r CategoryRepository) Delete(id string) (apiError *errors.ApiError) {
	category, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailedMessage
		return
	}

	err := r.DB.Delete(&category).Error
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailedMessage,
			Error:   err.Error(),
		}
	}

	return
}
