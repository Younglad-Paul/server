package model

type Plan struct {
    ID           string         `bson:"id"`
    Title        string         `bson:"title"`
    Amount       string         `bson:"amount"`
    Return       string         `bson:"return"`
    Duration     string         `bson:"duration"`
    ReferalBonus string         `bson:"referalBonus"`
    Description  []Description  `bson:"description"` 
}

type Description struct {
    ID    string `bson:"id"`
    Point string `bson:"point"`
}
