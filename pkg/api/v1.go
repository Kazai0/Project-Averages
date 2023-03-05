package api

type CreateAverage struct {
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

type CreateWeightedAverage struct {
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
	Peso1  int `json:"peso1"`
	Peso2  int `json:"peso2"`
}

type CreateMedian struct {
	Value1 []int `json:"value1"`
}
