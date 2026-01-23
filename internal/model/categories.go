package model

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Categories = []Category{
	{ID: 1, Name: "Foods", Description: "All kinds of food items"},
	{ID: 2, Name: "Cigarettes", Description: "All kinds of cigarettes"},
}
