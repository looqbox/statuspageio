package root

<<<<<<< HEAD
import (
        "encoding/json"
)

type BaseRequest struct {
        Url     string
        Headers []Header
}

type RequestJson struct {
        Url     string
        Method  string
        Headers []Header
        Body    *json.RawMessage
}

type RequestParam struct {
        Url     string
        Method  string
        Headers []Header
        body    string
}

type Header struct {
        Name  string
        Value string
}



type Component struct {
        ComponentId string `json:"component_id"`
=======
type BaseRequest struct {
	Url     string
	Headers []Header
}

type Header struct {
	Name  string
	Value string
>>>>>>> 7af2d93... Add root package
}
