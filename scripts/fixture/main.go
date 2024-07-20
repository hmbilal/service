package main

import "fmt"

//import (
//	apiRoutes "activities/internal/api/routes"
//	"activities/internal/auth"
//	"activities/internal/container"
//	"activities/scripts/fixture/model"
//)
//
//func main() {
//	container.Register(
//		apiRoutes.NewRoutesConfig,
//	)
//
//	c := container.Bootstrap()
//	defer container.RecoverPanic(c)
//
//	if err := c.Invoke(run); err != nil {
//		panic(err)
//	}
//}
//
//func run(
//	authRepository auth.Repository,
//) {
//	model.CreateProject(authRepository)
//}

func main() {
	fmt.Printf("It workds!")
}
