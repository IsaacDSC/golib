package l4g

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	b64 "encoding/base64"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	EnableTracing  = "tracer:true"
	DisableTracing = "tracer:false"
	RepoLayer      = "layer:repository"
	RecService     = "recover:service"
)

func ManageError(err error, tags ...string) (output bool) {
	errorID := uuid.NewString()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("RECOVERED: %v\n", err)
			output = true
		}
	}()
	if err != nil {
		if len(tags) > 0 {
			var enableTracing bool
			structuredTags := map[string]string{}
			structuredTags["errorID"] = errorID
			structuredTags["time"] = time.Now().Format(time.RFC3339)
			for index := range tags {
				t := strings.Split(tags[index], ":")
				structuredTags[t[0]] = t[1]
				err = errors.WithMessage(err, tags[index])
				if tags[index] == "tracer:true" {
					enableTracing = true
				}
			}
			if enableTracing {
				out, _ := json.Marshal(structuredTags)
				sEnc := b64.StdEncoding.EncodeToString(out)
				fmt.Printf("LOGGER IN TRACER: %+v\n", sEnc)
				sDec, _ := b64.StdEncoding.DecodeString(sEnc)
				fmt.Printf("CONVERT READ THE TRACER: %+v\n", string(sDec))
				return true
			}
			fmt.Printf("TRATAMENTO: %v\n", err)
			return true
		}
		fmt.Printf("TRATAMENTO: %v\n", err)
		return true
	}
	return false
}
