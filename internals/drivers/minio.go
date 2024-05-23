package drivers

import (
	"context"
	"fmt"

	"github.com/daqiancode/env"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinIO() (*minio.Client, error) {
	return minio.New(env.Get("S3_ENDPOINT"), &minio.Options{
		Region: env.Get("S3_REGION"),
		Creds:  credentials.NewStaticV4(env.Get("S3_ACCESS_KEY"), env.Get("S3_SECRET_KEY"), env.Get("S3_SECRET_TOKEN")),
		Secure: env.GetBoolMust("S3_USE_SSL", true),
	})
}

func CreateBucket(bucket string, isPublic bool) error {
	client, err := NewMinIO()
	if err != nil {
		return err
	}
	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	err = client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}
	if isPublic {
		// set public readonly policy
		policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::%s/*"]}]`
		return client.SetBucketPolicy(context.Background(), bucket, fmt.Sprintf(policy, bucket))
	}
	return nil
}
