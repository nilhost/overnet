/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-02-03
 */
package conf

import (
	"github.com/nilhost/overnet/app/onet"
)

func DefaultOnetConfig() *onet.Config {
	return &onet.Config{
		Ip:       "127.0.0.1",
		Port:     55155,
		Protocol: "tcp",
		PubId:    "default-p2p-pub-id",
		Seedlist: []*onet.Seed{
			{
				Protocol: "tcp",
				Ip:       "127.0.0.1",
				Port:     55255,
				PubId:    "seed1.pub.id",
			},
			{
				Protocol: "tcp",
				Ip:       "127.0.0.1",
				Port:     55355,
				PubId:    "seed2.pub.id",
			},
		},
	}
}

type _seed struct {
	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
	Port     int32  `json:"port"`
	PubID    string `json:"pub_id"`
}
type OnetConfig struct {
	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
	Port     int32  `json:"port"`
	PubID    string `json:"pub_id"`
	SeedList []_seed `json:"seedlist"`
}

func (v *OnetConfig) Build() *onet.Config {
	var seeds []*onet.Seed
	for _, value := range v.SeedList {
		seed := onet.Seed{
			Ip:       value.IP,
			Port:     value.Port,
			Protocol: value.Protocol,
			PubId:    value.PubID}
		seeds = append(seeds, &seed)
	}
	return &onet.Config{
		Ip:       v.IP,
		Port:     v.Port,
		Protocol: v.Protocol,
		PubId:    v.PubID,
		Seedlist: seeds,
	}
}
