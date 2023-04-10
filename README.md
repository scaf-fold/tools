### 引用

go get github.com/scaf-fold/tools@latest

~~~
func main() {
    e := gin.Default()
    courier.App(e, Root, func(e *gin.Engine) {
        e.Run(":8081")
    })
}
~~~

### 路由结构使用
~~~
controller
├── business
│   ├── get_resource.go
│   ├── get_meta.go
│   └── root.go
└── root.go
~~~

### 添加注册
#### root.go
~~~
func init(){
    V1:=root.Group(/v1)
    {
        V1.Register(&SimpleDemo)
    }
}

var rootName = "/service-name"
var Root = courier.NewGroup(rootName)
~~~
#### simple_demo.go
~~~
type SimpleDemo struct{
  courier.MethodGet
}

func (s *SimpleDemo)Path()string{
    return "/path"
}

func (s *SimpleDemo)Context(ctx *gin.Context){
    ctx.JSON(http.StatusNoContent,nil)
}
~~~
### gin 的使用
<url>https://github.com/gin-gonic/gin

