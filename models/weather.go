package models

type Weather struct {
    TempC     float64 
    Condition string   
    Icon      string  
    Humidity  int
    WindKph   float64
    City      string
}