package testingcommon

import (
	"bytes"
	"encoding/json"

	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
)

// Converts map created by json decoder to struct
// out should be pointer (&payload)
func ExtractPayload(response *httpo.ApiResponse, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(response.Payload)
	json.NewDecoder(buf).Decode(out)
}
