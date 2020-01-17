package httperrors

import "device-manager/pkg/models"

var EmailTakenError = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "001-001",
	Description: "Email is already taken",
}}

var UnauthorizedError = models.ErrorResult{Error: &models.ErrorResultError{
	Code:        "001-002",
	Description: "Access restricted",
}}