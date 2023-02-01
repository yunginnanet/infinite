package main

import (
	"flag"
	"net/http"
	"os"
	"path"

	"github.com/yunginnanet/infinite/components"
	"github.com/yunginnanet/infinite/components/progress"
)

var urlF = flag.String("d", "", "download url")

func init() {
	flag.Parse()
}

func main() {
	url := *urlF
	progress.NewGroupWithCount(1).
		AppendRunner(func(pro *components.Progress) func() {
			resp, err := http.Get(url)
			if err != nil {
				pro.Println("get error", err)
				resp.Body.Close()
				return func() {}
			}
			pro.WithTotal(resp.ContentLength)

			return func() {
				defer resp.Body.Close()

				dest, err := os.OpenFile(path.Base(url), os.O_CREATE|os.O_WRONLY, 0o777)
				defer dest.Close()
				if err != nil {
					pro.Println("dest open error", err)
					return
				}

				_, err = progress.StartTransfer(resp.Body, dest, pro)
				if err != nil {
					pro.Println("trans error", err)
				}
			}
		}).Display()
}
