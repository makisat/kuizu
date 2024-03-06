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

    app.Post("/add-user", addUser)

    log.Fatal(app.Listen(":8080"))
}

// add quiz
// delete quiz
// get users
func getUser(c *fiber.Ctx) error {
    rows, err := db.Query(fmt.Sprintf("SELECT * FROM users WHERE user_id=%v", c.Params("id")))
    if err != nil {
        log.Fatal("err at query: " + err.Error())
    }

    user := new(User)

    for rows.Next() {
        rows.Scan(&user.UserId, &user.Username, &user.Password)
    }

    return c.JSON(user)
}

func addUser(c *fiber.Ctx) error {

    new_user := &struct {
        Username string `json:"username"`
        Password string `json:"password"`
    } {}

    err := c.BodyParser(new_user)
    if err != nil {
        log.Fatal("err at bodyparser: " + err.Error())
    }

    fmt.Println("new_user: ", new_user.Username)

    _, err = db.Exec(fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%s', '%s')", new_user.Username, new_user.Password))
    if err != nil {
        log.Fatal("err at inserting data: " + err.Error())
        return c.SendStatus(fiber.StatusBadRequest)
    }

    return c.SendStatus(fiber.StatusOK)
}
