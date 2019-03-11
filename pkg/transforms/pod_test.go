package transforms

import (
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
)

func TestTransformPod(t *testing.T) {
	var p v1.Pod
	UnmarshalFile("../../test-data/pod.json", &p, t)
	node := TransformPod(&p)

	// Build time struct matching time in test data
	date := time.Date(2019, 02, 21, 21, 30, 33, 0, time.UTC)

	// Test only the fields that exist in pods - the common test will test the other bits
	AssertEqual("kind", node.Properties["kind"], "Pod", t)
	AssertEqual("hostIP", node.Properties["hostIP"], "1.1.1.1", t)
	AssertEqual("podIP", node.Properties["podIP"], "2.2.2.2", t)
	AssertEqual("restarts", node.Properties["restarts"], uint(2), t)
	AssertEqual("startedAt", node.Properties["startedAt"], date.UTC().Format(time.RFC3339), t)
	AssertEqual("status", node.Properties["status"], string(v1.PodRunning), t)

}
