package main

import (
	"log"
    "fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
    loadEnv()
    connectDb()
    defer db.Close()

    err := db.Ping()
    if err != nil {
        log.Fatal("err at ping: " + err.Error())
    } else {
        fmt.Println("db connected")
    }


    app := fiber.New()

    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.SendString("pong")
    })

    app.Get("/get-user/:id", getUser)

    log.Fatal(app.Listen(":8080"))
}

// add quiz
// delete quiz
// get users
func getUser(c *fiber.Ctx) error {
    rows, err := db.Query(fmt.Sprintf("SELECT * FROM users WHERE user_id=%v;", c.Params("id")))
    if err != nil {
        log.Fatal("err at query: " + err.Error())
    }

    for rows.Next() {
        var id int
        var username string
        var password string
        rows.Scan(&id, &username, &password)
        fmt.Println("id:", id)
        fmt.Println("username:", username)
        fmt.Println("password:", password)
    }

    return c.SendString("success")

}
