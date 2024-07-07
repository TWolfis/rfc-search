package fetch

import (
	"log"
	"os"
	"path"
	"sync"
	"time"

	"github.com/TWolfis/ietfrfc"
)

// fetch RFC's and safes them to disk
func FetcRfc(dstDir string, max int) error {
	var wg sync.WaitGroup

	for i := 1; i <= max; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			// fetch rfc
			rfc, err := ietfrfc.Get(num)
			if err != nil {
				log.Printf("Error %v while fetching RFC-%d\n", err, num)
				return
			}

			err = safeRfc(dstDir, rfc)
			if err != nil {
				log.Printf("Error %v while saving RFC-%d\n", err, num)
			}
		}(i)
		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()
	return nil
}

func safeRfc(dstDir string, rfc *ietfrfc.RFC) error {
	dstFile := path.Join(dstDir, "rfc-"+rfc.Number+".json")
	f, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(rfc.String())

	return nil
}
