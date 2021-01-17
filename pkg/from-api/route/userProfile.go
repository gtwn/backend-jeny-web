package route

import (

	"github.com/jenywebapp/pkg/from-api/model"
	"github.com/jenywebapp/pkg/from-api/svc"
	"github.com/labstack/echo/v4"
)

type UserProfileConfig struct {
	AccessToken	string
}

func UserProfile(cfg UserProfileConfig) echo.HandlerFunc {

	return func(c echo.Context) error {
		groupId := c.Param("id")
		userIds,err := svc.GetUserIDInGroup(groupId,cfg.AccessToken)
		if err != nil {
			return err
		}
		// spew.Dump(*userIds)
		// fmt.Println(*userIds)
		profiles, err := svc.GetUsersProfile(userIds,cfg.AccessToken)
		if err != nil {
			return c.NoContent(400)
		}
		// spew.Dump(*profiles)
		// fmt.Println(*profiles)

		return c.JSON(200, model.UserProfileInGroup{
			UsersProfile: *profiles,
		})
	}
}