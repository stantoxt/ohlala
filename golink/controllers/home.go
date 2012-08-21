package controllers

import (
    "github.com/QLeelulu/goku"
    "github.com/QLeelulu/ohlala/golink/models"
)

var _ = goku.Controller("home").
    Get("index", func(ctx *goku.HttpContext) goku.ActionResulter {
    u, ok := ctx.Data["user"]
    if !ok || u == nil {
        return ctx.Redirect("/discover")
    }
    user := u.(*models.User)
    ot := ctx.Get("o")
    if ot == "" {
        ot = "hot"
    }
    ctx.ViewData["Order"] = ot
    links, _ := models.Link_ForUser(user.Id, ot, 1, 20) //models.Link_GetByPage(1, 20)
    ctx.ViewData["Links"] = models.Link_ToVLink(links, ctx)
    return ctx.View(nil)
})
