package db

import (
	"log"

	"github.com/minio/minio-go/v7"
)

var Minio *minio.Client

func MinioInit() {
	endpoint := "localhost:9010"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
}
