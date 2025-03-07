package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	db "github.com/olartbaraq/spectrumshelf/db/sqlc"
	"github.com/olartbaraq/spectrumshelf/utils"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
	config  *utils.Config
	config2 *utils.Config
}

var tokenManager *utils.JWTToken

var cloudValues *CloudinaryValues

var Rdb *redis.Client

func NewServer(envPath string) *Server {

	config, err := utils.LoadDBConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Could not load env.env config: %v", err))
	}
	config2, err := utils.LoadOtherConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Could not load env.env config: %v", err))
	}

	conn, err := sql.Open(config2.DBdriverLive, config2.DBsourceLive)
	if err != nil {
		panic(fmt.Sprintf("There was an error connecting to database: %v", err))
	}

	tokenManager = utils.NewJWTToken(config2)

	cloudValues = NewCloudinaryValues(config2)

	q := db.New(conn)

	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	g.MaxMultipartMemory = 8 << 20

	g.Use(cors.Default())

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config2.RedisAddress,
		Password: config2.RedisPassword,
		DB:       0, // use default DB
	})
	return &Server{
		queries: q,
		router:  g,
		config:  config,
		config2: config2,
	}

}

func (s *Server) Start(port int) {

	if V, ok := binding.Validator.Engine().(*validator.Validate); ok {

		V.RegisterValidation("passwordStrength", ValidatePassword)
		V.RegisterValidation("isImageURL", ImageURLValidation)
		V.RegisterValidation("isPositive", PriceValidation)

	}

	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Home": "Welcome to Ra'Nkan Homepage...",
		})
	})

	User{}.router(s)
	Auth{}.router(s)
	Category{}.router(s)
	SubCategory{}.router(s)
	Shop{}.router(s)
	Product{}.router(s)
	Oauth{}.router(s)
	Order{}.router(s)

	s.router.Run(fmt.Sprintf(":%d", port))
}
