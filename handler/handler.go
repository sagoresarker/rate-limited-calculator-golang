package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sagoresarker/rate-limited-calculator-golang/calculator"
)

type CalculationHandler struct {
	calculator *calculator.RateLimitedCalculator
}

func NewCalculationHandler(calculator *calculator.RateLimitedCalculator) *CalculationHandler {
	return &CalculationHandler{calculator: calculator}
}

type CalculationRequest struct {
	Username       string `json:"username" validate:"required"`
	CalculatorType string `json:"calculator_type" validate:"required,oneof=add subtract multiply divide modulo power factorial"`
	Number1        int    `json:"number1"`
	Number2        int    `json:"number2"`
}

type CalculationResponse struct {
	Result int    `json:"result"`
	Error  string `json:"error,omitempty"`
}

func (h *CalculationHandler) HandleCalculation(c echo.Context) error {
	req := new(CalculationRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, CalculationResponse{Error: err.Error()})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, CalculationResponse{Error: "validator not registered"})
	}

	var result int
	var err error

	switch req.CalculatorType {
	case "add":
		result, err = h.calculator.Add(req.Number1, req.Number2)
	case "subtract":
		result, err = h.calculator.Subtract(req.Number1, req.Number2)
	case "multiply":
		result, err = h.calculator.Multiply(req.Number1, req.Number2)
	case "divide":
		result, err = h.calculator.Divide(req.Number1, req.Number2)
	case "modulo":
		result, err = h.calculator.Modulo(req.Number1, req.Number2)
	case "power":
		result, err = h.calculator.Power(req.Number1, req.Number2)
	case "factorial":
		result, err = h.calculator.Factorial(req.Number1)
	default:
		return c.JSON(http.StatusBadRequest, CalculationResponse{Error: "unknown calculator type"})
	}

	if err != nil {
		return c.JSON(http.StatusTooManyRequests, CalculationResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, CalculationResponse{Result: result})
}
