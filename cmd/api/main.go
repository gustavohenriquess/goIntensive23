package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gustavohenriquess/go-intensive23/internal/Infra/database"
	usecase "github.com/gustavohenriquess/go-intensive23/internal/useCase"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

type OrderRequest struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

func main() {
	//Chi
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", r)
	e := echo.New()
	e.POST("/order", OrderHandler)

	e.Logger.Fatal(e.Start(":8888"))

}

func OrderHandler(c echo.Context) error {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close() // espera tudo rodar e depois executa o close

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Decodifica o JSON
	var bodyJson usecase.OrderInput
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	output, err := uc.Execute(bodyJson)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, output)
}

// func OrderHandler(w http.ResponseWriter, r *http.Request) {
// 	order, _ := entity.NewOrder("4", 10.0, 0.1)
// 	err := order.CalculateFinalPrice()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}

// 	result := json.NewEncoder(w).Encode(order)
// 	if result != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// }
