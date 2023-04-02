package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/hoangtk0100/go-healthy/db/mock"
	db "github.com/hoangtk0100/go-healthy/db/sqlc"
	"github.com/hoangtk0100/go-healthy/util"
	"github.com/stretchr/testify/require"
)

func TestCreateMealAPI(t *testing.T) {
	username := currentUserName
	meal := randomMeal(username)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":        meal.Name,
				"description": meal.Description.String,
				"calories":    meal.Calories.Int32,
				"type":        meal.Type,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateMealParams{
					Name:        meal.Name,
					Description: meal.Description,
					Calories:    meal.Calories,
					Type:        meal.Type,
				}
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(meal.Username)).Times(1)
				store.EXPECT().CreateMeal(gomock.Any(), gomock.Eq(arg)).Times(1).Return(meal, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchMeal(t, recorder.Body, meal)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name":        meal.Name,
				"description": meal.Description.String,
				"calories":    meal.Calories.Int32,
				"type":        meal.Type,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(meal.Username)).Times(1)
				store.EXPECT().CreateMeal(gomock.Any(), gomock.Any()).Times(1).Return(db.Meal{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidName",
			body: gin.H{
				"name":        "",
				"description": meal.Description.String,
				"calories":    meal.Calories.Int32,
				"type":        meal.Type,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(0)
				store.EXPECT().CreateMeal(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidCalories",
			body: gin.H{
				"name":        meal.Name,
				"description": meal.Description.String,
				"calories":    -10,
				"type":        meal.Type,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(0)
				store.EXPECT().CreateMeal(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/meals"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestListMealsAPI(t *testing.T) {
	n := 5
	meals := make([]db.Meal, n)
	username := currentUserName
	for index := 0; index < n; index++ {
		meals[index] = randomMeal(username)
	}

	type Query struct {
		pageID   int
		pageSize int
		fromDate string
		toDate   string
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListMealsParams{
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(username)).Times(1)
				store.EXPECT().ListMeals(gomock.Any(), gomock.Eq(arg)).Times(1).Return(meals, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchMeals(t, recorder.Body, meals)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(username)).Times(1)
				store.EXPECT().
					ListMeals(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Meal{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListMeals(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: -1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListMeals(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/meals"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			q.Add("from_date", tc.query.fromDate)
			q.Add("to_date", tc.query.toDate)
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomMeal(username string) db.Meal {
	return db.Meal{
		ID:       util.RandomUID(),
		Username: username,
		Name:     util.RandomString(10),
		Description: sql.NullString{
			String: util.RandomString(20),
			Valid:  true,
		},
		Calories: sql.NullInt32{
			Int32: util.RandomAmount(),
			Valid: true,
		},
		Type: db.MealTypeMORNING,
	}
}

func parseResponse(t *testing.T, body *bytes.Buffer) map[string]interface{} {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var rsp BaseResponse
	err = json.Unmarshal(data, &rsp)
	require.NoError(t, err)
	return rsp.Data.((map[string]interface{}))
}

func parseMealFromMap(data map[string]interface{}) db.Meal {
	return db.Meal{
		Name: data["name"].(string),
		Type: db.MealType(data["type"].(string)),
		Description: sql.NullString{
			String: data["description"].(string),
			Valid:  true,
		},
		Calories: sql.NullInt32{
			Int32: (int32)(data["calories"].(float64)),
			Valid: true,
		},
		CreatedAt: time.Time{},
	}
}

func requireBodyMatchMeal(t *testing.T, body *bytes.Buffer, meal db.Meal) {
	dt := parseResponse(t, body)

	gotMeal := parseMealFromMap(dt)

	require.Equal(t, meal.Name, gotMeal.Name)
	require.Equal(t, meal.Description, gotMeal.Description)
	require.Equal(t, meal.Calories, gotMeal.Calories)
	require.Equal(t, meal.Type, gotMeal.Type)
}

func requireBodyMatchMeals(t *testing.T, body *bytes.Buffer, meals []db.Meal) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var rsp BaseResponse
	err = json.Unmarshal(data, &rsp)
	require.NoError(t, err)

	var gotMeals []mealResponse
	dt := rsp.Data.([]interface{})
	for _, val := range dt {
		valIndex, err := json.Marshal(val)
		require.NoError(t, err)
		var mealIndex mealResponse
		err = json.Unmarshal(valIndex, &mealIndex)
		require.NoError(t, err)
		gotMeals = append(gotMeals, mealIndex)
	}

	require.Equal(t, len(meals), len(gotMeals))
}
