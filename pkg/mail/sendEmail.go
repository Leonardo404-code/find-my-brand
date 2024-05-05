package mail

import (
	"crypto/tls"
	"fmt"

	"github.com/Leonargo404-code/find-my-brand/internal/env"
	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(ads *brand.Result, userEmail string) error {
	d := gomail.NewDialer(host, port, env.GetString(USER), env.GetString(PASSWORD))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", env.GetString(USER))
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Aqui estão os concorrentes que estão usando seus termos")

	htmlBody := createHtmlBody(ads)
	m.SetBody("text/html", htmlBody)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error in send e-mail: %v", err)
	}

	return nil
}

func createHtmlBody(result *brand.Result) string {
	tableRows := ""

	for _, ad := range result.Ads {
		tableRows += fmt.Sprintf(
			`<tr><td>%d</td><td>%s</td><td><a href="%s">%s</a></td><td><a href="%s">%s</a></td></tr>`,
			ad.Position,
			ad.Title,
			ad.Domain,
			ad.Domain,
			ad.Link,
			ad.Link,
		)
	}

	htmlBody := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Aqui estão suas concorentes</title>
		<style type="text/css">
			table, th, td {
				border: 1px solid black;
				padding: 0.3rem;
			}
		</style>
	</head>
	<body>
		<table>
			<thead>
				<th>Posição</th>
				<th>Titulo</th>
				<th>Domínio</th>
				<th>Link</th>
			</thead>
			<tbody>
				%s
			</tbody>
		</table>
	</body>
	</html>
`, tableRows)

	return htmlBody
}
