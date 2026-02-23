package main

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/joho/godotenv"
)

// App struct
type App struct {
	ctx context.Context
}

type Login struct {
	Username string
	Password string
}

type TeeTime struct {
	Course string
	Month  time.Month
	Day    time.Weekday
	Time   time.Time
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetTeeTimes(courses []string, t time.Time) (TeeTime, error) {
	for i := range courses {
		res, err := http.Get(courses[i])
		if err != nil {
			return TeeTime{}, err
		}
		fmt.Println(res, t)
	}
	return TeeTime{}, nil
}

func (a *App) ParseEnvFile(file string, vendor string) (Login, error) {

	login := Login{}
	// load config from .env file
	err := godotenv.Load(file)

	if err != nil {
		return login, err
	}

	var myEnv map[string]string
	myEnv, err = godotenv.Read()

	re := regexp.MustCompile(`^` + vendor + `/i`)
	pre := regexp.MustCompile(`(?i)\bpassword\s*$`)
	// search the env map values for password

	var matches []string
	for k := range myEnv {
		if re.MatchString(k) {
			matches = append(matches, k)
		}
	}

	if len(matches) > 2 {
		return login, fmt.Errorf("more than 2 matches in ENV for ", vendor)
	}

	for v := range matches {
		if pre.MatchString(matches[v]) {
			login.Password = matches[v]
		} else {
			login.Username = matches[v]
		}
	}

	return login, nil
}
