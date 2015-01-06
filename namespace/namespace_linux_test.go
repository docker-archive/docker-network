package namespace

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDeleteNamespace(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "docker-network-ns-test")
	if err != nil {
		t.Fatal(err)
	}
	nsPath := filepath.Join(tmpdir, "ns")
	ns, err := New(nsPath)
	if err != nil {
		t.Fatal(err)
	}
	if ns.Path != nsPath {
		t.Fatalf("ns.Path should be %q, got %q", nsPath, ns.Path)
	}
	if err := os.Remove(nsPath); err == nil {
		t.Fatal("You can remove mounted namespace")
	}
	if err := ns.Delete(); err != nil {
		t.Fatalf("Failed to remove namespace: %s", err)
	}
	if _, err := os.Stat(ns.Path); !os.IsNotExist(err) {
		t.Fatalf("Namespace file %q should be removed. Error: %s", ns.Path, err)
	}
}
