package main
import(
	"os"
)
const (
	VIDEO_DIR =`\videos\`
	MAX_UPLOAD_SIZE =1024*1024*1024
)
func GetCompletePath()string{
	gobin :=os.Getenv("GOBIN")
	complete_pt:=gobin +VIDEO_DIR
	return complete_pt
}