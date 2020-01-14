package httperrors

import "device-manager/pkg/models"

var InternalServerError = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "000-001",
	Description: "Internal server error",
}}
