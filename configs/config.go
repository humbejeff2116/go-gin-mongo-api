
package configs

import (
	"fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
)

var ServerConfigurations config = SetServerConfigurations();

type app struct {
	Port string
}
type mONGO_Db_Options struct {
	useNewUrlParser, useUnifiedTopology, useCreateIndex, useFindAndModify bool
}
type database struct {
	mONGO_Db_Options
	DatabaseURI string
}
type secret struct {
	CookieSecret, SessionSecret, JwtSecret, PssrptJwtSecret  string
}
type gmail struct {
	User, Password string
}
type cloudinaryConfig struct {
	CloudName, ApiKey, Secret, Url string
}

type config struct {
	app
	database
	secret
	gmail
	cloudinaryConfig
}

func SetServerConfigurations() config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	server := app{
		Port:  os.Getenv("SERVER_PORT") ,
	}

	db := database{
		mONGO_Db_Options: mONGO_Db_Options{
			useNewUrlParser: true,
			useUnifiedTopology: true,
			useCreateIndex: true,
			useFindAndModify: false,
		},
		DatabaseURI: os.Getenv("MONGO_URI"),
	}

	secret := secret{
		CookieSecret:  os.Getenv("COOKIE_SECRET"),
		SessionSecret:  os.Getenv("SESSION_SECRET"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
		PssrptJwtSecret:  os.Getenv("PSSRPT_JWT_SECRET"),
	}

	gmail := gmail{
		User: os.Getenv("GMAIL_USERNAME"),
		Password: os.Getenv("GMAIL_PASSWORD"),
	}

	cloudinary := cloudinaryConfig{
		CloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		ApiKey: os.Getenv("CLOUDINARY_KEY"),
		Secret: os.Getenv("CLOUDINARY_SECRET"),
		Url: fmt.Sprintf(`cloudinary://%v:%v@%v`, os.Getenv("CLOUDINARY_KEY"), os.Getenv("CLOUDINARY_SECRET"), os.Getenv("CLOUDINARY_CLOUD_NAME")),
	}
	
	serverConfigurations := config{
		app: server,
		database: db,
		secret: secret, 
		gmail: gmail,
		cloudinaryConfig: cloudinary,

	}
    return serverConfigurations;
}