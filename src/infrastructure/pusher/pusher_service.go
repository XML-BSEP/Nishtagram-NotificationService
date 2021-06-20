package pusher

import (
	"github.com/pusher/pusher-http-go/v5"
	"github.com/spf13/viper"
)

type pusherService struct {
	PusherClient *pusher.Client
}

type PusherService interface {
	Trigger(channel string, event string, data interface{}) error
}

func NewPusherService(pusherClient *pusher.Client) PusherService {
	return &pusherService{PusherClient: pusherClient}
}

func (p *pusherService) Trigger(channel string, event string, data interface{}) error {
	return p.PusherClient.Trigger(channel, event, data)
}

func init_viper() error{
	viper.SetConfigFile(`configuration/pusher.json`)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	if viper.GetBool(`debug`) {

	}
	return nil
}

func GetConnection() *pusher.Client{
	err := init_viper()
	if err != nil {
		return nil
	}

	appId := viper.GetString("pusher.APP_ID")
	appKey := viper.GetString("pusher.APP_KEY")
	appSecret := viper.GetString("pusher.APP_SECRET")
	appCluster := viper.GetString("pusher.APP_CLUSTER")

	return &pusher.Client{Key: appKey, AppID: appId, Secret: appSecret, Cluster: appCluster}
}
