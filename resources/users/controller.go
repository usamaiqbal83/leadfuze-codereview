package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"encoding/json"
	"sync"
	"time"
)

const rateLimiterBufferLength = 15

type VerifyEmailResponse struct {
	VerifiedEmail string `json:"verifiedEmail"`
}

// HandleCreateCdr create a cdr
func (resource *Resource) HandleFetchUser(c echo.Context) error {

	// bind input json with user model
	user := User{}
	if err := c.Bind(&user); err != nil {
		return CreateSuccessResponseWithoutData(&c, http.StatusBadRequest, "Please try again", "Operation failed")
	}

	emails := resource.Process(&user)

	fmt.Println("All emails")
	fmt.Println(emails)

	result, err := resource.FindVerifiedEmail(emails)
	if err != nil {
		res, _ := json.Marshal(VerifyEmailResponse{})
		return CreateSuccessResponse(&c,http.StatusOK, "No Data", "No verified email found", res)
	}

	res, _ := json.Marshal(VerifyEmailResponse{VerifiedEmail:result})
	return CreateSuccessResponse(&c,http.StatusOK, "Data found", "Verified email found", res)
}

func (resource *Resource)FindVerifiedEmail(emails []string) (string, error){
	// channel which includes tokens for rate limiting
	var tokens = make(chan int, rateLimiterBufferLength)
	resource.preFillTokens(tokens)
	// channel to indicate that required result had been found
	var found = make(chan string,1)

	var wg sync.WaitGroup
	wg.Add(len(emails))

	fmt.Println(len(emails))
	for _, email := range emails {
		resource.consumeToken(tokens)

		select {
		case d := <-found:
			return d, nil
		default:
			go func(resource *Resource, email string, ch chan int, res chan string) {
				defer wg.Done()
				fmt.Println("Processing for email : ", email)
				success, _ := resource.EmailService.VerifyEmail(email)
				if success {
					found <- email
				}
				resource.releaseToken(ch)
			}(resource, email, tokens, found)
		}
	}

	// if program reaches here that means all batches have been sent to webservices
	// all responses may not have arrived yet
	// so wait for the responses to arrive
	var done = make(chan struct{})
	go func() {
		// wait for processing of all emails
		wg.Wait()
		close(done)
	}()

	// following conditions actually breaks the waiting if all emails have been sent for processing and now waiting for results
	// i) wait group's waiting ends as all email responses are received and no one was verified
	// ii) time out after 1 minute
	// iii) any email got verified
	select {
		case <- done:
			return "", errors.New("No verified email found")
		case <-time.After(time.Minute * 1):
			return "", errors.New("No verified email found")
		case d := <-found:
			return d, nil
	}
	return "", errors.New("No verified email found")
}

func (resource *Resource) preFillTokens(buffer chan int) {
	for i := 0; i < rateLimiterBufferLength ; i++ {
		buffer <- 1
	}
}

func (resource *Resource) consumeToken(buffer chan int) {
	<- buffer
}

func (resource *Resource) releaseToken(buffer chan int) {
	buffer <- 1
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

func CreateSuccessResponse(c *echo.Context, requestCode int, message string, subMessage string, data []byte) error {

	localC := *c
	response := fmt.Sprintf("{\"data\":%s,\"message\":%q,\"submessage\":%q}", data, message, subMessage)
	fmt.Print(response)
	return localC.JSONBlob(requestCode, []byte(response))
}