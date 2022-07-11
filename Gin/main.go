package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type News struct {
	Title   string
	Content string
}
type User struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	// 1.创建一个默认的路由
	r := gin.Default()
	// 2.配置路由handler
	r.Static("./static", "static")  // 加载静态资源文件（映射本地资源）
	r.LoadHTMLGlob("template/**/*") // 加载html模板
	r.LoadHTMLFiles("./template/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{
			"title": "我是后台数据",
			"news":  News{Title: "这是标题", Content: "这是内容"},
		})
	})
	r.GET("/string", func(c *gin.Context) {
		c.String(200, "值：%v", "你好gin")
	})
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "<h1>JSON我的天那</h1>", //< 转义
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		c.AsciiJSON(200, gin.H{
			"message": "<h1>AsciiJSON我的天哪</h1>", //<、非ascii转义
		})
	})
	r.GET("/json3", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"message": "<h1>PureJSON我的天那</h1>", //完全不转义
		})
	})
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "JSONP",
		})
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"message": "XML",
		})
	})
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"score": 89,
			"news":  News{Title: "标题", Content: "内容"},
			"newsList": []interface{}{
				News{Title: "标题1", Content: "内容1"},
				News{Title: "标题2", Content: "内容2"},
				News{Title: "标题3", Content: "内容3"},
				News{Title: "标题4", Content: "内容4"},
				News{Title: "标题5", Content: "内容5"},
			},
			"time": time.Now(),
		})
	})
	r.GET("/demo/1", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")
		c.JSONP(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", nil)
	})
	r.POST("/login1", func(c *gin.Context) {
		// 获取表单传递的数据
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSONP(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	r.POST("/login2", func(c *gin.Context) {
		user := &User{}
		if err := c.ShouldBind(user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})
	// 解析xml
	r.POST("xml", func(c *gin.Context) {
		xmlSliceData, err := c.GetRawData()
		if err != nil {
			return
		}
		user := &User{}
		err = xml.Unmarshal(xmlSliceData, user)
		if err != nil {
			return
		}
		c.JSONP(http.StatusOK, user)
	})
	r.GET("/getUser", func(c *gin.Context) {
		user := &User{}
		if err := c.Bind(user); err == nil {
			c.JSONP(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})
	// 如果没有 /demo/user 的路由，则匹配  /demo/user、/demo/user/、/demo/user/action
	// 有的话  匹配  /demo/user/、/demo/user/action
	r.GET("/demo/:user/*action", func(c *gin.Context) {
		user := c.Param("user")
		action := c.Param("action")
		c.JSON(http.StatusOK, gin.H{
			"user":   user,
			"action": action,
		})
	})
	//r.GET("/demo/:user", func(c *gin.Context) {
	//	user := c.Param("user")
	//	c.JSON(http.StatusOK, gin.H{
	//		"user": user,
	//	})
	//})
	// 外部重定向
	r.GET("/redirect1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	// 内部路由重定向
	r.GET("/redirect2", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		log.Printf("Cookie value: %s \n", cookie)
	})
	r.POST("upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.SaveUploadedFile(file, "./")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	// 模拟一些私人数据
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}
	v1 := r.Group("/secrets").Use(gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	{
		v1.GET("/1", func(c *gin.Context) {
			// 获取用户，它是由 BasicAuth 中间件设置的
			user := c.MustGet(gin.AuthUserKey).(string)
			if secret, ok := secrets[user]; ok {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
			} else {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
			}
		})
	}
	// 3.启动服务器
	r.Run(":8080")
	//m := autocert.Manager{
	//	Prompt:     autocert.AcceptTOS,
	//	HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
	//	Cache:      autocert.DirCache("C:\\.cache"),
	//}
	//log.Fatal(autotls.RunWithManager(r, &m)) //https
}
