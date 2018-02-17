package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// HandleCreateCdr create a cdr
func (resource *Resource) HandleFetchUser(c echo.Context) error {

	// bind input json with user model
	user := User{}
	if err := c.Bind(&user); err != nil {
		return CreateSuccessResponseWithoutData(&c, http.StatusBadRequest, "Please try again", "Operation failed")
	}

	go func(resource *Resource, user *User) {
		emails := resource.Process(user)

		fmt.Println("All emails")
		fmt.Println(emails)

		var vEmails []string

		for _, email := range emails {
			success, _ := resource.EmailService.VerifyEmail(email)
			if success {
				vEmails = append(vEmails, email)
			}
		}

		fmt.Println("Verified emails")
		fmt.Println(vEmails)

	}(resource, &user)

	return CreateSuccessResponseWithoutData(&c, http.StatusOK, "Success", "Processing started")
}

func (resource *Resource) Process(user *User) []string {
	tlds := tlds()
	var result []string

	comb1 := user.FullName()
	comb2 := user.Initials()
	comb3 := user.Combination1()
	comb4 := user.Combination2()

	result = append(result, resource.generatePossibleEmail(tlds, comb1, user.Company)...)
	result = append(result, resource.generatePossibleEmail(tlds, comb2, user.Company)...)
	result = append(result, resource.generatePossibleEmail(tlds, comb3, user.Company)...)
	result = append(result, resource.generatePossibleEmail(tlds, comb4, user.Company)...)

	// possible infixes
	infixes := infixes()
	for _, infix := range infixes {
		comb5, err := user.Combination3(infix)
		if err == nil {
			result = append(result, resource.generatePossibleEmail(tlds, comb5, user.Company)...)
		}
	}

	return result
}

func (resource *Resource) generatePossibleEmail(tlds []string, src string, company string) []string {
	if src == "" || company == "" {
		return []string{}
	}

	var result []string
	var c = strings.TrimSpace(company)

	for _, tld := range tlds {
		email := fmt.Sprintf("%s@%s%s", src, c, tld)
		result = append(result, email)
	}
	return result
}

func tlds() []string {
	return []string{".org", ".com", ".edu", ".net", ".uk", ".us"}
}

func infixes() []string {
	return []string{".", "_", "-"}
}

// response
func CreateBadResponse(c *echo.Context, requestCode int, message string, subMessage string) error {

	localC := *c
	response := fmt.Sprintf("{\"data\":{},\"message\":%q,\"submessage\":%q}", message, subMessage)
	return localC.JSONBlob(requestCode, []byte(response))
}

func CreateSuccessResponseWithoutData(c *echo.Context, requestCode int, message string, subMessage string) error {

	localC := *c
	response := fmt.Sprintf("{\"data\":{},\"message\":%q,\"submessage\":%q}", message, subMessage)
	return localC.JSONBlob(requestCode, []byte(response))
}
