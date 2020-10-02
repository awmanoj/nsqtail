package nsq

import "log"

type nsqModule struct {
	nsqLookupdAddr string
}

var instantiated *nsqModule = nil

func Init(nsqLookupdAddr string) *nsqModule {
	if instantiated == nil {
		instantiated = &nsqModule{
			nsqLookupdAddr: nsqLookupdAddr,
		}
	}
	log.Printf("Init NSQLookupdAddr: %s\n", nsqLookupdAddr)
	return instantiated
}
