/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2023-05-24 17:35:34
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2023-05-24 17:35:47
 * @FileName: crowData
 * @Desc:
 *
 * Copyright (c) 2023 by small_ant, All Rights Reserved.
 */

package cron

import "github.com/robfig/cron"

func Run() {
    c := cron.New()
    // c.AddFunc("@every 3s", CheckWithDraw)
    // c.AddFunc("0 30 16 * * ?", Reward)
    c.Start()
}
