package GraphBlender

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func TestLoadDS(t *testing.T) {
	open, _ := os.Open("datasets/health_and_population.csv")
	reader := csv.NewReader(open)
	//reader.LazyQuotes = true
	strings, _ := reader.Read()
	for _, s := range strings {
		fmt.Println(s)
	}
}
