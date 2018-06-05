package main

type RequestFormat struct {
    Length *int       `json:"length,omitempty"`
    IncludeCaps *bool `json:"caps,omitempty"`
}

type Return struct {
    Status bool
    ErrorCode int
    Value string
}

type RedisElement struct {
    Length int
    IncludeCaps bool
    Value string
    Time int32
}