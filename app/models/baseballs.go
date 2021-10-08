package models
 
type Baseballs struct {
    ID int
    Name string `json:"name"`
    Manager string `json:"manager"`
    Home string `json:"home"`
}