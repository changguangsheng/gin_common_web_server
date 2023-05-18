package router

import "ewa_admin_server/router/system"

//import "gin_common_web_server/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
