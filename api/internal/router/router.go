/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2023-05-24 16:23:12
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2023-05-24 16:27:45
 * @FileName: router
 * @Desc:
 *
 * Copyright (c) 2023 by small_ant, All Rights Reserved.
 */

package router

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
    v1 := r.Group("v1")
    {
        v1.GET("hello",
            func(context *gin.Context) {
                context.JSON(200, gin.H{"msg": "v1-登录成功"})
            },
        )
    }
}
