package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/julienschmidt/httprouter"
)

const (
	bucketName  = "arn:aws:s3:eu-central-1:566415016137:accesspoint/china-get" //"trackit.file"
	keyName     = "trackit"
	accountID   = "566415016137"
	accessPoint = "china-get"
)

func main() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	router := httprouter.New()
	router.POST("/upload", HandleUpload)
	router.GET("/download", HandleDownload)

	log.Fatal(http.ListenAndServe(":8080", router))

	// Create an Access Point referring to the bucket
	// fmt.Println("create an access point")
	// _, err := s3ControlSvc.CreateAccessPoint(&s3control.CreateAccessPointInput{
	// 	AccountId: aws.String(accountID),
	// 	Bucket:    aws.String(bucketName),
	// 	Name:      aws.String(accessPoint),
	// })
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to create access point: %v", err))
	// }

	// Use the SDK's ARN builder to create an ARN for the Access Point.
	// apARN := arn.ARN{
	// 	Partition: "aws",
	// 	Service:   "s3",
	// 	Region:    aws.StringValue(sess.Config.Region),
	// 	AccountID: accountID,
	// 	Resource:  "accesspoint/" + accessPoint,
	// }

	// And Use Access Point ARN where bucket parameters are accepted
	// fmt.Println("get object using access point")

	// params := &s3.GetObjectInput{
	// 	Bucket: aws.String(apARN.String()),
	// 	Key:    aws.String(keyName),
	// }
	// //params.SetRequestPayer(s3.RequestPayerRequester)
	// r, _ := s3Svc.GetObjectRequest(params)
	// q := r.HTTPRequest.URL.Query()
	// q.Set("x-amz-request-payer", "requester")
	// r.HTTPRequest.URL.RawQuery = q.Encode()
	// url, err := r.Presign(15 * time.Minute)

	// if err != nil {
	// 	panic(fmt.Sprintf("failed put object request: %v", err))
	// }

	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(req)
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(resp.Body)
	// if resp.StatusCode != http.StatusOK {
	// 	_ = resp.Write(os.Stdout)
	// 	log.Fatal(resp.Status)
	// }
	// _, err = ioutil.ReadAll(getObjectOutput.Body)
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to read object body: %v", err))
	// }
}

// HandleDownload provide download file for s3 accesspoint
func HandleDownload(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fileName := r.URL.Query().Get("name")
	filePath := r.URL.Query().Get("action")

	//var buff []byte

	// combine filePath + fileName
	switch filePath {
	case "avatar", "feedback", "public", "version":
		filePath = keyName + "/" + filePath + "/"
		log.Infof("action key %s \n", filePath)
	default:
		log.Infof("Upload FormFile unKnow Action :", filePath)
		filePath = keyName + "/public/"
	}

	// touch upload url
	url, err := getObjectRequest(bucketName, filePath+fileName)
	if err != nil {
		log.Fatal("fail to Generate url: ", err)
	}
	fmt.Println(url)
	// create new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("NewRequest Err : ", err)
	}

	//upload file for Binary
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("fail to upload: ", err)
	}

	w.Header().Set("Content-disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "binary/octet-stream")

	io.Copy(w, resp.Body)
	log.Info("Download SUCCESS ")
}

func HandleUpload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fileName := r.URL.Query().Get("name")
	filePath := r.URL.Query().Get("action")

	var buff []byte

	// combine filePath + fileName
	switch filePath {
	case "avatar", "feedback", "public", "version":
		filePath = keyName + "/" + filePath + "/"
		log.Infof("action key %s \n", filePath)
	default:
		log.Infof("Upload FormFile unKnow Action :", filePath)
		filePath = keyName + "/public/"
	}

	// touch upload url
	url, err := putObjectRequest(bucketName, filePath+fileName)
	if err != nil {
		log.Fatal("fail to Generate url: ", err)
	}

	// determine "Content-Type" value
	switch r.Header["Content-Type"][0] {
	case "image/jpeg":
		buff, err = ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Read r.Body error: ", err)
		}
	default:
		file, _, err := r.FormFile("file")
		buff, err = ioutil.ReadAll(file)
		if err != nil {
			log.Error("Read form-data error: ", err)
		}
	}

	// create new request
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(buff))
	if err != nil {
		log.Fatal("NewRequest Err : ", err)
	}

	//upload file for Binary
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("fail to upload: ", err)
	}

	// determine upload result
	if resp.StatusCode != http.StatusOK {
		_ = resp.Write(os.Stdout)
		log.Fatal(resp.Status)
		fmt.Fprintf(w, "hello, %s!\n", fileName)
	} else {
		fmt.Fprintf(w, "Upload SUCCESS ,Code is : %v", resp.StatusCode)
	}
}

// Generate signed url
// return get signature url
func getObjectRequest(bucket string, key string) (url string, err error) {
	sess := session.Must(session.NewSession())
	s3Svc := s3.New(sess)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	r, _ := s3Svc.GetObjectRequest(params)
	q := r.HTTPRequest.URL.Query()
	r.HTTPRequest.URL.RawQuery = q.Encode()
	return r.Presign(15 * time.Minute)
}

// Generate signed url
// return put signature url
func putObjectRequest(bucket string, key string) (url string, err error) {
	sess := session.Must(session.NewSession())
	s3Svc := s3.New(sess)

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	r, _ := s3Svc.PutObjectRequest(params)
	q := r.HTTPRequest.URL.Query()
	//q.Set("Content-Type", "application/octet-stream")
	r.HTTPRequest.URL.RawQuery = q.Encode()
	return r.Presign(15 * time.Minute)
}
