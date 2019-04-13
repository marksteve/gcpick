package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
)

func getProjects() []*cloudresourcemanager.Project {
	ctx := context.Background()
	client, _ := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformReadOnlyScope)
	svc, _ := cloudresourcemanager.New(client)
	resp, _ := svc.Projects.List().Do()
	return resp.Projects
}

func listProjects(projects []*cloudresourcemanager.Project) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, projects)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	projects := getProjects()
	e.GET("/projects", listProjects(projects))
	e.Logger.Fatal(e.Start(":1323"))
}
