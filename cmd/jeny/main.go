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

// type User struct {
// 	ID 		primitive.ObjectID	`bson:"_id,omitempty"`
// 	User 	string				`bson:"user,omitempty"`
// 	Gen 	[]string			`bson:"gen,omitempty"`

// }

// type Task struct {
// 	ID          primitive.ObjectID 	`bson:"_id,omitempty"`
// 	OrderBy     string             	`bson:"order_by,omitempty"`
// 	Task        string             	`bson:"task,omitempty"`
// 	OrderTo     string             	`bson:"order_to,omitempty"`
// 	Deadline    string          	`bson:"deadline,omitempty"`
// 	CreatedAt   time.Time          	`bson:"created_at,omitempty"`
// 	DoneAt      string          	`bson:"done_at,omitempty"`
// }

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
	

	// user := db.Collection("user")
	// var result User
	// find := user.FindOne(ctx, bson.D{
	// 	{"gen",bson.D{{"$all",bson.A{"male","female"}}}},			// get All sub array
	// 	}).Decode(&result)
	// if find != nil {
	// 	logrus.WithError(find).Fatal("find not Found")
	// }
	// fmt.Println(result)
	// userResult, err := user.InsertOne(ctx, bson.D{
	// 	{"user","game"},
	// 	{"gen", bson.A{"male","male","female"}},
	// })
	// if err != nil {
	// 	logrus.WithError(err).Fatal("Insert Error")
	// }
	// fmt.Printf("Insert %v into Collections",userResult.InsertedID)

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
	
	e.Logger.Fatal(e.Start(fmt.Sprint(":", cfg.Port)))
}