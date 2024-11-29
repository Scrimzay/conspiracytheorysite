package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("htmlfiles/*/*.html")
	r.Static("/static", "./static")

	r.GET("/", indexHandler)

	// Level 1 conspiracy theory groups
	r.GET("/level1conspiracies", level1ConspiraciesHandler)
	r.GET("/fakemoonlanding", func(ctx *gin.Context) {
		ctx.HTML(200, "fakemoonlanding.html", nil)
	})
	r.GET("/aliensarereal", func(c *gin.Context) {
		c.HTML(200, "aliensarereal.html", nil)
	})
	r.GET("/lizardpeople", func(c *gin.Context) {
		c.HTML(200, "lizardpeople.html", nil)
	})
	r.GET("/mkultra", func(c *gin.Context) {
		c.HTML(200, "mkultra.html", nil)
	})
	r.GET("/elitespraisesatan", func(c *gin.Context) {
		c.HTML(200, "elitespraisesatan.html", nil)
	})
	r.GET("/newworldorder", func(c *gin.Context) {
		c.HTML(200, "newworldorder.html", nil)
	})
	r.GET("/israeldestabilizingmiddleeast", func(c *gin.Context) {
		c.HTML(200, "israeldestabilizingmiddleeast.html", nil)
	})
	r.GET("/jewscontroleverything", func(c *gin.Context) {
		c.HTML(200, "jewscontroleverything.html", nil)
	})
	r.GET("/flatearth", func(c *gin.Context) {
		c.HTML(200, "flatearth.html", nil)
	})
	r.GET("/nsaprivacybreach", func(c *gin.Context) {
		c.HTML(200, "ciaprivacybreach.html", nil)
	})
	r.GET("/socialmediaspying", func(c *gin.Context) {
		c.HTML(200, "socialmediaspying.html", nil)
	})


	// Level 2 conspiracies
	r.GET("/level2conspiracies", level2ConspiraciesHandler)
	r.GET("/governmentcontrolscomputers", func(c *gin.Context) {
		c.HTML(200, "governmentcontrolscomputers.html", nil)
	})
	r.GET("/smartdevicesrecordyou", func(ctx *gin.Context) {
		ctx.HTML(200, "smartdevicesrecordyou.html", nil)
	})
	r.GET("/blackrockdestroyinghousing", func(ctx *gin.Context) {
		ctx.HTML(200, "blackrockdestroyshousing.html", nil)
	})
	r.GET("/climatechangeonlyforwest", func(ctx *gin.Context) {
		ctx.HTML(200, "climatechangeonlyforwest.html", nil)
	})
	r.GET("/socialmediacensorship", func(ctx *gin.Context) {
		ctx.HTML(200, "socialmediacensorship.html", nil)
	})
	r.GET("/holocaustdenial", func(ctx *gin.Context) {
		ctx.HTML(200, "holocaustdenial.html", nil)
	})
	r.GET("/chemtrails", func(ctx *gin.Context) {
		ctx.HTML(200, "chemtrails.html", nil)
	})


	// Level 3 theories
	r.GET("/level3conspiracies", level3ConspiraciesHandler)
	r.GET("/gangstalking", func(ctx *gin.Context) {
		ctx.HTML(200, "gangstalking.html", nil)
	})
	r.GET("/governmentorbitallaser", func(ctx *gin.Context) {
		ctx.HTML(200, "governmentorbitallaser.html", nil)
	})
	r.GET("/fbirunbypedos", func(ctx *gin.Context) {
		ctx.HTML(200, "fbirunbypedos.html", nil)
	})
	r.GET("/fbiplansshootings", func(ctx *gin.Context) {
		ctx.HTML(200, "fbiplansshootings.html", nil)
	})
	r.GET("/internetbotpropaganda", func(ctx *gin.Context) {
		ctx.HTML(200, "internetbotpropaganda.html", nil)
	})
	r.GET("/deathandrebirthfromsleep", func(ctx *gin.Context) {
		ctx.HTML(200, "deathandrebirthfromsleep.html", nil)
	})
	r.GET("/undestroysnations", func(ctx *gin.Context) {
		ctx.HTML(200, "undestroysnations.html", nil)
	})
	r.GET("/epsteinstillalive", func(ctx *gin.Context) {
		ctx.HTML(200, "epsteinstillalive.html", nil)
	})
	r.GET("/holocaustnumbersfake", func(ctx *gin.Context) {
		ctx.HTML(200, "holocaustnumbersfake.html", nil)
	})
	r.GET("/replacementtheory", func(ctx *gin.Context) {
		ctx.HTML(200, "replacementtheory.html", nil)
	})



	// Level 4 theories
	r.GET("/level4conspiracies", level4ConspiraciesHandler)
	r.GET("/demonsinsidechips", func(ctx *gin.Context) {
		ctx.HTML(200, "demonsinsidechips.html", nil)
	})
	r.GET("/conciousgateway", func(ctx *gin.Context) {
		ctx.HTML(200, "conciousgateway.html", nil)
	})

	
	
	err := r.Run(":3372")
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func level1ConspiraciesHandler(c *gin.Context) {
	c.HTML(200, "level1.html", nil)
}

func level2ConspiraciesHandler(c *gin.Context) {
	c.HTML(200, "level2.html", nil)
}

func level3ConspiraciesHandler(c *gin.Context) {
	c.HTML(200, "level3.html", nil)
}

func level4ConspiraciesHandler(c *gin.Context) {
	c.HTML(200, "level4.html", nil)
}