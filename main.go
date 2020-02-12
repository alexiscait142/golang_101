package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


// Robot is the model for our db robot
type Robot struct {
	ID string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Type string `json:"type" db:"type"`
	Dangerous bool `json:"dangerous" db:"dangerous"`
}

func main(){
	r:= gin.Default()
	r.GET("/robots", handleRobots)
	r.Run()
}

func handleRobots(c *gin.Context) {
		db, err := sqlx.Connect("postgres", "user=alexischilinski dbname=postgres sslmode=disable")
		if err != nil {
			c.JSON(500, err.Error())
		}
		robots := []Robot{}
		// db.Select(&robots, "SELECT * FROM robot")
		err = db.Select(&robots, "SELECT * FROM robot")
		if err != nil {
			c.JSON(500, err.Error())
			// log.Fatalln(err)
		}
		c.JSON(200, robots)
	})
	r.Run()
}