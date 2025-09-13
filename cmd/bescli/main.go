package main

import (
	"log"

	"github.com/bazelbuild/bazelapis/src/main/java/com/google/devtools/build/lib/buildeventstream/proto/build_event_stream/build_event_stream"
)

func main() {
	log.Printf("hello, %+v", build_event_stream.BuildEventId{})
}
