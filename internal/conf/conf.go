package conf

import (
	"errors"
	"time"

	"github.com/Kambar-ZH/simple-service/pkg/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	GormDB *gorm.DB

	Database Database
	JWT      JWT
	App      App
	Lgr      logger.Logger
}

func (gc *Config) Init() (err error) {
	gc.InitLogger()
	if err = gc.InitVars(); err != nil {
		panic(err)
	}
	if err = gc.InitGormDB(); err != nil {
		panic(err)
	}

	return
}

func (gc *Config) InitVars() (err error) {
	viper.SetConfigFile("app.env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// DATABASE CONFIGURATION
	{
		gc.Database.User = viper.GetString("POSTGRES_USER")
		gc.Database.Name = viper.GetString("POSTGRES_DB")
		gc.Database.Host = viper.GetString("POSTGRES_HOST")
		gc.Database.HostName = viper.GetString("POSTGRES_HOSTNAME")
		gc.Database.Port = viper.GetInt("DATABASE_PORT")
		gc.Database.Password = viper.GetString("POSTGRES_PASSWORD")
	}

	// JWT CONFIGURATION
	{
		gc.JWT.SecretKey = viper.GetString("JWT_SECRET_KEY")
	}

	// APP CONFIGURATION
	{
		gc.App.Name = viper.GetString("APP_NAME")
	}

	return
}

func (gc *Config) InitGormDB() (err error) {
	conn, err := gorm.Open(postgres.Open(gc.Database.DSN()), &gorm.Config{})
	if err != nil {
		return
	}

	start := time.Now()
	db, err := conn.DB()
	if err != nil {
		return
	}
	for db.Ping() != nil {
		if start.After(start.Add(5 * time.Second)) {
			err = errors.New("failed connect to db")
			break
		}
	}

	gc.GormDB = conn
	return
}

func (gc *Config) InitLogger() {
	tops := []logger.TeeOption{
		{
			Filename: "logs/access.log",
			Ropt: logger.RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3,
				Compress:   true,
			},
			Lef: func(lvl logger.Level) bool {
				return lvl <= logger.InfoLevel
			},
		},
		{
			Filename: "logs/error.log",
			Ropt: logger.RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3,
				Compress:   true,
			},
			Lef: func(lvl logger.Level) bool {
				return lvl > logger.InfoLevel
			},
		},
	}

	lgr := logger.NewTeeWithRotate(tops)
	gc.Lgr = lgr
}
