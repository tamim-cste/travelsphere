package models

type Country struct {
    Name       string  `json:"name"`
    Capital    string  `json:"capital"`
    Population int64   `json:"population"`
    Region     string  `json:"region"`
    Subregion  string  `json:"subregion"`
    FlagURL    string  `json:"flag"`
    Currency   string  `json:"currency"`
    Languages  string  `json:"languages"`
    Slug       string  `json:"slug"`
    Lat        float64 `json:"lat"`
    Lon        float64 `json:"lon"` 
}