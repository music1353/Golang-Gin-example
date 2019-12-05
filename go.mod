module my-gin

go 1.13

require (
    github.com/gin-gonic/gin v1.5.0
    github.com/lib/pq v1.2.0
)

replace (
    my-gin/database => /Users/jensonsu/Go-Application/my-gin/database
    my-gin/models => /Users/jensonsu/Go-Application/my-gin/models
    my-gin/router => /Users/jensonsu/Go-Application/my-gin/router
)
