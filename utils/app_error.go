package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type: %d, code: %s, err: %s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

// universal error

func ReqBodyNotValidError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Bad request",
		ErrorType:    http.StatusBadRequest,
	}
}

func CreateIdError() error {
	return AppError{
		ErrorCode:    "500",
		ErrorMessage: "Created Id Failed",
		ErrorType:    http.StatusInternalServerError,
	}
}

//invoice

func GetInvoiceError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Get Invoice Error",
		ErrorType:    http.StatusBadRequest,
	}
}

// item

func GetItemError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "No Item Found",
		ErrorType:    http.StatusBadRequest,
	}
}

func CreateItemsError() error {
	return AppError{
		ErrorCode:    "409",
		ErrorMessage: "Create Item Failed",
		ErrorType:    http.StatusConflict,
	}
}

func DuplicateItemError() error {
	return AppError{
		ErrorCode:    "409",
		ErrorMessage: "Item's already exist",
		ErrorType:    http.StatusConflict,
	}
}

// customer

func DuplicateCustomer() error {
	return AppError{
		ErrorCode:    "409",
		ErrorMessage: "Customer already exist",
		ErrorType:    http.StatusConflict,
	}
}

func CreateCustomerError() error {
	return AppError{
		ErrorCode:    "409",
		ErrorMessage: "Create Customer Failed",
		ErrorType:    http.StatusConflict,
	}
}

func GetCustomerError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "No Customer Found",
		ErrorType:    http.StatusConflict,
	}
}
