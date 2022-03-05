package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Hackintosh struct {
    ID      string `json:"id"`
    Title   string `json:"title"`
    Year    string `json:"year"`
}

var ha = []Hackintosh{
    {ID: "10.15", Title: "macOS Catalina", Year: "2020"},
    {ID: "11.0", Title: "macOS Big Sur", Year: "2021"},
    {ID: "12.0", Title: "macOS Monterey", Year: "2022"},
}

func getHackintosh(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, ha)
}

func getHackintoshByID(c *gin.Context) {
    id := c.Param("id")
    for _, h := range ha {
        if h.ID == id {
            c.IndentedJSON(http.StatusOK, h)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "hackintosh not found"})
}

func postHackintosh(c *gin.Context) {
    var h Hackintosh
    if err := c.BindJSON(&h); err != nil {
        return
    }
    ha = append(ha, h)
    c.IndentedJSON(http.StatusCreated, h)
}

func deleteHackintosh(c *gin.Context) {
    ha = make([]Hackintosh, 0) 
    c.IndentedJSON(http.StatusOK, ha)
}

func deleteHackintoshByID(c *gin.Context) {
    id := c.Param("id")
    for _, h := range ha {
        if h.ID != id {
            c.IndentedJSON(http.StatusOK, h)
        }
    }
}

func main() {
    router := gin.Default()

    router.GET("/hackintosh", getHackintosh)
    router.GET("/hackintosh/:id", getHackintoshByID)
    router.POST("/hackintosh", postHackintosh)
    router.DELETE("/hackintosh", deleteHackintosh)
    router.DELETE("/hackintosh/:id", deleteHackintoshByID)

    router.Run("localhost:8080")
}
