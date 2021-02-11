package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jenywebapp/cmd/jeny/Config"
	route "github.com/jenywebapp/pkg/from-api/route"
	routeTask "github.com/jenywebapp/pkg/task/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func main() {

	cfg,err := config.Read()
	if err != nil {
		logrus.WithError(err).Fatal("Read config")
	}
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.aarl2.mongodb.net/%s?retryWrites=true&w=majority",cfg.DBUserName,cfg.DBPassword,cfg.DBName)
	fmt.Println(uri)
	client,err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logrus.WithError(err).Fatal("Database Failed !")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("Connect Database Failed ! ")
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.WithError(err).Fatal("Ping Error")
	}

	db := client.Database(cfg.DBName)
	

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Middleware debug
	e.Use(middleware.Logger())
	e.GET("/", routeTask.Hello)
	e.GET("/hello", routeTask.Hello)
	e.GET("/dashboard", route.LineToken(route.LineTokenConfig{
		LineAPI 		: cfg.LineTokenAPI,
		ChannelID 		: cfg.ChannelID,
		ChannelSecret	: cfg.ChannelSecret,
		},db))

	e.GET("/logout",route.Revoke())
	e.GET("/group/:id/profile", route.UserProfile(route.UserProfileConfig{
		AccessToken: cfg.ChannelAccessToken,
	}))
	taskGrp := e.Group("/task")

	taskGrp.GET("",routeTask.GetTask(db))
	taskGrp.GET("/follow",routeTask.GetFollow(db))
	taskGrp.GET("/review",routeTask.GetReview(db))
	taskGrp.GET("/history",routeTask.GetHistory(db))

	requestGrp := e.Group("/request")

	requestGrp.POST("/review/:id",routeTask.SendReviewTask(routeTask.ReviewTaskConfig{
		AccessToken: cfg.ChannelAccessToken,
	},db))
	requestGrp.POST("/follow/:id",routeTask.SendFollowTask(routeTask.FollowTaskConfig{
		AccessToken: cfg.ChannelAccessToken,
	},db))
	requestGrp.POST("/accept/:id",routeTask.SendDoneTask(routeTask.DoneTaskConfig{
		AccessToken: cfg.ChannelAccessToken,
	},db))
	requestGrp.POST("/reject/:id",routeTask.SendRejectTask(routeTask.RejectTaskConfig{
		AccessToken: cfg.ChannelAccessToken,
	},db))
	requestGrp.POST("/cancel/:id", routeTask.SendCancelTask(routeTask.CancelTaskConfig{
		AccessToken: cfg.ChannelAccessToken,
	}, db))
	
	e.Logger.Fatal(e.Start(fmt.Sprint(":", cfg.Port)))
	
}