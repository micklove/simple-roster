package UUID

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"time"
)

type KSUUIDGenerator struct {
}

func (KSUUIDGenerator) Create() (id string, err error) {
	var tmpId ksuid.KSUID
	if tmpId, err = ksuid.NewRandomWithTime(time.Now()); err != nil {
		fmt.Printf("error getting new UUID for User ID, err [%v]", err)
		return "", err
	}
	return tmpId.String(), nil
}
