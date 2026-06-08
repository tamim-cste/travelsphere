package models

type Country struct {
    Name       string `json:"name"`
    Capital    string `json:"capital"`
    Population int64  `json:"population"`
    Region     string `json:"region"`
    FlagURL    string `json:"flag"`
    Currency   string `json:"currency"`
    Languages  string `json:"languages"`
    Slug       string `json:"slug"` 
}