package models

type Test struct {
    Id int
    Name string
}

type Location struct {
    Latitude float32 `json:"lat"`
    Longitude float32 `json:"long"`
}

type Challenge struct {
    Id int
    Icon string
    Is_Active bool
    Location Location
    Picture_Url string
    Title string
}

type Response struct {
    Id int
    Challenge_Id string
    User_Id string
    Status string
    Picture_Url string
}

type User_Challenge struct {
    Challenge_Id int
    Challenger_Id int
    Challenged_Id int
}

type Create_Challenge struct {
    Challenger_Id int `json:"challenger_id" binding:"required"`
    Title string `json:"title" binding:"required"`
    Location Location `json:"location"`
    Picture_Url string `json:"picture_url"`
    Icon string `json:"icon"`
    Challenged_Ids []int `json:"challenged_ids"`
}
