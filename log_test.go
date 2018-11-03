package kclogs

import(
	"testing"
	// "os"
)
func TestDebug(t *testing.T){
	InitLog("test.log","debug","json")
	Log.Info("hello")
}