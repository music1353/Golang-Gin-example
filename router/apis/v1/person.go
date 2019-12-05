package v1

import (
    "net/http"
    "log"
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
    // . "my-gin/models" // py: from models import *
    models "my-gin/models" // py: import models
)

func IndexApi(c *gin.Context) {
    c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
    firstName := c.Request.FormValue("first_name")
    lastName := c.Request.FormValue("last_name")
    p := models.Person{FirstName: firstName, LastName: lastName}

    ra, err := p.AddPerson()
    if err != nil {
        log.Fatalln(err)
    }

    msg := fmt.Sprintf("insert successful %d", ra)
    c.JSON(http.StatusOK, gin.H {
        "data": true,
        "msg": msg,
    })
}

func ModPersonApi(c *gin.Context) {
    firstName := c.Request.FormValue("first_name")
    lastName := c.Request.FormValue("last_name")
    id, err := strconv.Atoi(c.Request.FormValue("id"))
    if err != nil {
        log.Fatalln(err)
    }

    p := Person{Id: id}
    p.GetPerson()
    
    if p.FirstName != '' {
        p.FirstName = firstName
        p.LastName = lastName
        num, err = p.ModPerson()
        if err != nil {
            log.Fatalln(err)
        }
        msg := fmt.Sprintf("update successful %d", num)
        c.JSON(http.StatusOK, gin.H {
            "data": true,
            "msg": msg,
        })
    } else {
        c.JSON(http.StatusOK, gin.H {
            "data": nil,
            "msg": "Person not found"
        })
    }
}

func DelPersonApi(c *ginContext) {

}

func GetPersonsApi(c *ginContext) {
    
}

func GetPersonApi(c *gin.Context) {
    id, err := strconv.Atoi(c.Query("id"))
    if err != nil {
        log.Fatalln(err)
    }
    person := models.Person{Id: id}

    person.GetPerson()
    if person.FirstName != "" {
        c.JSON(http.StatusOK, gin.H {
            "data": person,
            "msg": "success",
        })
    } else {
        c.JSON(http.StatusOK, gin.H {
            "data": nil,
            "msg": "Person not found",
        })
    }
}

