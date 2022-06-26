package main

import (
    "context"
    "fmt"

    "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/auth"
)

func main() {
    ctx := context.Background()

    app, err := firebase.NewApp(ctx, nil)
    if err != nil {
        panic(err)
    }

    authCli, err := app.Auth(ctx)
    if err != nil {
        panic(err)
    }

    user, err := authCli.CreateUser(ctx, &auth.UserToCreate{})
    if err != nil {
        panic(err)
    }

    fmt.Println(user)
}
