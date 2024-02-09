package api_controller

import (
	dbconn "cms-fello/app/databases/connection"
	"cms-fello/app/databases/repository"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func LoginUser(ctx echo.Context) (err error) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	serverCms, err := db
	if err != nil {
		return err
	}

	dbCms := repository.AccessRepositoryCms(serverCms)
	result, err := dbCms.ViewCmsUserTableByUsername(context.Background(), username, ctx.RealIP())
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errors.New("Akun anda tidak terdaftar")
		} else {
			return err
		}
	}

	hasher := sha256.New()
	hasher.Write([]byte(password))
	sha := hex.EncodeToString(hasher.Sum(nil))

	if result.Password_user != sha {
		return errors.New("password tidak sesuai")
	}

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 1, // Days
	}

	sess.Values["UserID"] = result.Id_user
	sess.Values["UserName"] = result.Username
	if result.Level_user == "1" || result.Level_user == "2" {
		sess.Values["UserLevel"] = "markom"
	}

	sess.Save(ctx.Request(), ctx.Response())

	return ctx.JSON(http.StatusOK, &Response{
		Rc:   "00",
		Val:  "",
		Desc: "Sukses",
	})
}

func LogoutUser(ctx echo.Context) (err error) {
	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 1, // Days
	}

	sess.Values["UserID"] = ""
	sess.Values["UserName"] = ""
	sess.Values["UserLevel"] = ""

	sess.Save(ctx.Request(), ctx.Response())

	return ctx.NoContent(http.StatusOK)
}

func CheckSession(ctx echo.Context) (err error) {
	sess, _ := session.Get("session", ctx)
	sess.Save(ctx.Request(), ctx.Response())

	fmt.Printf("sess: %v\n", sess)

	return ctx.NoContent(http.StatusOK)
}
