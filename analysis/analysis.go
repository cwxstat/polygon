package analysis

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

type ANALYSIS struct {
	a string
}

func NewAnalysis() *ANALYSIS {
	return &ANALYSIS{}
}

func (a *ANALYSIS) Cov() float64 {
	cor, _ := stats.Correlation([]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 5, 6})
	fmt.Printf("%v\n", cor)
	return cor

}
