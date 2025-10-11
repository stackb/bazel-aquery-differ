package protobuf

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type marshaler func(m protoreflect.ProtoMessage) ([]byte, error)
type unmarshaler func(b []byte, m protoreflect.ProtoMessage) error

func unmarshalerForFilename(filename string) (unmarshaler, string) {
	ext := filepath.Ext(filename)
	switch ext {
	case ".json", ".jsonproto":
		return protojson.Unmarshal, "json"
	case ".textproto":
		return prototext.Unmarshal, "text"
	case ".proto", ".pb":
		return proto.Unmarshal, "proto"
	default:
		// Default to binary proto for unknown extensions
		return proto.Unmarshal, "proto"
	}
}

func marshalerForFilename(filename string) marshaler {
	ext := filepath.Ext(filename)
	switch ext {
	case ".json", ".jsonproto":
		return protojson.Marshal
	case ".textproto":
		return prototext.Marshal
	case ".proto", ".pb":
		return proto.Marshal
	default:
		// Default to binary proto for unknown extensions
		return proto.Marshal
	}
}

func ReadFile(filename string, message protoreflect.ProtoMessage) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read %q: %w", filename, err)
	}
	unmarshaler, name := unmarshalerForFilename(filename)
	if err := unmarshaler(data, message); err != nil {
		return fmt.Errorf("unmarshal %q: %w", name, err)
	}
	return nil
}

func WriteFile(filename string, message protoreflect.ProtoMessage) error {
	data, err := marshalerForFilename(filename)(message)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

func WritePrettyJSONFile(filename string, message protoreflect.ProtoMessage) error {
	marshaler := protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		EmitUnpopulated: false,
	}
	data, err := marshaler.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

func WritePrettyTextFile(filename string, message protoreflect.ProtoMessage) error {
	marshaler := prototext.MarshalOptions{
		Multiline: true,
		EmitASCII: true,
	}
	data, err := marshaler.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

func FormatProtoText(message protoreflect.ProtoMessage) string {
	marshaler := prototext.MarshalOptions{
		Multiline: true,
		EmitASCII: true,
	}
	return marshaler.Format(message)
}
