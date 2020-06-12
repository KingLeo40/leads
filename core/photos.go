package core

import "github.com/louisevanderlith/husk"

type Photos struct {
	Front    husk.Key
	Left     husk.Key
	Right    husk.Key
	Rear     husk.Key
	Interior husk.Key
	Engine   husk.Key
}
