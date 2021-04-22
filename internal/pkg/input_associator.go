package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

// AssociateBookInput is responsible of associate the params to the book model.
func AssociateBookInput(c *gin.Context) (book model.Book, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&book)
	if err != nil {
		return model.Book{}, &errors.ApiError{
			Status:  http.StatusUnprocessableEntity,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateCategoryInput is responsible of associating the params to the category model.
func AssociateCategoryInput(c *gin.Context) (category model.Category, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return model.Category{}, &errors.ApiError{
			Status:  http.StatusUnprocessableEntity,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateStudentInput is responsible of associating the params to the student model.
func AssociateStudentInput(c *gin.Context) (student model.Student, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&student)
	if err != nil {
		return model.Student{}, &errors.ApiError{
			Status:  http.StatusUnprocessableEntity,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateLendingInput is responsible of associating the params to the lending model.
func AssociateLendingInput(c *gin.Context) (lending model.Lending, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&lending)
	if err != nil {
		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusUnprocessableEntity,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateAccountInput is responsible of associating the params to the account model.
func AssociateAccountInput(c *gin.Context) (account model.Account, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&account)
	if err != nil {
		return model.Account{}, &errors.ApiError{
			Status:  http.StatusUnprocessableEntity,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateLoginInput is responsible of associating the params to the account model.
func AssociateLoginInput(c *gin.Context) (login model.Credential, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&login)
	if err != nil {
		return model.Credential{}, &errors.ApiError{
			Status:  http.StatusUnprocessableEntity,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
