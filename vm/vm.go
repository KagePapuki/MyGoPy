package vm

import (
	"github.com/KagePapuki/MyGoPy/decode"
)

var Path string
var Mem []string
var Address map[string]string
var pointer int

func init(p string) string {
	Path = p
	Mem = []string{}
	Address = map[string]string{}
	return "Succeed"
}

func run(command []*decode.Token) string {
	pointer = 0
	for pointer < len(command) {
		switch command[pointer].Trawtype() {
		case 0:
			//pass
		case 1:
			//pass
		case 2:
			//pass
		case 3:
			//pass
		case 4:
			//pass
		case 5:
			//pass
		case 6:
			//pass
		case 7:
			//pass
		default:
			//pass
		}
		
		pointer += 1
	}
}