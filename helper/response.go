package helper

import "net/http"

func BadRequest(msg string) (int, map[string]interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"status":  "error",
		"message": msg,
	}
}

func SiccessCreate() (int, map[string]interface{}) {
	return http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"status":  "success",
		"message": "insert success",
	}
}

func FailedBadRequest() (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"message": "bad request",
		"code":    400,
	}
}

func AuthOK(data interface{}, token string) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    200,
		"data":    data,
		"token":   token,
	}
}

func FailedNotFound() (int, interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"message": "data not found",
		"code":    404,
	}
}
