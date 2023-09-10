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
}

func getEnvironmentVariables() environment {
	return environment{
		FOUNDER_EMAIL: os.Getenv("FOUNDER_EMAIL"),
		EMAIL_API_KEY: os.Getenv("EMAIL_API_KEY"),
		VERIFY_TEMPLATE_ID: os.Getenv("VERIFY_TEMPLATE_ID"),
	}
}

func SendVerificationEmail(user models.User, email string) error {

		fmt.Println(getEnvironmentVariables())

		 m := mail.NewV3Mail()
		 address := getEnvironmentVariables().FOUNDER_EMAIL
		 name := "Panoups"
		 e := mail.NewEmail(name, address)
		 m.SetFrom(e)
		 m.SetTemplateID(getEnvironmentVariables().VERIFY_TEMPLATE_ID)

		 p := mail.NewPersonalization()
		 to := mail.NewEmail("", "delgeoffrey1@gmail.com");
		 p.AddTos(to)

		 p.SetDynamicTemplateData("url", "https://google.com");

		 m.AddPersonalizations(p)

		 sendGridRequest := sendgrid.GetRequest(getEnvironmentVariables().EMAIL_API_KEY, "/v3/mail/send", "https://api.sendgrid.com");
		 sendGridRequest.Method = "POST"
		 sendGridRequest.Body = mail.GetRequestBody(m)
		 response, err := sendgrid.API(sendGridRequest)

		 if err != nil {
			 return err;
		 }

		 fmt.Println(response);

		 fmt.Println("Email envoyé avec succès");

		return nil;
}
