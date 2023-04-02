package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hoangtk0100/go-healthy/db/sqlc"
)

type createMealRequest struct {
	Name        string      `json:"name" binding:"required,min=1"`
	Description string      `json:"description"`
	Calories    int32       `json:"calories" binding:"required"`
	Type        db.MealType `json:"type" binding:"required"`
}

type mealResponse struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Calories    int32       `json:"calories"`
	Type        db.MealType `json:"type"`
	CreatedAt   time.Time   `json:"created_at"`
}

func newMealResponse(meal db.Meal) mealResponse {
	return mealResponse{
		Name:        meal.Name,
		Description: meal.Description.String,
		Calories:    meal.Calories.Int32,
		Type:        meal.Type,
		CreatedAt:   meal.CreatedAt,
	}
}

func (server *Server) createMeal(ctx *gin.Context) {
	var req createMealRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := server.store.GetUser(ctx, currentUserName)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	arg := db.CreateMealParams{
		Username: user.Username,
		Name:     req.Name,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
		Calories: sql.NullInt32{
			Int32: req.Calories,
			Valid: true,
		},
		Type: req.Type,
	}

	meal, err := server.store.CreateMeal(ctx, arg)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	rsp := newMealResponse(meal)
	apiResponse(ctx, rsp)
}

type listMealsRequest struct {
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
	FromDate string `form:"from_date"`
	ToDate   string `form:"to_date"`
}

func (server *Server) listMeals(ctx *gin.Context) {
	var req listMealsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := server.store.GetUser(ctx, currentUserName)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	arg := db.ListMealsParams{
		Username: user.Username,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
		FromDate: parseDateTimeStrToNullTime(req.FromDate),
		ToDate:   parseEndOfDayNullTime(req.ToDate),
	}

	meals, err := server.store.ListMeals(ctx, arg)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	rsp := []mealResponse{}
	for _, meal := range meals {
		rsp = append(rsp, newMealResponse(meal))
	}

	apiResponse(ctx, rsp)
}
