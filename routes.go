package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Podologie Wundsam - Ihre medizinische Fußpflege in Dachau",
			"desc":  "Fachpraxis Podologie Wundsam, Ihr Ansprechpartner für medizinische Fußpflege und Podologie in Dachau bei München, in allen Fällen von Schmerzen und Deformationen an Ihren Füßen.",
		})
	})
	r.GET("/services", func(c *gin.Context) {
		c.HTML(http.StatusOK, "service.tmpl", gin.H{
			"title": "Übersicht über unsere Leistungen | Podologie Wundsam",
			"desc":  "Erhalten Sie eine Übersicht über alle Leistungen die wir Ihnen als Podologische Einrichtung in Dachau anbieten können.",
		})
	})
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.tmpl", gin.H{
			"title": "Neuigkeiten, News und Nachrichten | Podologie Wundsam",
			"desc":  "Finden Sie aktuelle Infos und Neuigkeiten rund um das Thema Podologie und unsere Fachpraxis in Dachau.",
		})
	})
	r.GET("/career", func(c *gin.Context) {
		c.HTML(http.StatusOK, "career.tmpl", gin.H{
			"title": "Karierremöglichkeiten in unserer Praxis | Podologie Wundsam ",
			"desc":  "Unsere Praxis-Team wächst vortlaufend und ist immer auf der Suche nach Unterstützung. Finden Sie hier eine Übersicht der Karierremöglichkeiten, die wir aktuell anbieten.",
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
			"title": "Kontakt mit der Praxis aufnehmen | Podologie Wundsam",
			"desc":  "Tretem Sie mit uns in Kontakt. Sie erreichen uns telefonisch, per Email und können auch gerne persönlich in Dachau vorbei kommen.",
		})
	})
	r.GET("/imprint", func(c *gin.Context) {
		c.HTML(http.StatusOK, "imprint.tmpl", gin.H{
			"title": "Impressum: Fachpraxis Podologie Wundsam | +49 (0) 8131-9980933",
			"desc":  "Das Impressum unserer Fachpraxis in Dachau",
		})
	})
	r.GET("/privacy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy.tmpl", gin.H{
			"title": "Datenschutzerklärung | Podologie Wundsam",
			"desc":  "Hier führen wir auf, wie personenbezogene Daten auf unserer Website verarbeitet werden.",
		})
	})
	r.GET("/licences", func(c *gin.Context) {
		c.HTML(http.StatusOK, "licence.tmpl", gin.H{
			"title": "Lizenzen unserer Ressourcen | Podologie Wundsam",
			"desc":  "Sehen Sie sich alle Lizenzen der verwendeten Ressourcen ein.",
		})
	})
	r.GET("/sitemap.xml", func(c *gin.Context) {
		c.File("public/files/sitemap.xml")
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
			"title": "404 | Podologie Wundsam",
			"desc":  "Die von Ihnen gewünschte Seite konnte leider nicht gefunden werden.",
		})
	})
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("public/files/favicon.ico")
	})
}
