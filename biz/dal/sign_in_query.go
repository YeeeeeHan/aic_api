package dal

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"image"
// 	"image/jpeg"
// 	"io/ioutil"
// 	"net/http"
// 	"os"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// 	"gitlab.com/aic/aic_api/biz/model/gitlab.com/aic/http_idl_gen/gen/aic/data/models"
// 	"gitlab.com/aic/aic_api/biz/util/errors"
// 	"gitlab.com/aic/aic_api/biz/util/helpers"
// 	"gitlab.com/aic/aic_api/logs"
// )

// type JsonDP struct {
// 	ProfilePictureDP `json:"profilePicture"`
// }

// type ProfilePictureDP struct {
// 	DisplayImageDP `json:"displayImage~"`
// }

// type DisplayImageDP struct {
// 	Elements []ElementsDP `json:"elements"`
// }

// type ElementsDP struct {
// 	Identifiers []IdentifiersDP `json:"identifiers"`
// }

// type IdentifiersDP struct {
// 	Identifier string `json:"identifier"`
// }

// func RegisterNewUser(ctx context.Context, d *models.LinkedInData, accessToken string) (*models.Profile, string, error) {
// 	emailUrl := "https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))"

// 	req, err := http.NewRequest("GET", emailUrl, nil)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+accessToken)

// 	client := &http.Client{}
// 	response, err := client.Do(req)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	defer response.Body.Close()

// 	type Element struct {
// 		Handle struct {
// 			EmailAddress string `json:"emailAddress"`
// 		} `json:"handle~"`
// 	}

// 	type Data struct {
// 		Elements []Element `json:"elements"`
// 	}

// 	content, err := ioutil.ReadAll(response.Body)

// 	if err != nil {
// 		return nil, "", err
// 	}

// 	var data Data
// 	err = json.Unmarshal(content, &data)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	if len(data.Elements) != 1 {
// 		return nil, "", errors.NewInternalError("unable to fetch email properly")
// 	}

// 	email := data.Elements[0].Handle.EmailAddress

// 	//Getting profile pic metadata using accessToken
// 	imageURL := "https://api.linkedin.com/v2/me?projection=(id,firstName,lastName,emailAddress,profilePicture(displayImage~:playableStreams))&oauth2_access_token=" + accessToken
// 	response, err = http.Get(imageURL)
// 	if err != nil {
// 		logs.CtxError(ctx, "error downloading image: %v", err)
// 		return nil, "", err
// 	}
// 	defer response.Body.Close()

// 	//Reading profile pic metadata
// 	imageData, err := ioutil.ReadAll(response.Body)

// 	if err != nil {
// 		logs.CtxError(ctx, "error reading json image data: %v", err)
// 		return nil, "", err

// 	}

// 	var jsonDP JsonDP
// 	err = json.Unmarshal(imageData, &jsonDP)

// 	if err != nil {
// 		logs.CtxError(ctx, "error unmarshalling image data: %v", err)
// 		return nil, "", err
// 	}

// 	var dpData []byte //Profile picture data in bytes

// 	if jsonDP.ProfilePictureDP.DisplayImageDP.Elements != nil {

// 		//Obtaining the URL of the profile pic stored on LinkedIn server
// 		dpUrl := jsonDP.ProfilePictureDP.DisplayImageDP.Elements[2].Identifiers[0].Identifier
// 		response, err = http.Get(dpUrl)
// 		if err != nil {
// 			logs.CtxError(ctx, "error reading image data: %v", err)
// 			return nil, "", err

// 		}
// 		dpData, err = ioutil.ReadAll(response.Body)
// 		if err != nil {
// 			logs.CtxError(ctx, "error reading image data with ioutil: %v", err)
// 			return nil, "", err
// 		}

// 	} else { //Profile picture is not available from Linkedin
// 		dpData = getDefaultDp(ctx)
// 	}

// 	return createProfile(ctx, d.Id, dpData, d.LocalizedFirstName, d.LocalizedLastName, email)
// }

// func RegisterNewAppleUser(ctx context.Context, params *models.AppleSignInRequestParams) (*models.Profile, string, error) {
// 	dpData := getDefaultDp(ctx)
// 	return createProfile(ctx, params.LinkedinId, dpData, params.GivenName, params.FamilyName, params.Email)

// }

// func createProfile(ctx context.Context, linkedInId string, dpData []byte, firstName, lastName, email string) (*models.Profile, string, error) {

// 	//Creating new AWS S3 session
// 	sess, err := session.NewSession(&aws.Config{
// 		Region:      aws.String("ap-southeast-1"),
// 		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")})

// 	if err != nil {
// 		logs.CtxError(ctx, "error creating aws session: %v", err)
// 		return nil, "", err

// 	}

// 	svc := s3.New(sess)

// 	// Set the name of the bucket and the key for the object
// 	bucket := "aic-bucket"
// 	key := linkedInId + ".jpg"

// 	_, err = svc.PutObject(&s3.PutObjectInput{
// 		Bucket: aws.String(bucket),
// 		Key:    aws.String(key),
// 		Body:   bytes.NewReader(dpData),
// 	})

// 	if err != nil {
// 		logs.CtxError(ctx, "error uploading image: %v", err)
// 		return nil, "", err

// 	}

// 	//Concatenating strings to obtain dp URL
// 	linkedInProfileURL := "https://aic-bucket.s3.ap-southeast-1.amazonaws.com/" + key

// 	id := helpers.GenerateID()

// 	newUser := &models.Profile{
// 		Id:          id,
// 		AccountType: "Student",
// 		ProfileLinkedin: &models.ProfileLinkedIn{
// 			LinkedinId:             linkedInId,
// 			LinkedinFirstName:      helpers.Ternary(firstName == "", "Please set FirstName", firstName),
// 			LinkedinLastName:       helpers.Ternary(lastName == "", "Please set LastName", lastName),
// 			LinkedinAdditionalName: "",
// 			LinkedinHeadline:       "",
// 			LinkedinProfileUrl:     linkedInProfileURL,
// 			Id:                     id,
// 		},
// 		ProfilePretexts: []string{
// 			"Hello, my name is XXX. I am interested in XXX, pleased to meet you!",
// 			"I wanted to reach out to you because XXX, and I believe that our shared interests could potentially lead to valuable collaboration.",
// 			"I hope we can meet up at XXX over a cup of coffee!",
// 		},
// 		Modules:              []string{},
// 		Interests:            []string{},
// 		Email:                email,
// 		NetworkingIntentions: []string{"Academic"},
// 	}

// 	err = newProfile(ctx, newUser, id)

// 	if err != nil {
// 		logs.CtxError(ctx, "rrror from db.NewProfile while registering new user: %v", err)
// 		return nil, "", err

// 	}

// 	logs.CtxInfo(ctx, "registered new user")
// 	return newUser, id, nil
// }

// func getDefaultDp(ctx context.Context) []byte {
// 	dir, _ := os.Getwd()
// 	img_path := dir + "/common/empty.jpg"
// 	emptyImg, err := os.Open(img_path)
// 	if err != nil {
// 		logs.CtxError(ctx, "error opening empty image file from local: %v", err)
// 	}

// 	img_Data, _, err := image.Decode(emptyImg)
// 	if err != nil {
// 		logs.CtxError(ctx, "error decoding jpg: %v", err)
// 	}

// 	var buf bytes.Buffer
// 	err = jpeg.Encode(&buf, img_Data, nil)
// 	if err != nil {
// 		logs.CtxError(ctx, "error converting jpg to bytes: %v", err)
// 	}
// 	return buf.Bytes()
// }
