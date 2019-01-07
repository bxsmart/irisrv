package public

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "cookie-session-name-id"
	msessions              = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func Secret(ctx iris.Context) {
	//验证用户授权
	if auth, _ := msessions.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	//输出消息
	ctx.WriteString("The cake is a lie!")
}

func Login(ctx iris.Context) {
	session := msessions.Start(ctx)
	// 在里执行验证
	// ...
	//把验证状态保存为true
	session.Set("authenticated", true)
}

func Logout(ctx iris.Context) {
	session := msessions.Start(ctx)

	// 撤消用户身份验证
	session.Set("authenticated", false)
}

func SetSession(ctx iris.Context, k string, o interface{}) {
	session := msessions.Start(ctx)
	session.Set(k, o)
}

func GetSession(ctx iris.Context, k string) interface{} {
	session := msessions.Start(ctx)
	return session.Get(k)
}

func SaveUser(ctx iris.Context, user interface{}) {
	SetSession(ctx, "user", user)
}

func LoadUser(ctx iris.Context) interface{} {
	return GetSession(ctx, "user")
}

func SaveRoleId(ctx iris.Context, roleId interface{}) {
	SetSession(ctx, "roleid", roleId)
}

func LoadRoleId(ctx iris.Context) interface{} {
	return GetSession(ctx, "roleid")
}

func ClearAllSession(ctx iris.Context) {
	session := msessions.Start(ctx)
	session.Clear()
}
