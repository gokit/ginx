// Copyright 2020 The niqingyang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://acme.top
// Author: niqingyang	niqy@qq.com
// TIME: 2020/11/19 00:08
package boot

import (
	"github.com/gin-gonic/gin"
	"log"
)

// gin group map
var groupMap = make(map[string]*gin.RouterGroup)

// 存放中间件的回调，回调函数内部实现注册中间件的功能
var middlewaresCallback []func(app *gin.Engine)

// 存放注册的路由回调，回调函数内部实现注册路由的功能
var routersCallback []func(app *gin.Engine)

// 存放注册的路由群组回调，回调函数内部实现注册路由群组的功能
var groupsCallback []func(app *gin.Engine)

// 注册路由群组
func AddGroup(name string, group *gin.RouterGroup) {
	groupMap[name] = group
}

// 获取指定路由群组
func GetGroup(name string) *gin.RouterGroup {
	return groupMap[name]
}

// 向指定的 Echo 实例中注入中间件回调
func Middleware(callback func(app *gin.Engine)) {
	middlewaresCallback = append(middlewaresCallback, callback)
}

// 向指定的 Echo 实例中注入路由回调
func Route(callback func(app *gin.Engine)) {
	routersCallback = append(routersCallback, callback)
}

// 向指定的 Echo 实例中注入路由群组回调
func Group(callback func(app *gin.Engine)) {
	groupsCallback = append(groupsCallback, callback)
}

// 向指定的路由群组中中注入路由回调
func GroupByName(groupName string, callback func(group *gin.RouterGroup)) {
	routersCallback = append(routersCallback, func(app *gin.Engine) {
		if g, ok := groupMap[groupName]; ok {
			callback(g)
		} else {
			log.Printf("[GINX-BOOT] gin group %s not exists", groupName)
		}
	})
}

// 在注册 Gin 实例后，进行初始化
func Init(app *gin.Engine) {

	// 中间件
	for _, callback := range middlewaresCallback {
		callback(app)
	}

	// 路由群组
	for _, callback := range groupsCallback {
		callback(app)
	}

	// 路由
	for _, callback := range routersCallback {
		callback(app)
	}

}
