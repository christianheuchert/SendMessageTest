package SendPriorityMessageToAssets

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"fmt"
	// "io"
	"net/http"
	"net/url"

	"github.com/AiRISTAFlowInc/fs-core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	customerId := getCustomerId(input.IP, input.Username, input.Password)
	Status := SendPriorityMessageToAssets(input, customerId)

	output := &Output{Status: Status}

	// fmt.Println("Output: ", output.Status)
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

func SendPriorityMessageToAssets(input *Input, CustomerId string)string{

	// filtering url string's special characters
	FilteredMessage := url.QueryEscape(input.Message)

	// Create an HTTP client
	client := &http.Client{}

	var asset Asset
	assetCheck := json.Unmarshal([]byte(input.StaffIdList), &asset) // check if Asset Obj
	if (assetCheck == nil){ // if no error puting into asset struct, then obj
		input.StaffIdList = strconv.Itoa(asset.ID)
	}

	//Hardcoded variables
	MessagePriority := "1" //high
	MessageOptions := displayItemToValue(input.MessageOptions)
	MessageId := "1"


	// Create the request
	url := "http://" + input.IP + "/XpertRestApi/api/Device/DisplayPriorityMessageOnTagNotif?" +
		"StaffIdList=" + input.StaffIdList + "&" +
		"Message=" + FilteredMessage + "&" +
		"MessageId=" + MessageId + "&" +
		"MessagePriority=" + MessagePriority + "&" +
		"MessageOptions=" + MessageOptions + "&" +
		"MessageExpirationSeconds=" + strconv.Itoa(input.MessageExpiration) + "&" +
		"CustomerId=" + CustomerId
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err.Error()
	}

	req.Header.Add("Content-Type", "application/json")

	// Add basic authentication to the request header
	auth := input.Username + ":" + input.Password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err.Error()
	}
	defer resp.Body.Close()

	returnedStatus := "false"
	if(resp.Status == "200 OK"){
		returnedStatus = "true"
	}

	return returnedStatus
}

func displayItemToValue(displayItem string) string {
  switch displayItem {
  case "Beep":
    return "tone-four"
  case "Twinkle":
    return "tone-two"
  case "Stairs":
    return "tone-three"
  case "Siren":
    return "tone-one"
  case "Silent":
    return "silent"
  default:
    return "tone-four"
  }
}

// Call login route for customerId
func getCustomerId(IP string,  username string, password string )string {
	// Create an HTTP client
	client := &http.Client{}
	IpNoPort := strings.Split(IP, ":")[0]
	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/Login/Login?IpAddress="+IpNoPort
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the config JSON into a response struct object
	var response Login
	errUnmarshal := json.Unmarshal([]byte(body), &response)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
		return ""
	}

	return  strconv.Itoa(response.CustomerID)
}