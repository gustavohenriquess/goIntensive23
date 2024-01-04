package main

import (
	"net/http"

	"github.com/gustavohenriquess/go-intensive23/internal/entity"
	"github.com/labstack/echo"
)

func main() {
	//Chi
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", r)
	e := echo.New()
	e.GET("/order", OrderHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

func OrderHandler(c echo.Context) error {
	order, _ := entity.NewOrder("4", 10.0, 0.1)
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, order)
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
