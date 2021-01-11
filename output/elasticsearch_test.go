package output

import (
	"log"
	"testing"
)

func TestNewEs(t *testing.T) {
	toEs, err := NewEs([]string{"http://localhost:9200"},[]string{"www"},"2","1","7",true,
	"","","elastic","password")
	if err != nil {
		log.Println(err)
		return
	}
	toEs.ToEs(Storage)
}
