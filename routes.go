package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	r.GET("/services", func(c *gin.Context) {
		c.HTML(http.StatusOK, "service.tmpl", gin.H{
			"title": "Leistungen | ",
		})
	})
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.tmpl", gin.H{
			"title": "Neuigkeiten | ",
		})
	})
	r.GET("/career", func(c *gin.Context) {
		c.HTML(http.StatusOK, "career.tmpl", gin.H{
			"title": "Karierre | ",
		})
	})
	r.GET("/career/podologe", func(c *gin.Context) {
		c.File("public/files/Stellenanzeige_Podologie_Wundsam.pdf")
	})
	r.GET("/career/buero", func(c *gin.Context) {
		c.File("public/files/Stellenanzeige_Podologie_Wundsam_Buero.pdf")
	})
	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.tmpl", gin.H{
			"title": "Kontakt | ",
		})
	})
	r.GET("/imprint", func(c *gin.Context) {
		c.HTML(http.StatusOK, "imprint.tmpl", gin.H{
			"title": "Impressum | ",
		})
	})
	r.GET("/privacy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy.tmpl", gin.H{
			"title": "Datenschutz | ",
		})
	})
	r.GET("/licences", func(c *gin.Context) {
		c.HTML(http.StatusOK, "licence.tmpl", gin.H{
			"title": "Lizenzen | ",
		})
	})
}
