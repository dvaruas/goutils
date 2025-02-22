package protoutils

import (
	"bytes"
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func PrettyPrintProto(
	m proto.Message,
) string {
	b, err := protojson.Marshal(m)
	if err != nil {
		panic(err)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, b, "", " ")
	if err != nil {
		panic(err)
	}

	return prettyJSON.String()
}
