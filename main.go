package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	browser, err := CreateBrowser()
	if err != nil {
		panic(err)
	}
	defer browser.Quit() // This line makes the browser quit

	browser.Get("https://www.instagram.com//")
	time.Sleep(time.Second * 3)

	username_textbox, err := browser.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div/div[2]/div/div/div/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[1]/div/label/input")
	if err != nil {
		fmt.Println("err")
		os.Exit(0)
	}
	username_textbox.SendKeys("") //Enter instagram username

	password_textbox, err := browser.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div/div[2]/div/div/div/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[2]/div/label/input")
	if err != nil {
		fmt.Println("err")
		os.Exit(0)
	}
	password_textbox.SendKeys("") //Enter instagram password

	login_button, err := browser.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div/div[2]/div/div/div/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[3]")
	if err != nil {
		fmt.Println("err")
		os.Exit(0)
	}
	login_button.Click()
	time.Sleep(time.Second * 5)

	browser.Get("https://www.instagram.com/eminem")
	time.Sleep(time.Second * 4)

	follow_button, err := browser.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div/div[2]/div/div/div/div[1]/div[1]/div[2]/div[2]/section/main/div/header/section/div[1]/div[2]/div/div[1]/button")
	if err != nil {
		fmt.Println("err")
		os.Exit(0)
	}
	follow_button.Click()
	time.Sleep(time.Second * 20)

}

func CreateBrowser() (selenium.WebDriver, error) {
	browserPath := os.Getenv("CHROME_PATH")
	port := 8080

	var opts []selenium.ServiceOption
	_, err := selenium.NewChromeDriverService(os.Getenv("CHROMEDRIVER_PATH"), port, opts...)

	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	// NB: Path is the important part here
	caps.AddChrome(
		chrome.Capabilities{
			Path: browserPath,
			Args: []string{
				"--window-size=800,640",
				// "--headless",
				"--no-sandbox",
			},
		},
	)

	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:"+strconv.Itoa(port)+"/wd/hub")

	return wd, err

}
