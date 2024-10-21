package investment

import (
	"fmt"
	"log"
	"math"
)

// InvestmentResponse represents the results of the investment calculation
type InvestmentResponse struct {
	FutureValue        float64 `json:"futureValue"`
	TotalContributions float64 `json:"totalContributions"`
	TotalInterest      float64 `json:"totalInterest"`
}

// InvestmentRequest represents the input parameters for the investment calculation
type InvestmentRequest struct {
	InitialAmount       float64 `json:"initialAmount"`
	MonthlyContribution float64 `json:"monthlyContribution"`
	AnnualInterestRate  float64 `json:"annualInterestRate"`
	Years               int     `json:"years"`
}

// SimpleCalculator defines the interface for investment calculations
type SimpleCalculator interface {
	CalculateFutureValue(request InvestmentRequest) (InvestmentResponse, error)
}

// Calculator implements the SimpleCalculator interface
type Calculator struct {
	// Future: Add dependencies here, such as a logger
}

// CalculateFutureValue calculates the future value of an investment
func (c *Calculator) CalculateFutureValue(request InvestmentRequest) (InvestmentResponse, error) {
	// Input validation
	if request.InitialAmount < 0 || request.MonthlyContribution <= 0 || request.AnnualInterestRate < 0 || request.Years < 0 {
		err := fmt.Errorf("alle Eingabewerte müssen positiv sein")
		log.Println("Error in CalculateFutureValue:", err)
		return InvestmentResponse{}, err
	}

	// Convert annual interest rate to monthly and term to months
	monthlyRate := request.AnnualInterestRate / 100 / 12
	totalMonths := request.Years * 12

	// Calculate future value using the formula:
	// FV = P * (1 + r)^n + PMT * ((1 + r)^n - 1) / r
	futureValue := request.InitialAmount * math.Pow(1+monthlyRate, float64(totalMonths))
	if monthlyRate != 0 {
		futureValue += request.MonthlyContribution * (math.Pow(1+monthlyRate, float64(totalMonths)) - 1) / monthlyRate
	} else {
		futureValue += request.MonthlyContribution * float64(totalMonths)
	}

	totalContributions := request.InitialAmount + request.MonthlyContribution*float64(totalMonths)
	totalInterest := futureValue - totalContributions

	// Log the calculation
	log.Printf("CalculateFutureValue: capital=%.2f, deposit=%.2f, interest=%.2f, term=%d, futureValue=%.2f\n",
		request.InitialAmount, request.MonthlyContribution, request.AnnualInterestRate, request.Years, futureValue)

	// Create the response
	response := InvestmentResponse{
		FutureValue:        math.Round(futureValue*100) / 100,
		TotalContributions: math.Round(totalContributions*100) / 100,
		TotalInterest:      math.Round(totalInterest*100) / 100,
	}

	return response, nil
}
