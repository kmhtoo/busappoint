package app

import (
	"bytes"
	"fmt"
	"github.com/bamzi/jobrunner"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	section := os.Getenv("ENV_MODE")
	configPath := "./conf/config.json"
	if section != "prod" {
		configPath = "./conf/local/config.json"
	}

	viper.SetConfigType("json")
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = viper.ReadConfig(bytes.NewBuffer(b))
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	jobrunner.Start()
	jobrunner.Schedule("@every 5s", ReminderEmails{})

	// jobrunner.Schedule("* */5 * * * *", DoSomething{}) // every 5min do something
	//    jobrunner.Schedule("@every 1h30m10s", ReminderEmails{})
	//    jobrunner.Schedule("@midnight", DataStats{}) // every midnight do this..
	//    jobrunner.Every(16*time.Minute, CleanS3{}) // evey 16 min clean...
	//    jobrunner.In(10*time.Second, WelcomeEmail{}) // one time job. starts after 10sec
	//    jobrunner.Now(NowDo{}) // do the job as soon as it's triggered

}

func ParamString(param string) string {
	return viper.GetString(param)
}

func ParamInt(param string) int {
	return viper.GetInt(param)
}

func ParamBool(param string) bool {
	return viper.GetBool(param)
}

func ParamDouelb(param string) float64 {
	return viper.GetFloat64(param)
}

func Params(param string) map[string]interface{} {
	return viper.GetStringMap(param)
}

// Job Specific Functions
type ReminderEmails struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	fmt.Printf("Every 5 sec send reminder emails \n")
}
