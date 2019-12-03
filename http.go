package main

import (
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/deepch/vdk/format/mp4f"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func serveHTTP() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/", func(c *gin.Context) {
		fi, all := Config.list()
		sort.Strings(all)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"port":     Config.Server.HTTPPort,
			"suuid":    fi,
			"suuidMap": all,
			"version":  time.Now().String(),
		})
	})
	router.GET("/player/:suuid", func(c *gin.Context) {
		_, all := Config.list()
		sort.Strings(all)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"port":     Config.Server.HTTPPort,
			"suuid":    c.Param("suuid"),
			"suuidMap": all,
			"version":  time.Now().String(),
		})
	})
	router.GET("/ws/:suuid", func(c *gin.Context) {
		handler := websocket.Handler(ws)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	router.StaticFS("/static", http.Dir("web/static"))
	err := router.Run(Config.Server.HTTPPort)
	if err != nil {
		log.Fatalln(err)
	}
}
func ws(ws *websocket.Conn) {
	defer ws.Close()
	suuid := ws.Request().FormValue("suuid")
	log.Println("Request", suuid)
	if Config.ext(suuid) {
		ws.SetWriteDeadline(time.Now().Add(5 * time.Second))
		suuid := ws.Request().FormValue("suuid")
		cuuid, ch := Config.clAd(suuid)
		defer Config.clDe(suuid, cuuid)
		codecs := Config.coGe(suuid)
		if codecs == nil {
			log.Println("No Codec Info")
			return
		}
		muxer := mp4f.NewMuxer(nil)
		muxer.WriteHeader(codecs)
		meta, init := muxer.GetInit(codecs)
		err := websocket.Message.Send(ws, append([]byte{9}, meta...))
		if err != nil {
			return
		}
		err = websocket.Message.Send(ws, init)
		if err != nil {
			return
		}
		var start bool
		for {
			select {
			case pck := <-ch:
				if pck.IsKeyFrame {
					start = true
				}
				if !start {
					continue
				}
				ready, buf, _ := muxer.WritePacket(pck, false)
				if ready {
					ws.SetWriteDeadline(time.Now().Add(10 * time.Second))
					err := websocket.Message.Send(ws, buf)
					if err != nil {
						return
					}
				}
			}
		}
	}
}
