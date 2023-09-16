package utils

import (
	"fmt"
	"os"
	"todoapp/database/models"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type environment struct {
	FOUNDER_EMAIL 		string
	EMAIL_API_KEY 		string
	VERIFY_TEMPLATE_ID 	string
	DEV_CLIENT_URL		string
}

func getEnvironmentVariables() environment {
	return environment{
		FOUNDER_EMAIL: os.Getenv("FOUNDER_EMAIL"),
		EMAIL_API_KEY: os.Getenv("EMAIL_API_KEY"),
		VERIFY_TEMPLATE_ID: os.Getenv("VERIFY_TEMPLATE_ID"),
		DEV_CLIENT_URL: os.Getenv("DEV_CLIENT_URL"),
	}
}

func SendVerificationEmail(user models.User, email, token string) error {
		 m := mail.NewV3Mail()
		 address := getEnvironmentVariables().FOUNDER_EMAIL
		 name := "Panoups"
		 e := mail.NewEmail(name, address)
		 m.SetFrom(e)
		 m.SetTemplateID(getEnvironmentVariables().VERIFY_TEMPLATE_ID)

		 p := mail.NewPersonalization()
		 to := mail.NewEmail("", email);
		 p.AddTos(to)

		 p.SetDynamicTemplateData("url", getEnvironmentVariables().DEV_CLIENT_URL + "register/" + token);

		 m.AddPersonalizations(p)

		 sendGridRequest := sendgrid.GetRequest(getEnvironmentVariables().EMAIL_API_KEY, "/v3/mail/send", "https://api.sendgrid.com");
		 sendGridRequest.Method = "POST"
		 sendGridRequest.Body = mail.GetRequestBody(m)
		 response, err := sendgrid.API(sendGridRequest)

		 if err != nil {
			 return err;
		 }

		 fmt.Println(response);

		return nil;
}
