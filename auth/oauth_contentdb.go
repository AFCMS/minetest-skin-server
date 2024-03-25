package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"

	"minetest-skin-server/utils"
)

func CDBRedirectURL() string {
	var cdbURL, err = url.Parse(utils.ConfigOAuthRedirectHost)
	if err != nil {
		log.Fatalln(err)
	}
	cdbURL.Path = "/api/account/providers/contentdb/callback"
	return cdbURL.String()
}

type ContentDBTokenResponse struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Error       string `json:"error"`
}

type ContentDBUser struct {
	Username        string `json:"username"`
	IsAuthenticated bool   `json:"is_authenticated:"`
}

func ContentDBAuthorize(c fiber.Ctx) error {
	cdbURL, _ := url.Parse(utils.ConfigOAuthContentDBURL)
	cdbURL.Path = "/oauth/authorize/"
	cdbURL.RawQuery = url.Values{
		"response_type": {"code"},
		"client_id":     {utils.ConfigOAuthContentDBClientID},
		"redirect_uri":  {CDBRedirectURL()},
	}.Encode()
	log.Println("Authorizing")

	return c.Redirect().To(cdbURL.String())
}

func ContentDBExchange(code string) (string, error) {
	c := client.New()
	req := c.R()

	req.SetHeader(fiber.HeaderContentType, fiber.MIMEMultipartForm)
	req.SetFormDatas(map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     utils.ConfigOAuthContentDBClientID,
		"client_secret": utils.ConfigOAuthContentDBClientSecret,
		"code":          code,
	})

	cdbURL, _ := url.Parse(utils.ConfigOAuthContentDBURL)
	cdbURL.Path = "/oauth/token/"

	log.Println(cdbURL.String())
	resp, err := req.Post(cdbURL.String())
	if err != nil {
		return "", err
	}

	result := ContentDBTokenResponse{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", err
	}

	if !result.Success {
		return "", errors.New("CDB: " + result.Error)
	}

	return result.AccessToken, nil
}

func ContentDBFetchUser(token string) (*ContentDBUser, error) {
	c := client.New()
	req := c.R()

	req.SetHeaders(map[string]string{
		fiber.HeaderAuthorization: "Bearer " + token,
		fiber.HeaderAccept:        fiber.MIMEApplicationJSON,
	})
	resp, err := req.Get("https://content.minetest.net/api/whoami/")
	if err != nil {
		return nil, err
	}

	d := new(ContentDBUser)

	err = resp.JSON(&d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func RegisterEndpoints(group fiber.Router) {
	group.Get("/contentdb", ContentDBAuthorize)

	group.Get("/contentdb/callback", func(c fiber.Ctx) error {
		code := c.Query("code")

		token, err := ContentDBExchange(code)
		if err != nil {
			return err
		}

		user, err := ContentDBFetchUser(token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "ContentDB error", "data": err.Error(),
			})
		}

		log.Println(user.Username)

		return c.Redirect().To("/")
	})
}
