package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func apiResponse(ctx *gin.Context, data interface{}) {
	response := BaseResponse{
		Code:    0,
		Message: "Success",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func errorResponse(ctx *gin.Context, code int, message string) {
	response := BaseResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	ctx.JSON(code, response)
}

func parseDateTimeStrToNullTime(dateStr string) sql.NullTime {
	if dateStr == "" {
		return sql.NullTime{}
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  date,
		Valid: true,
	}
}

func parseEndOfDayNullTime(dateStr string) sql.NullTime {
	result := parseDateTimeStrToNullTime(dateStr)
	if result.Valid {
		date := result.Time
		endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())
		result.Time = endOfDay
	}

	return result
}

func sliceToString(s interface{}) string {
	switch s := s.(type) {
	case []string:
		return strings.Join(s, ",")
	case []int:
		strSlice := make([]string, len(s))
		for i, v := range s {
			strSlice[i] = strconv.Itoa(v)
		}
		return strings.Join(strSlice, ",")
	default:
		return ""
	}
}
