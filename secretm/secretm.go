package secretm

import (
	//"encoding/json"
	"fmt"
	//"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	// "github.com/Kkamerunn/gambitUser/awsgo"
	"github.com/Kkamerunn/gambitUser/models"
)

func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	//var dataSecret models.SecretRDSJson
	fmt.Println(" > Ask for secret " + nameSecret)

	//svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	//key, err := svc.GetSecretValue()
}