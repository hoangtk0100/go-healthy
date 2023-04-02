package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/hoangtk0100/go-healthy/util"
	"github.com/stretchr/testify/require"
)

func createRandomMeal(t *testing.T, user User) Meal {
	arg := CreateMealParams{
		Username: user.Username,
		Name:     util.RandomString(8),
		Description: sql.NullString{
			String: util.RandomString(20),
			Valid:  true,
		},
		Calories: sql.NullInt32{
			Int32: util.RandomAmount(),
			Valid: true,
		},
		Type: MealTypeMORNING,
	}

	meal, err := testQueries.CreateMeal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, meal)

	require.Equal(t, meal.Username, arg.Username)
	require.Equal(t, meal.Name, arg.Name)
	require.Equal(t, meal.Description, arg.Description)
	require.Equal(t, meal.Calories, arg.Calories)
	require.Equal(t, meal.Type, arg.Type)
	require.NotZero(t, meal.CreatedAt)

	return meal
}

func TestCreateMeal(t *testing.T) {
	user := createRandomUser(t)
	createRandomMeal(t, user)
}

func TestGetMeal(t *testing.T) {
	user := createRandomUser(t)
	meal1 := createRandomMeal(t, user)
	meal2, err := testQueries.GetMeal(context.Background(), meal1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, meal2)

	require.Equal(t, meal1.Username, meal2.Username)
	require.Equal(t, meal1.Name, meal2.Name)
	require.Equal(t, meal1.Description, meal2.Description)
	require.Equal(t, meal1.Calories, meal2.Calories)
	require.Equal(t, meal1.Type, meal2.Type)
	require.WithinDuration(t, meal1.CreatedAt, meal2.CreatedAt, time.Second)
}

func TestListMeals(t *testing.T) {
	var lastMeal Meal
	user := createRandomUser(t)
	for index := 0; index < 4; index++ {
		lastMeal = createRandomMeal(t, user)
	}

	arg := ListMealsParams{
		Username: lastMeal.Username,
		Limit:    5,
		Offset:   0,
	}

	meals, err := testQueries.ListMeals(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, meals)
	require.Equal(t, 4, len(meals))

	for _, mealIndex := range meals {
		require.NotEmpty(t, mealIndex)
		require.Equal(t, lastMeal.Username, mealIndex.Username)
		require.Equal(t, lastMeal.Type, mealIndex.Type)
	}
}
