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

type ResponseStatus int64

const (
    Open ResponseStatus = iota
    Accepted
    Declined
)

func (e ResponseStatus) String() string {
    switch e {
    case Open:
        return "open"
    case Accepted:
        return "accepted"
    case Declined:
        return "declined"
    }
    // TODO: Might just want to throw error
    return ""
}

func ResponseStringToEnum(str string) ResponseStatus {
    switch str {
    case "open":
        return Open
    case "accepted":
        return Accepted
    case "declined":
        return Declined
    }
    // TODO: Might just want to throw error
    return Open
}

type Response struct {
    Id int
    Challenge_Id string
    User_Id string
    Status ResponseStatus
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
