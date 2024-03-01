package utils

import (
	"Middleware-test/internal/models"
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net"
	"net/smtp"
	"os"
)

var configFile = Path + "config/config.json"

func generateConfirmationID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func generateMessageID(domain string) string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("<%s@%s>", base64.StdEncoding.EncodeToString(b), domain)

}

func fetchConfig() models.MailConfig {
	var config models.MailConfig

	data, err := os.ReadFile(configFile)

	if len(data) == 0 {
		Logger.Error(GetCurrentFuncName() + " No JSON config data found!")
		return config
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		Logger.Error(GetCurrentFuncName()+" JSON MarshalIndent error!", slog.Any("output", err))
		return config
	}

	return config
}

func SendMail(temp *models.TempUser) {
	// Fetching mail configuration
	config := fetchConfig()

	// Recipient information
	recipientMail := []string{temp.User.Email}

	// Generating confirmation Id
	temp.ConfirmID = generateConfirmationID()

	header := make(map[string]string)
	header["From"] = "Account management" + "<" + config.Email + ">"
	header["To"] = temp.User.Email
	header["Subject"] = "Email verification"
	header["Message-ID"] = generateMessageID(config.Hostname)
	header["Content-Type"] = "text/html; charset=UTF-8"

	t, err := template.ParseFiles(Path + "templates/mail.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	// Create a buffer to hold the formatted message
	var body bytes.Buffer

	// Execute template with data
	err = t.Execute(&body, struct {
		Username  string
		ConfirmID string
	}{
		Username:  temp.User.Username,
		ConfirmID: temp.ConfirmID,
	})
	if err != nil {
		log.Fatal(err)
	}

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body.String()

	auth := smtp.PlainAuth(
		"",
		config.Email,
		config.Auth,
		config.Hostname,
	)

	err = SendMailUsingTLS(
		fmt.Sprintf("%s:%d", config.Hostname, config.Port),
		auth,
		config.Email,
		recipientMail,
		[]byte(message),
	)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

// return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//Explode Host Port String
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

// refer to net/smtp func SendMail()
// When using net.Dial to connect to the tls (SSL) port, smtp. NewClient() will be stuck and will not prompt err
// When len (to)>1, to [1] starts to prompt that it is secret delivery
func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smtp client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
